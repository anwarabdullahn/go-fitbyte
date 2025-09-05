package routes

import (
	"go-fitbyte/src/api/handlers"
	"go-fitbyte/src/api/middleware"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/user"
	"go-fitbyte/src/pkg/userfile"

	"github.com/gofiber/fiber/v2"
)

func UserfileRouter(app fiber.Router, userService user.Service, userfileService userfile.Service, jm *auth.JWTManager) {
	app.Post(
		"/upload-file",
		middleware.JWTProtected(jm),
		handlers.UploadUserFile(userService, userfileService),
	)

}
