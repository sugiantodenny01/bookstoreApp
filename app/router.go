package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugiantodenny01/bookstoreApp/controller"
	"github.com/sugiantodenny01/bookstoreApp/middleware"
)

func NewRouter(authorController controller.AuthorController, bookController controller.BookController, salesController controller.SalesController) *fiber.App {
	router := fiber.New()
	//conf := GetConfig()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//load assets
	router.Static("/assets/", "./assets")

	//Author
	router.Post("/author/register", authorController.RegisterController)
	router.Post("/author/login", authorController.LoginController)
	router.Post("/author/logout", middleware.IsAuthenticatedAccessToken, authorController.LogOutController)
	router.Post("/author/forgot_password", authorController.ForgotPasswordController)
	router.Post("/author/change_password", middleware.IsAuthenticatedAccessToken, authorController.ChangePasswordController)
	router.Post("/author/refresh_token", authorController.RefreshTokenController)
	router.Put("/author/update", middleware.IsAuthenticatedAccessToken, authorController.UpdateProfileAuthorController)
	router.Get("/author/get_my_profile", middleware.IsAuthenticatedAccessToken, authorController.AuthorProfileController)
	router.Delete("/author/delete", middleware.IsAuthenticatedAccessToken, authorController.DeleteAuthorController)

	//Book
	router.Post("/book/add", middleware.IsAuthenticatedAccessToken, bookController.AddBookController)
	router.Get("/book/get", bookController.GetAllBookController)
	router.Get("/book/get_my_book", middleware.IsAuthenticatedAccessToken, bookController.GetMyBookController)
	router.Get("/book/get/:id", bookController.GetBookByIdController)
	router.Post("/book/update/:id", middleware.IsAuthenticatedAccessToken, bookController.UpdateMyBookController)
	router.Post("/book/update_cover/:id", middleware.IsAuthenticatedAccessToken, bookController.UpdateCoverBookController)
	router.Delete("/book/delete/:id", middleware.IsAuthenticatedAccessToken, bookController.DeleteBookController)

	//sales
	router.Post("/sales/add", middleware.IsAuthenticatedAccessToken, salesController.AddSalesController)
	router.Get("/sales/get/:id", middleware.IsAuthenticatedAccessToken, salesController.MySalesGetByIdController)

	return router
}
