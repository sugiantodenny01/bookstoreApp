package repository

import (
	"database/sql"
	"github.com/sugiantodenny01/bookstoreApp/model"
)

type BookRepository interface {
	AddBookRepository(tx *sql.Tx, book model.Book) error
	//GetBookByIdRepository(tx *sql.Tx, book model.Book) (model.Book, error)
}
