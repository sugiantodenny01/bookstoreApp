package services

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/repository"
	"image/jpeg"
	"image/png"
	"os"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
}

func NewBookService(book repository.BookRepository, DB *sql.DB) BookService {
	return &BookServiceImpl{
		BookRepository: book,
		DB:             DB,
	}
}

func (service *BookServiceImpl) AddBookService(request web.AddBookRequest, c *fiber.Ctx) error {
	var f *os.File
	tx, err := service.DB.Begin()

	if err != nil {
		return errors.New("error_internal_server")
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	if (request.Image_Extension != "jpg") && (request.Image_Extension != "jpeg") && (request.Image_Extension != "png") {
		err = errors.New("error_type_support")
		return err
	}

	unbased, _ := base64.StdEncoding.DecodeString(string(request.Cover_Image))
	nameImageAndPath := ""

	if request.Image_Extension == "jpg" || request.Image_Extension == "jpeg" {

		dataGambar, err := jpeg.Decode(bytes.NewReader(unbased))

		if err != nil {
			return err
		}

		path, _ := os.Getwd()
		newpath := path + "/assets/" + request.Title
		f, _ = os.OpenFile(newpath+".jpg", os.O_WRONLY|os.O_CREATE, 0777)
		jpeg.Encode(f, dataGambar, &jpeg.Options{Quality: 75})
		nameImageAndPath = "/assets/" + request.Title + ".jpg"

	} else if request.Image_Extension == "png" {

		dataGambar, err := png.Decode(bytes.NewReader(unbased))

		if err != nil {
			return err
		}

		path, _ := os.Getwd()
		newpath := path + "/assets/" + request.Title
		f, _ = os.OpenFile(newpath+".png", os.O_WRONLY|os.O_CREATE, 0777)
		png.Encode(f, dataGambar)
		nameImageAndPath = "/assets/" + request.Title + ".png"

	}

	authorInformation := model.Book{
		Author_ID: authorId,
		Title:     request.Title,
		Summary:   request.Summary,
		Stock:     request.Stock,
		Price:     request.Price,
		Cover_URL: nameImageAndPath,
	}

	err = service.BookRepository.AddBookRepository(tx, authorInformation)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}
