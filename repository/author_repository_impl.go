package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"math/rand"
	"time"
)

type AuthorRepositoryImpl struct {
}

func NewAuthorRepository() AuthorRepository {
	return &AuthorRepositoryImpl{}
}

func (repo *AuthorRepositoryImpl) RegisterRepository(tx *sql.Tx, ar model.Author) error {

	var author model.Author

	SQL := "select * from author where Email = (?)"
	resultCheckExists := tx.QueryRow(SQL, ar.Email)
	errorResultCheckExists := resultCheckExists.Scan(author)

	if errorResultCheckExists == sql.ErrNoRows {

		SQLInsert := "insert into author(Name, Pen_Name, Email,Password, Created_Time) Values (?,?,?,?,?)"
		currentTime := time.Now()
		_, err := tx.Exec(SQLInsert, ar.Email, ar.Pen_Name, ar.Email, ar.Password, currentTime)

		if err != nil {
			err := errors.New("error_internal_server")
			return err
		}

		return err

	}

	err := errors.New("error_email_duplicate")
	return err
}

func (repo *AuthorRepositoryImpl) LoginRepository(tx *sql.Tx, ar model.Author) (model.Author, error) {
	var author model.Author

	SQL := "select * from author where Email = (?) and Is_Disabled = (?) "
	resultCheckExists := tx.QueryRow(SQL, ar.Email, 0)
	errorResultCheckExists := resultCheckExists.Scan(&author.Author_ID, &author.Name, &author.Pen_Name, &author.Email, &author.Password, &author.Is_Disabled, &author.Created_Time)

	if errorResultCheckExists != nil {
		fmt.Println(errorResultCheckExists)

	}

	if errorResultCheckExists == sql.ErrNoRows {

		err := errors.New("error_email_not_found")
		return author, err

	}

	if author.Password != ar.Password {
		err := errors.New("error_invalid_password")
		return author, err
	}

	return author, nil

}

func (repo *AuthorRepositoryImpl) ForgotPasswordRepository(tx *sql.Tx, ar model.Author) (model.Author, error) {

	var author model.Author

	SQL := "select Email from author where Email = (?)"
	resultCheckEmailExists := tx.QueryRow(SQL, ar.Email)
	errorResultCheckEmailExists := resultCheckEmailExists.Scan(author.Email)

	if errorResultCheckEmailExists == sql.ErrNoRows {
		err := errors.New("error_email_not_found")
		return author, err
	}

	rand.Seed(time.Now().UnixNano())
	randomConfig := make([]byte, 10)
	rand.Read(randomConfig)
	randomPassword := fmt.Sprintf("%x", randomConfig)[:10]

	SQLUpdate := "update author set password = ? where email = ?"
	_, err := tx.Exec(SQLUpdate, randomPassword, ar.Email)
	if err != nil {
		err := errors.New("error_internal_server")
		return author, err
	}

	author.Password = randomPassword
	return author, nil

}

func (repo *AuthorRepositoryImpl) ChangePasswordRepository(tx *sql.Tx, ar model.Author, newPassword string) error {
	var author model.Author

	SQL := "select Password from author where Author_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, ar.Author_ID)
	errorResultCheckExists := resultCheckExists.Scan(&author.Password)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return err
	}

	if author.Password != ar.Password {
		err := errors.New("error_invalid_password")
		return err
	}

	SQLUpdate := "update author set password = ? where Author_ID = ?"
	_, err := tx.Exec(SQLUpdate, newPassword, ar.Author_ID)

	if err != nil {
		err := errors.New("error_internal_server")
		return err
	}

	return nil

}

func (repo *AuthorRepositoryImpl) UpdateAuthorRepository(tx *sql.Tx, ar web.AuthorUpdateProfileRequest) error {
	var author model.Author
	SQL := "select Author_ID from author where Author_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, ar.Author_ID)
	errorResultCheckExists := resultCheckExists.Scan(&author.Author_ID)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return err
	}

	SQLUpdate := "update author set Name = (?), Pen_Name = (?) where Author_ID = (?)"
	_, err := tx.Exec(SQLUpdate, ar.Name, ar.Pen_Name, ar.Author_ID)

	if err != nil {
		err := errors.New("error_internal_server")
		return err
	}

	return nil

}

func (repo *AuthorRepositoryImpl) AuthorProfileRepository(tx *sql.Tx, ar model.Author) (web.AuthorProfileResponse, error) {

	var author web.AuthorProfileResponse
	SQL := "select Author_ID, NAME , Pen_Name, Email from author where Author_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, ar.Author_ID)
	errorResultCheckExists := resultCheckExists.Scan(&author.Author_ID, &author.Name, &author.Pen_Name, &author.Email)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return author, err
	}

	return author, nil

}

func (repo *AuthorRepositoryImpl) DeleteAuthorRepository(tx *sql.Tx, ar model.Author) error {

	var author model.Author
	SQL := "select Author_ID from author where Author_ID = (?)"
	resultCheckExists := tx.QueryRow(SQL, ar.Author_ID)
	errorResultCheckExists := resultCheckExists.Scan(&author.Author_ID)

	if errorResultCheckExists == sql.ErrNoRows {
		err := errors.New("error_author_id_not_found")
		return err
	}

	SQLUpdate := "update author set Is_Disabled = (?) where Author_ID = (?)"
	_, err := tx.Exec(SQLUpdate, ar.Is_Disabled, ar.Author_ID)

	if err != nil {
		err := errors.New("error_internal_server")
		return err
	}

	return nil

}
