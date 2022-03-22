package repository

import (
	"database/sql"
	"errors"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"os"
	"strconv"
	"strings"
	"time"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (b *BookRepositoryImpl) AddBookRepository(tx *sql.Tx, book model.Book) error {

	var author model.Author
	SQL := "select Author_ID from author where Author_ID = (?)"
	resultCheckEmailExists := tx.QueryRow(SQL, book.Author_ID)
	errorResultCheckEmailExists := resultCheckEmailExists.Scan(&author.Author_ID)

	if errorResultCheckEmailExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return err
	}

	SQLInsert := "insert into book(Author_ID, Title, Summary, Stock, Price, Cover_URL, Created_Time) Values (?,?,?,?,?,?,?)"
	currentTime := time.Now()
	_, err := tx.Exec(SQLInsert, book.Author_ID, book.Title, book.Summary, book.Stock, book.Price, book.Cover_URL, currentTime)

	if err != nil {
		err := errors.New("error_internal_server")
		return err
	}

	return nil

}

func (b *BookRepositoryImpl) GetBookByIdRepository(tx *sql.Tx, book model.Book) (web.BookResponse, error) {

	var bookResponse web.BookResponse

	SQL := "select Book_ID,Author_ID,Title,Summary,Stock, Price, Cover_URL from book where Book_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, book.Book_ID)
	errorResultCheckExists := resultCheckExists.Scan(&bookResponse.Book_ID, &bookResponse.Author_ID, &bookResponse.Title, &bookResponse.Summary, &bookResponse.Stock, &bookResponse.Price, &bookResponse.Cover_URL)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_id_not_found")
		return bookResponse, err
	}

	SQLInformationUser := "select Pen_Name from author where Author_ID = (?)"
	resultCheckUserExists := tx.QueryRow(SQLInformationUser, bookResponse.Author_ID)
	errorResultCheckUserExists := resultCheckUserExists.Scan(&bookResponse.Author_Pen_name)

	if errorResultCheckUserExists == sql.ErrNoRows {
		err := errors.New("error_id_not_found")
		return bookResponse, err
	}

	return bookResponse, nil

}

func (b *BookRepositoryImpl) GetAllBookRepository(tx *sql.Tx, max int) ([]web.BookResponse, int, error) {

	var bookResponse web.BookResponse
	var arrObj []web.BookResponse

	SQL := `select book.Book_ID, book.Author_ID, book.Title, book.Summary, book.Stock, book.Price, book.Cover_URL, author.Pen_Name
			from (select * from book limit ? ) as book join author on book.Author_ID = author.Author_ID 
			where author.Is_Disabled = false`
	resultCheckExists, err := tx.Query(SQL, max)

	if err != nil {
		err := errors.New("error_internal_server")
		return nil, 0, err
	}

	for resultCheckExists.Next() {

		err = resultCheckExists.Scan(&bookResponse.Book_ID, &bookResponse.Author_ID, &bookResponse.Title, &bookResponse.Summary, &bookResponse.Stock, &bookResponse.Price, &bookResponse.Cover_URL, &bookResponse.Author_Pen_name)
		if err != nil {
			return nil, 0, err
		}
		arrObj = append(arrObj, bookResponse)

	}

	var countValData int
	SQLCountData := "select count(book.Book_ID) from book join author on book.Author_ID = author.Author_ID where author.Is_Disabled = false"
	resultCountData := tx.QueryRow(SQLCountData)
	errorResultCountData := resultCountData.Scan(&countValData)

	if errorResultCountData == sql.ErrNoRows {
		err := errors.New("error_internal_server")
		return nil, 0, err
	}

	return arrObj, countValData, nil

}

func (b *BookRepositoryImpl) GetMyBookRepository(tx *sql.Tx, authorID int, max int) ([]web.BookResponse, int, error) {

	var bookResponse web.BookResponse
	var arrObj []web.BookResponse

	SQL := `select book.Book_ID, book.Author_ID, book.Title, book.Summary, book.Stock, book.Price, book.Cover_URL, author.Pen_Name
			from  book join author on book.Author_ID = author.Author_ID 
			where author.Is_Disabled = false and author.Author_ID = (?) limit ?`
	resultCheckExists, err := tx.Query(SQL, authorID, max)

	if err != nil {
		err := errors.New("error_internal_server")
		return nil, 0, err
	}

	for resultCheckExists.Next() {

		err = resultCheckExists.Scan(&bookResponse.Book_ID, &bookResponse.Author_ID, &bookResponse.Title, &bookResponse.Summary, &bookResponse.Stock, &bookResponse.Price, &bookResponse.Cover_URL, &bookResponse.Author_Pen_name)
		if err != nil {
			return nil, 0, err
		}
		arrObj = append(arrObj, bookResponse)

	}

	var countValData int
	SQLCountData := "select count(book.Book_ID) from book join author on book.Author_ID = author.Author_ID where author.Is_Disabled = false and author.Author_ID = (?)"
	resultCountData := tx.QueryRow(SQLCountData, authorID)
	errorResultCountData := resultCountData.Scan(&countValData)

	if errorResultCountData == sql.ErrNoRows {
		err := errors.New("error_internal_server")
		return nil, 0, err
	}

	return arrObj, countValData, nil

}

func (b *BookRepositoryImpl) UpdateBookRepository(tx *sql.Tx, book model.Book) (model.Book, error) {

	var mock model.Book

	SQL := "select Author_ID from book where Book_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, book.Book_ID)
	errorResultCheckExists := resultCheckExists.Scan(&book.Author_ID)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return mock, err
	}

	SQLUpdate := "update book set Title = (?), Summary = (?), Price = (?), Stock (?) where Book_ID = (?)"
	_, err := tx.Exec(SQLUpdate, book.Title, book.Summary, book.Price, book.Stock, book.Book_ID)

	if err != nil {
		err := errors.New("error_internal_server")
		return mock, err
	}

	return book, nil

}

func (b *BookRepositoryImpl) UpdateCoverBookRepository(tx *sql.Tx, book model.Book) (model.Book, error) {

	var mock model.Book

	SQL := "select Author_ID, Cover_URL from book where Book_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, book.Book_ID)
	errorResultCheckExists := resultCheckExists.Scan(&book.Author_ID, &mock.Cover_URL)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return mock, err
	}

	//remove old image
	path, _ := os.Getwd()
	fileImageOld := strings.Replace(mock.Cover_URL, "/", `\`, -1)
	pathImage := path + fileImageOld
	err := os.Remove(pathImage)

	if err != nil {
		return mock, err
	}

	SQLUpdate := "update book set Cover_URL = (?) where Book_ID = (?)"
	_, err = tx.Exec(SQLUpdate, book.Cover_URL, book.Book_ID)

	if err != nil {
		err := errors.New("error_internal_server")
		return mock, err
	}

	return book, nil

}

func (b *BookRepositoryImpl) DeleteBookRepository(tx *sql.Tx, book model.Book) error {

	SQL := "delete from book where Book_ID = (?)"
	bookIdString := strconv.Itoa(book.Book_ID)
	_, err := tx.Query(SQL, bookIdString)

	if err != nil {
		err := errors.New("error_id_not_found")
		return err
	}

	return nil

}
