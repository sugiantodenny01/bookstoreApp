package controller

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/services"
)

type AuthorControllerImpl struct {
	AuthorService services.AuthorService
}

func NewAuthorController(AuthorService services.AuthorService) AuthorController {
	return &AuthorControllerImpl{
		AuthorService: AuthorService,
	}
}

func (ax *AuthorControllerImpl) RegisterController(c *fiber.Ctx) error {

	var author web.AuthorRegisterRequest

	err := c.BodyParser(&author)
	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	err = ax.AuthorService.RegisterService(author)
	if err != nil {

		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		} else {
			getInformationError := web.ToFailResponse(err, "Email sudah ada pada data Author dan tidak bisa di registrasi lagi")
			return c.JSON(getInformationError)
		}

	}

	return c.JSON(map[string]string{"message": "Success"})

}

func (ax *AuthorControllerImpl) LoginController(c *fiber.Ctx) error {

	var author web.AuthorLoginRequest

	err := c.BodyParser(&author)
	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	data, err := ax.AuthorService.LoginService(author)
	if err != nil {
		getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
		return c.JSON(getInformationError)
	}

	result := web.SuccessResponse(data)
	return c.JSON(result)

}
func (ax *AuthorControllerImpl) LogOutController(c *fiber.Ctx) error {
	//Checking error dihandle middleware
	ax.AuthorService.LogOutService(c)
	return c.JSON(map[string]string{"message": "Success"})
}

func (ax *AuthorControllerImpl) ForgotPasswordController(c *fiber.Ctx) error {

	var author web.AuthorForgotPasswordRequest

	err := c.BodyParser(&author)
	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	data, err := ax.AuthorService.ForgotPasswordService(author)
	if err != nil {
		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		} else {
			getInformationError := web.ToFailResponse(err, "Email tidak ditemukan di data Author")
			return c.JSON(getInformationError)
		}
	}

	result := web.SuccessResponse(data)
	return c.JSON(result)

}

func (ax *AuthorControllerImpl) ChangePasswordController(c *fiber.Ctx) error {

	var author web.AuthorChangePassword

	err := c.BodyParser(&author)
	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	err = ax.AuthorService.ChangePasswordService(author, c)
	if err != nil {
		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)

		} else if err.Error() == "error_author_id_not_found" {
			getInformationError := web.ToFailResponse(err, "Error ID yang di supply tidak ada di database")
			return c.JSON(getInformationError)

		} else if err.Error() == "error_invalid_password" {
			getInformationError := web.ToFailResponse(err, "Password tidak sesuai")
			return c.JSON(getInformationError)
		}
	}

	return c.JSON(map[string]string{"message": "Success"})
}

func (ax *AuthorControllerImpl) RefreshTokenController(c *fiber.Ctx) error {

	var token web.RefreshTokenRequest

	err := c.BodyParser(&token)
	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}

	data, err := ax.AuthorService.RefreshTokenService(token)

	if err != nil {

		if err.Error() == "error_internal_server" {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)

		} else if err.Error() == "signature is invalid" {
			setError := errors.New("error_internal_server")
			getInformationError := web.ToFailResponse(setError, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)

		} else if err.Error() == "Token is expired" {
			getInformationError := web.ToFailResponse(err, "Refresh Token yang di supply sudah kadaluarsa")
			return c.JSON(getInformationError)

		} else {
			getInformationError := web.ToFailResponse(err, "Error Selain yang tercantum di sini")
			return c.JSON(getInformationError)
		}

	}

	result := web.SuccessResponse(data)
	return c.JSON(result)
}

func (ax *AuthorControllerImpl) UpdateProfileAuthorController(c *fiber.Ctx) error {
	var author web.AuthorUpdateProfileRequest
	fmt.Println("hello controller")

	err := c.BodyParser(&author)
	if err != nil {
		err = errors.New("error_param")
		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
		return c.JSON(getInformationError)
	}
	fmt.Println(author)

	err = ax.AuthorService.UpdateProfileAuthorService(&author, c)
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

func (ax *AuthorControllerImpl) AuthorProfileController(c *fiber.Ctx) error {

	data, err := ax.AuthorService.AuthorProfileService(c)

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

	result := web.SuccessResponse(data)
	return c.JSON(result)
}

func (ax *AuthorControllerImpl) DeleteAuthorController(c *fiber.Ctx) error {

	err := ax.AuthorService.DeleteAuthorService(c)

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
