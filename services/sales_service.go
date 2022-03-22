package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type SalesService interface {
	AddSalesService(b web.SalesAddRequest, c *fiber.Ctx) error
	GetMySalesByIdService(sales model.Sales, c *fiber.Ctx) (web.SalesByIdResponse, error)
}
