package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type BookService interface {
	AddBookService(b web.AddBookRequest, c *fiber.Ctx) error
}
