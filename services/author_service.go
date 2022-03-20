package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type AuthorService interface {
	RegisterService(ax web.AuthorRegisterRequest) error
	LoginService(ax web.AuthorLoginRequest) (map[string]string, error)
	LogOutService(c *fiber.Ctx)
	ForgotPasswordService(ax web.AuthorForgotPasswordRequest) (map[string]string, error)
	ChangePasswordService(ax web.AuthorChangePassword, c *fiber.Ctx) error
	RefreshTokenService(token web.RefreshTokenRequest) (map[string]string, error)
	UpdateProfileAuthorService(ax *web.AuthorUpdateProfileRequest, c *fiber.Ctx) error
	AuthorProfileService(c *fiber.Ctx) (map[string]string, error)
}
