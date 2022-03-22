package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/services"
	"strconv"
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

func (ax *BookControllerImpl) GetBookByIdController(c *fiber.Ctx) error {
	//var book web.GetBookByIdRequest
	idString := c.Params("id")

	data, err := ax.BookService.GetBookByIdService(idString)

	if err != nil {

		if err.Error() == "error_id_not_found" {
			getInformationError := web.ToFailResponse(err, "Error ID yang di supply tidak ada di database")
			return c.JSON(getInformationError)
		} else {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		}

	}

	result := web.SuccessResponse(data)
	return c.JSON(result)

}

func (ax *BookControllerImpl) GetAllBookController(c *fiber.Ctx) error {
	page := c.Query("Page")
	max := c.Query("Limit")

	if page == "" || max == "" {
		err := errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	data, pageInfo, err := ax.BookService.GetAllBookService(page, max)

	if err != nil {

		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error ID yang di supply tidak ada di database")
			return c.JSON(getInformationError)
		} else {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		}

	}

	result := web.SuccessBookByPage(data, pageInfo)
	return c.JSON(result)

}

func (ax *BookControllerImpl) GetMyBookController(c *fiber.Ctx) error {
	page := c.Query("Page")
	max := c.Query("Limit")

	if page == "" || max == "" {
		err := errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	data, pageInfo, err := ax.BookService.GetMyBookService(page, max, c)

	if err != nil {

		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error ID yang di supply tidak ada di database")
			return c.JSON(getInformationError)
		} else {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		}

	}

	result := web.SuccessBookByPage(data, pageInfo)
	return c.JSON(result)

}

func (ax *BookControllerImpl) UpdateMyBookController(c *fiber.Ctx) error {

	var book model.Book

	idString := c.Params("id")
	if idString == "" {
		err := errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	idInt, _ := strconv.Atoi(idString)
	book.Book_ID = idInt

	err := c.BodyParser(&book)

	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	err = ax.BookService.UpdateBookService(book, c)
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

func (ax *BookControllerImpl) UpdateCoverBookController(c *fiber.Ctx) error {

	var book web.UpdateCoverBookRequest

	idString := c.Params("id")
	if idString == "" {
		err := errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	idInt, _ := strconv.Atoi(idString)
	book.Book_ID = idInt

	err := c.BodyParser(&book)

	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	err = ax.BookService.UpdateCoverBookService(book, c)
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

func (ax *BookControllerImpl) DeleteBookController(c *fiber.Ctx) error {
	var book model.Book

	idString := c.Params("id")
	if idString == "" {
		err := errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	idInt, _ := strconv.Atoi(idString)
	book.Book_ID = idInt

	err := ax.BookService.DeleteBookService(book, c)
	if err != nil {

		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		} else if err.Error() == "error_id_not_found" {
			getInformationError := web.ToFailResponse(err, "Error ID yang di supply tidak ada di database")
			return c.JSON(getInformationError)
		} else {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		}

	}

	return c.JSON(map[string]string{"message": "Success"})

}
