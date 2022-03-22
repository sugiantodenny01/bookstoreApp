package controller

import "github.com/gofiber/fiber/v2"

type SalesController interface {
	AddSalesController(c *fiber.Ctx) error
}
