package controller

import "github.com/gofiber/fiber/v2"

type BookController interface {
	AddBookController(c *fiber.Ctx) error
}
