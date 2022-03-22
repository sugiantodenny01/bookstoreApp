package services

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/repository"
)

type SalesServiceImpl struct {
	salesRepo repository.SalesRepository
	DB        *sql.DB
}

func NewSalesService(sales repository.SalesRepository, DB *sql.DB) SalesService {
	return &SalesServiceImpl{
		salesRepo: sales,
		DB:        DB,
	}
}

func (s *SalesServiceImpl) AddSalesService(sales web.SalesAddRequest, c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := s.DB.Begin()

	if err != nil {
		tx.Rollback()
		return errors.New("error_internal_server")
	}

	bookInformation := web.SalesAddRequest{
		Book_ID: sales.Book_ID,
	}

	dataBook, err := s.salesRepo.GetInformationBookRepo(tx, bookInformation)

	if err != nil {
		tx.Rollback()
		return err
	}

	totalPrice := float64(dataBook.Price) * float64(sales.Quantity)
	sales.Price_Per_Unit = float64(dataBook.Price)
	sales.Price_Total = totalPrice
	sales.Author_ID = authorId

	err = s.salesRepo.AddSalesRepository(tx, sales)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}
func (s *SalesServiceImpl) GetMySalesByIdService(sales model.Sales, c *fiber.Ctx) (web.SalesByIdResponse, error) {

	var mock web.SalesByIdResponse
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := s.DB.Begin()

	if err != nil {
		tx.Rollback()
		return mock, errors.New("error_internal_server")
	}

	salesInformation := model.Sales{
		Author_ID: authorId,
		Sales_ID:  sales.Sales_ID,
	}

	dataSales, err := s.salesRepo.GetInformationSalesById(tx, salesInformation)
	if err != nil {
		tx.Rollback()
		return mock, err
	}
	tx.Commit()
	return dataSales, nil

}
