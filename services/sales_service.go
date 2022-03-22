package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type SalesService interface {
	AddSalesService(b web.SalesAddRequest, c *fiber.Ctx) error
}
