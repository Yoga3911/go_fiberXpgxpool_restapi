package controllers

import (
	"coba/helper"
	"coba/models"
	"coba/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (a *authController) Login(c *fiber.Ctx) error {
	login := new(models.Login)
	err := c.BodyParser(login)
	
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.BuildResponse(err.Error(), false))
	}

	errors := helper.ErrorHandler(login)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BuildResponse(errors, false))
	}

	e := a.authService.VerifyUser(*login, c.Context())
	if e != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(e.Error(), false))
	}

	return c.JSON(helper.BuildResponse(nil, false))
}

func (a *authController) Register(c *fiber.Ctx) error {
	register := new(models.Register)
	err := c.BodyParser(register)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BuildResponse(err.Error(), false))
	}

	errors := helper.ErrorHandler(register)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BuildResponse(errors, false))
	}

	e := a.authService.CreateUser(*register, c.Context())
	if e != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(e.Error(), false))
	}
	return c.JSON(helper.BuildResponse(nil, false))
}
