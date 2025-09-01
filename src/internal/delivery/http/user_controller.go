package http

import (
	"go-fitbyte/src/internal/model"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	return ctx.JSON(request)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {

		return fiber.ErrBadRequest
	}

	return ctx.JSON(request)
}
