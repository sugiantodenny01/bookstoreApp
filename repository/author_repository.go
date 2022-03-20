package repository

import (
	"database/sql"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
)

type AuthorRepository interface {
	RegisterRepository(tx *sql.Tx, author model.Author) error
	LoginRepository(tx *sql.Tx, author model.Author) (model.Author, error)
	ForgotPasswordRepository(tx *sql.Tx, author model.Author) (model.Author, error)
	ChangePasswordRepository(tx *sql.Tx, author model.Author, newPassword string) error
	UpdateAuthorRepository(tx *sql.Tx, author web.AuthorUpdateProfileRequest) error
	AuthorProfileRepository(tx *sql.Tx, author model.Author) (web.AuthorProfileResponse, error)
}
