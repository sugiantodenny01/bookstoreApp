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
	//router.Post(conf.Host+"/upload", uploadImage)
	router.Static(conf.Host+"/assets/", "./assets")

	//Book
	router.Post(conf.Host+"/book/add", middleware.IsAuthenticatedAccessToken, bookController.AddBookController)

	return router
}

//func uploadImage(c *fiber.Ctx) error {
//
//	type upload struct {
//		Image string `json:"Image"`
//	}
//
//	var img upload
//	var f *os.File
//
//	err := c.BodyParser(&img)
//	if err != nil {
//		err = errors.New("error_param")
//		getInformationError := web.ToFailResponse(err, "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai")
//		return c.JSON(getInformationError)
//	}
//
//	fmt.Println(img)
//	unbased, _ := base64.StdEncoding.DecodeString(string(img.Image))
//	jpgI, err := jpeg.Decode(bytes.NewReader(unbased))
//
//	if err != nil {
//		return err
//	}
//	path, _ := os.Getwd()
//	newpath := path + "/assets/" + "\\image"
//	f, _ = os.OpenFile(newpath+".jpg", os.O_WRONLY|os.O_CREATE, 0777)
//	png.Encode(f, jpgI)
//	fmt.Println("jpg generated")
//
//	return nil
//
//}
