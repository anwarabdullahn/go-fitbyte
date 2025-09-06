package handlers

import (
	"fmt"
	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/entities"
	"go-fitbyte/src/pkg/user"
	"go-fitbyte/src/pkg/userfile"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// UploadUserFile is handler/controller which upload file
// @Summary      Upload user file
// @Description  Upload user file
// @Tags         Upload File
// @Accept       multipart/form-data
// @Produce      json
// @Security BearerAuth
// @Param        file  formData  file  true  "User File"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/upload-file [post]
func UploadUserFile(userService user.Service, userfileService userfile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Ambil user_id dari JWT
		userIDStr, ok := c.Locals("user_id").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid token user_id"))
		}

		uid64, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid user_id format"))
		}

		// 2. Cek apakah user ada
		user, err := userService.FetchUserById(uint(uid64))
		if err != nil {
			return c.Status(fiber.StatusNotFound).
				JSON(presenter.ErrorResponse("user not found"))
		}

		// 3. Ambil data user_file lama
		existingFile, err := userfileService.GetUserFile(uint(uid64))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(presenter.UserFileErrorResponse(err))
		}

		var userfileID uint
		if existingFile != nil {
			userfileID = existingFile.ID
			if existingFile.URI != "" {
				// hapus file lama jika ada
				if _, err := os.Stat(existingFile.URI); err == nil {
					if err := os.Remove(existingFile.URI); err != nil {
						return c.Status(fiber.StatusInternalServerError).
							JSON(presenter.ErrorResponse("failed to remove old file: " + err.Error()))
					}
				}
			}
		}

		// 4. Ambil file baru dari form-data
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(presenter.ErrorResponse("file is required"))
		}
		//validate file
		ext := strings.ToLower(filepath.Ext(file.Filename))
		fmt.Println(ext)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return c.Status(fiber.StatusBadRequest).SendString("Hanya file gambar yang diperbolehkan")
		}

		//validasi ukuran
		// Validasi ukuran file (max 100 KB)
		const maxSize = 100 * 1024 // 100 KB
		if file.Size > maxSize {
			return c.Status(fiber.StatusBadRequest).SendString("Ukuran file maksimal 100KB")
		}

		// bikin folder user -> uploads/<email>/
		dirPath := filepath.Join("uploads", user.Email)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(presenter.ErrorResponse("failed to create dir: " + err.Error()))
		}

		savePath := filepath.Join(dirPath, file.Filename)

		// 5. Simpan file ke folder
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(presenter.ErrorResponse("failed to save file: " + err.Error()))
		}

		// 6. Simpan metadata user_file
		userFile := entities.UserFile{
			ID:     userfileID,
			UserID: user.ID,
			URI:    savePath,
		}

		result, err := userfileService.UploadUserFile(&userFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(presenter.UserFileErrorResponse(err))
		}

		return c.JSON(presenter.UserFileSuccessResponse(result))
	}
}
