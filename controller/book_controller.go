package controller

import "github.com/gofiber/fiber/v2"

type BookController interface {
	AddBookController(c *fiber.Ctx) error
	GetBookByIdController(c *fiber.Ctx) error
	GetAllBookController(c *fiber.Ctx) error
	GetMyBookController(c *fiber.Ctx) error
	UpdateMyBookController(c *fiber.Ctx) error
	UpdateCoverBookController(c *fiber.Ctx) error
	DeleteBookController(c *fiber.Ctx) error
}
