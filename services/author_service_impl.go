package services

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sugiantodenny01/bookstoreApp/model"
	"github.com/sugiantodenny01/bookstoreApp/model/web"
	"github.com/sugiantodenny01/bookstoreApp/repository"
	"strconv"
	"time"
)

type CategoryServiceImpl struct {
	AuthorRepository repository.AuthorRepository
	DB               *sql.DB
}

func NewAuthorService(author repository.AuthorRepository, DB *sql.DB) AuthorService {
	return &CategoryServiceImpl{
		AuthorRepository: author,
		DB:               DB,
	}
}
func (service *CategoryServiceImpl) RegisterService(request web.AuthorRegisterRequest) error {

	tx, err := service.DB.Begin()

	if err != nil {
		return errors.New("error_internal_server")
	}

	authorInformation := model.Author{
		Name:     request.Name,
		Pen_Name: request.Pen_Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = service.AuthorRepository.RegisterRepository(tx, authorInformation)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func (service *CategoryServiceImpl) LoginService(request web.AuthorLoginRequest) (map[string]string, error) {

	tx, err := service.DB.Begin()

	if err != nil {
		return nil, errors.New("error_internal_server")
	}
	authorInformation := model.Author{
		Email:    request.Email,
		Password: request.Password,
	}

	author, err := service.AuthorRepository.LoginRepository(tx, authorInformation)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	//set Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Author_ID"] = author.Author_ID
	claims["Name"] = author.Name
	claims["Pen_Name"] = author.Pen_Name
	claims["Email"] = author.Email
	claims["exp"] = time.Now().Add(time.Second * 30).Unix()
	t, err := token.SignedString([]byte("accessToken"))
	if err != nil {
		err = errors.New("error_internal_server")
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["Author_ID"] = author.Author_ID
	rtClaims["Name"] = author.Name
	rtClaims["Pen_Name"] = author.Pen_Name
	rtClaims["Email"] = author.Email
	rtClaims["exp"] = time.Now().Add(time.Second * 30).Unix()

	rt, err := refreshToken.SignedString([]byte("refreshToken"))
	if err != nil {
		err = errors.New("error_internal_server")
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil

}

func (service *CategoryServiceImpl) LogOutService(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Second * 1).Unix()
}

func (service *CategoryServiceImpl) ForgotPasswordService(request web.AuthorForgotPasswordRequest) (map[string]string, error) {
	tx, err := service.DB.Begin()

	if err != nil {
		return nil, errors.New("error_internal_server")
	}
	authorInformation := model.Author{
		Email: request.Email,
	}

	author, err := service.AuthorRepository.ForgotPasswordRepository(tx, authorInformation)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return map[string]string{
		"New_Password": author.Password,
	}, nil

}

func (service *CategoryServiceImpl) ChangePasswordService(request web.AuthorChangePassword, c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := service.DB.Begin()

	if err != nil {
		return errors.New("error_internal_server")
	}

	authorInformation := model.Author{
		Author_ID: authorId,
		Password:  request.Old_Password,
	}

	err = service.AuthorRepository.ChangePasswordRepository(tx, authorInformation, request.New_Password)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func (service *CategoryServiceImpl) RefreshTokenService(tokenReq web.RefreshTokenRequest) (map[string]string, error) {

	token, err := jwt.Parse(tokenReq.Refresh_Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errorInternal := errors.New("error_internal_server")
			return nil, errorInternal
		}
		return []byte("refreshToken"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		authorId := int(claims["Author_ID"].(float64))
		authorName := claims["Name"]
		authorPenName := claims["Pen_Name"]
		authorEmail := claims["Email"]

		//set Token
		NewToken := jwt.New(jwt.SigningMethodHS256)
		newClaims := NewToken.Claims.(jwt.MapClaims)
		newClaims["Author_ID"] = authorId
		newClaims["Name"] = authorName
		newClaims["Pen_Name"] = authorPenName
		newClaims["Email"] = authorEmail
		newClaims["exp"] = time.Now().Add(time.Second * 40).Unix()
		t, err := NewToken.SignedString([]byte("accessToken"))
		if err != nil {
			err = errors.New("error_internal_server")
			return nil, err
		}

		refreshToken := jwt.New(jwt.SigningMethodHS256)
		rtClaims := refreshToken.Claims.(jwt.MapClaims)
		rtClaims["Author_ID"] = authorId
		rtClaims["Name"] = authorName
		rtClaims["Pen_Name"] = authorPenName
		rtClaims["Email"] = authorEmail
		rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		rt, err := refreshToken.SignedString([]byte("refreshToken"))
		if err != nil {
			err = errors.New("error_internal_server")
			return nil, err
		}

		return map[string]string{
			"access_token":  t,
			"refresh_token": rt,
		}, nil

	}

	//token invalid
	err = errors.New("error_refresh_token_invalid")
	return nil, err

}

func (service *CategoryServiceImpl) UpdateProfileAuthorService(request *web.AuthorUpdateProfileRequest, c *fiber.Ctx) error {

	tx, err := service.DB.Begin()

	if err != nil {
		return errors.New("error_internal_server")
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))
	request.Author_ID = authorId

	err = service.AuthorRepository.UpdateAuthorRepository(tx, *request)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (service *CategoryServiceImpl) AuthorProfileService(c *fiber.Ctx) (map[string]string, error) {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := int(claims["Author_ID"].(float64))

	tx, err := service.DB.Begin()

	if err != nil {
		return nil, errors.New("error_internal_server")
	}

	authorInformation := model.Author{
		Author_ID: authorId,
	}

	result, err := service.AuthorRepository.AuthorProfileRepository(tx, authorInformation)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	authorIdTostring := strconv.Itoa(result.Author_ID)

	mappingData := map[string]string{
		"Author_ID": authorIdTostring,
		"Name":      result.Name,
		"Pen_Name":  result.Pen_Name,
		"Email":     result.Email,
	}

	return mappingData, nil

}
