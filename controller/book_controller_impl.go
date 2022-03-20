package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/services"
)

type BookControllerImpl struct {
	BookService services.BookService
}

func NewBookController(bookService services.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (ax *BookControllerImpl) AddBookController(c *fiber.Ctx) error {

	var book web.AddBookRequest
	err := c.BodyParser(&book)

	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	err = ax.BookService.AddBookService(book, c)
	if err != nil {

		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		} else if err.Error() == "error_author_id_not_found" {
			getInformationError := web.ToFailResponse(err, "Error ID yang di supply tidak ada di database")
			return c.JSON(getInformationError)
		} else {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		}

	}

	return c.JSON(map[string]string{"message": "Success"})

}
