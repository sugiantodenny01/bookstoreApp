package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/controller"
	"github.com/sugiantodenny01/bookstoreApp/middleware"
)

func NewRouter(authorController controller.AuthorController, bookController controller.BookController) *fiber.App {
	router := fiber.New()
	conf := GetConfig()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//Author
	router.Post(conf.Host+"/author/register", authorController.RegisterController)
	router.Post(conf.Host+"/author/login", authorController.LoginController)
	router.Post(conf.Host+"/author/logout", middleware.IsAuthenticatedAccessToken, authorController.LogOutController)
	router.Post(conf.Host+"/author/forgot_password", authorController.ForgotPasswordController)
	router.Post(conf.Host+"/author/change_password", middleware.IsAuthenticatedAccessToken, authorController.ChangePasswordController)
	router.Post(conf.Host+"/author/refresh_token", authorController.RefreshTokenController)
	router.Put(conf.Host+"/author/update", middleware.IsAuthenticatedAccessToken, authorController.UpdateProfileAuthorController)
	router.Get(conf.Host+"/author/get_my_profile", middleware.IsAuthenticatedAccessToken, authorController.AuthorProfileController)
	router.Static(conf.Host+"/assets/", "./assets")

	//Book
	router.Post(conf.Host+"/book/add", middleware.IsAuthenticatedAccessToken, bookController.AddBookController)

	return router
}
