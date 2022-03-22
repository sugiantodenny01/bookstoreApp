package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type BookService interface {
	AddBookService(b web.AddBookRequest, c *fiber.Ctx) error
	GetBookByIdService(id string) (web.BookResponse, error)
	GetAllBookService(page string, max string) ([]web.BookResponse, web.PageInfo, error)
	GetMyBookService(page string, max string, c *fiber.Ctx) ([]web.BookResponse, web.PageInfo, error)
	UpdateBookService(book model.Book, c *fiber.Ctx) error
}
