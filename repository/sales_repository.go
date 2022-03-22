package repository

import (
	"database/sql"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type SalesRepository interface {
	GetInformationBookRepo(tx *sql.Tx, sales web.SalesAddRequest) (web.BookResponse, error)
	AddSalesRepository(tx *sql.Tx, sales web.SalesAddRequest) error
	//GetInformationSalesById(tx *sql.Tx, sales web.SalesAddRequest) error
}
