package controller

import "github.com/gofiber/fiber/v2"

type AuthorController interface {
	RegisterController(c *fiber.Ctx) error
	LoginController(c *fiber.Ctx) error
	LogOutController(c *fiber.Ctx) error
	ForgotPasswordController(c *fiber.Ctx) error
	ChangePasswordController(c *fiber.Ctx) error
	RefreshTokenController(c *fiber.Ctx) error
	UpdateProfileAuthorController(c *fiber.Ctx) error
	AuthorProfileController(c *fiber.Ctx) error
	DeleteAuthorController(c *fiber.Ctx) error
}
