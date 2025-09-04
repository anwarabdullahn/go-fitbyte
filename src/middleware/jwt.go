package middleware

import (
	"go-fitbyte/src/api/presenter"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).
				JSON(presenter.ErrorResponse("missing token"))
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid token"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("user_id", claims["user_id"])
		}

		return c.Next()
	}
}
