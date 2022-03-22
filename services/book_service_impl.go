package services

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/repository"
	"image/jpeg"
	"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	port           string
}

func NewBookService(book repository.BookRepository, DB *sql.DB, p string) BookService {
	return &BookServiceImpl{
		BookRepository: book,
		DB:             DB,
		port:           p,
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
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

func (service *BookServiceImpl) GetBookByIdService(request string) (web.BookResponse, error) {

	tx, err := service.DB.Begin()
	var mock web.BookResponse

	if err != nil {
		return mock, errors.New("error_internal_server")
	}

	idInt, err := strconv.Atoi(request)

	bookInformation := model.Book{
		Book_ID: idInt,
	}

	data, err := service.BookRepository.GetBookByIdRepository(tx, bookInformation)

	locationImageUrl := service.port + data.Cover_URL
	data.Cover_URL = locationImageUrl

	if err != nil {
		tx.Rollback()
		return mock, err
	}
	tx.Commit()

	return data, nil

}

func (service *BookServiceImpl) GetAllBookService(page string, max string) ([]web.BookResponse, web.PageInfo, error) {

	tx, err := service.DB.Begin()
	var pageInfo web.PageInfo

	if err != nil {
		return nil, pageInfo, errors.New("error_internal_server")
		tx.Rollback()
	}

	pageStatus, err := strconv.Atoi(page)
	maxStatus, err := strconv.Atoi(max)

	data, countValData, err := service.BookRepository.GetAllBookRepository(tx, maxStatus)

	for i := range data {
		data[i].Cover_URL = service.port + data[i].Cover_URL
	}

	maxPage := math.Ceil(float64(countValData) / float64(maxStatus))

	pageInfo.Current_Page = pageStatus
	pageInfo.Max_Data_Per_Page = maxStatus
	pageInfo.Max_Page = int(maxPage)
	pageInfo.Total_All_Data = countValData
	tx.Commit()

	return data, pageInfo, nil
}

func (service *BookServiceImpl) GetMyBookService(page string, max string, c *fiber.Ctx) ([]web.BookResponse, web.PageInfo, error) {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := service.DB.Begin()
	var pageInfo web.PageInfo

	if err != nil {
		return nil, pageInfo, errors.New("error_internal_server")
		tx.Rollback()

	}

	pageStatus, _ := strconv.Atoi(page)
	maxStatus, _ := strconv.Atoi(max)

	data, countValData, err := service.BookRepository.GetMyBookRepository(tx, authorId, maxStatus)

	if err != nil {
		return nil, pageInfo, errors.New("error_internal_server")
		tx.Rollback()

	}

	for i := range data {
		data[i].Cover_URL = service.port + data[i].Cover_URL
	}

	maxPage := math.Ceil(float64(countValData) / float64(maxStatus))

	pageInfo.Current_Page = pageStatus
	pageInfo.Max_Data_Per_Page = maxStatus
	pageInfo.Max_Page = int(maxPage)
	pageInfo.Total_All_Data = countValData
	tx.Commit()

	return data, pageInfo, nil
}

func (service *BookServiceImpl) UpdateBookService(book model.Book, c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := service.DB.Begin()

	if err != nil {
		tx.Rollback()
		return err
	}

	book, err = service.BookRepository.UpdateBookRepository(tx, book)

	if err != nil {
		tx.Rollback()
		return err
	}

	if authorId != book.Author_ID {
		tx.Rollback()
		return errors.New("error_internal_server")
	}

	tx.Commit()
	return nil

}

func (service *BookServiceImpl) UpdateCoverBookService(request web.UpdateCoverBookRequest, c *fiber.Ctx) error {
	var f *os.File

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := service.DB.Begin()

	if err != nil {
		tx.Rollback()
		return err
	}

	if (request.Image_Extension != "jpg") && (request.Image_Extension != "jpeg") && (request.Image_Extension != "png") {
		err = errors.New("error_type_support")
		return err
	}

	unbased, _ := base64.StdEncoding.DecodeString(string(request.Cover_Image))
	nameImageAndPath := ""
	randomName := randomString(10)

	if request.Image_Extension == "jpg" || request.Image_Extension == "jpeg" {

		dataGambar, err := jpeg.Decode(bytes.NewReader(unbased))

		if err != nil {
			return err
		}

		path, _ := os.Getwd()
		newpath := path + "/assets/" + randomName
		f, _ = os.OpenFile(newpath+".jpg", os.O_WRONLY|os.O_CREATE, 0777)
		jpeg.Encode(f, dataGambar, &jpeg.Options{Quality: 75})
		nameImageAndPath = "/assets/" + randomName + ".jpg"

	} else if request.Image_Extension == "png" {

		dataGambar, err := png.Decode(bytes.NewReader(unbased))

		if err != nil {
			return err
		}

		path, _ := os.Getwd()
		newpath := path + "/assets/" + randomName
		f, _ = os.OpenFile(newpath+".png", os.O_WRONLY|os.O_CREATE, 0777)
		png.Encode(f, dataGambar)
		nameImageAndPath = "/assets/" + randomName + ".png"

	}

	bookInformation := model.Book{
		Book_ID:   request.Book_ID,
		Author_ID: authorId,
		Cover_URL: nameImageAndPath,
	}

	book, err := service.BookRepository.UpdateCoverBookRepository(tx, bookInformation)

	if err != nil {
		path, _ := os.Getwd()
		fileImageOld := strings.Replace(nameImageAndPath, "/", `\`, -1)
		pathImage := path + fileImageOld
		err := os.Remove(pathImage)

		if err != nil {
			return err
		}

		tx.Rollback()
		return err
	}

	if authorId != book.Author_ID {
		tx.Rollback()
		return errors.New("error_internal_server")
	}

	tx.Commit()
	return nil

}

func (service *BookServiceImpl) DeleteBookService(book model.Book, c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := service.DB.Begin()

	if err != nil {
		tx.Rollback()
		return errors.New("error_internal_server")
	}

	bookInformation := model.Book{
		Book_ID: book.Book_ID,
	}

	data, err := service.BookRepository.GetBookByIdRepository(tx, bookInformation)

	if authorId != data.Author_ID {
		tx.Rollback()
		return errors.New("error_id_not_found")
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	err = service.BookRepository.DeleteBookRepository(tx, book)

	if err != nil {
		tx.Rollback()
		return err
	}

	if authorId != book.Author_ID {
		tx.Rollback()
		return errors.New("error_internal_server")
	}

	tx.Commit()
	return nil
}
