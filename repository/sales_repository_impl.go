package repository

import (
	"database/sql"
	"errors"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type SalesRepositoryImpl struct {
}

func NewSalesRepository() SalesRepository {
	return &SalesRepositoryImpl{}
}

func (s *SalesRepositoryImpl) GetInformationBookRepo(tx *sql.Tx, sales web.SalesAddRequest) (web.BookResponse, error) {

	var bookData web.BookResponse
	SQL := "select Price_Per_Unit from book where Book_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, sales.Book_ID)
	errorResultCheckExists := resultCheckExists.Scan(&bookData.Price)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return bookData, err
	}

	return bookData, nil

}

func (s *SalesRepositoryImpl) AddSalesRepository(tx *sql.Tx, sales web.SalesAddRequest) error {

	var bookData string
	SQL := "select Title from book where Book_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, sales.Book_ID)
	errorResultCheckExists := resultCheckExists.Scan(&bookData)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_id_not_found")
		return err
	}

	SQlInsert := `insert into sales(Author_id, Recipient_Name, Recipient_Email, Book_Title, Quantity, Price_Per_Unit, Price_Total)" +
				  values(?,?,?,?,?,?,?)`
	_, err := tx.Exec(SQlInsert, sales.Author_ID, sales.Name, sales.Email, bookData, sales.Quantity, sales.Price_Per_Unit, sales.Price_Total)

	if err != nil {
		err := errors.New("error_internal_server")
		return err
	}

	return nil

}

func (s *SalesRepositoryImpl) GetInformationSalesById(tx *sql.Tx, sales model.Sales) (web.SalesByIdResponse, error) {

	var salesResponse web.SalesByIdResponse
	SQL := "select Sales_ID, Recipient_Name, Recipient_Email, Book_Title Author_ID, Quantity, Price_Per_Unit, Total_Price, Created_Time from sales where Sales_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, sales.Sales_ID)
	errorResultCheckExists := resultCheckExists.Scan(&salesResponse.Sales_ID, salesResponse.Recipient_Name, salesResponse.Recipient_Email, salesResponse.Book_Title, salesResponse.Author_ID, salesResponse.Quantity, salesResponse.Price_Per_Unit, salesResponse.Total_Price, salesResponse.Created_Time)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_id_not_found")
		return salesResponse, err
	}

	return salesResponse, nil
}
