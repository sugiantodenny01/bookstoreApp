package repository

import (
	"database/sql"
	"errors"
	"github.com/sugiantodenny01/bookstoreApp/model"
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

//func (b *BookRepositoryImpl) GetBookByIdRepository(tx *sql.Tx, book model.Book) (model.Book, error) {
//
//	var author model.Author
//	SQL := "select Book_ID,Author_ID,Title,Summary,Stock, Price, from author where Author_ID = (?)"
//	resultCheckEmailExists := tx.QueryRow(SQL, book.Author_ID)
//	errorResultCheckEmailExists := resultCheckEmailExists.Scan(&author.Author_ID)
//
//	if errorResultCheckEmailExists == sql.ErrNoRows {
//		err := errors.New("error_author_id_not_found")
//		return author, err
//	}
//
//}
