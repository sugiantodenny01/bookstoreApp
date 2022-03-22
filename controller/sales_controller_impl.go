package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/services"
	"strconv"
)

type SalesControllerImpl struct {
	SalesService services.SalesService
}

func NewSalesController(sales services.SalesService) SalesController {
	return &SalesControllerImpl{
		SalesService: sales,
	}
}

func (ax *SalesControllerImpl) AddSalesController(c *fiber.Ctx) error {

	var sales web.SalesAddRequest
	err := c.BodyParser(&sales)

	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	err = ax.SalesService.AddSalesService(sales, c)
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

func (ax *SalesControllerImpl) MySalesGetByIdController(c *fiber.Ctx) error {

	idString := c.Params("id")

	if idString == "" {
		err := errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	var sales model.Sales
	idInt, _ := strconv.Atoi(idString)
	sales.Sales_ID = idInt

	data, err := ax.SalesService.GetMySalesByIdService(sales, c)
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

	result := web.SuccessResponse(data)
	return c.JSON(result)

}
