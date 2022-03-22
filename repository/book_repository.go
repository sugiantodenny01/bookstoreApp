package repository

import (
	"database/sql"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type BookRepository interface {
	AddBookRepository(tx *sql.Tx, book model.Book) error
	GetBookByIdRepository(tx *sql.Tx, book model.Book) (web.BookResponse, error)
	GetAllBookRepository(tx *sql.Tx, max int) ([]web.BookResponse, int, error)
	GetMyBookRepository(tx *sql.Tx, authorID int, max int) ([]web.BookResponse, int, error)
	UpdateBookRepository(tx *sql.Tx, book model.Book) (model.Book, error)
}
