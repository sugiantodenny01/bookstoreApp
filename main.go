package main

import (
	"github.com/sugiantodenny01/bookstoreApp/app"
	"github.com/sugiantodenny01/bookstoreApp/controller"
	"github.com/sugiantodenny01/bookstoreApp/repository"
	"github.com/sugiantodenny01/bookstoreApp/services"
)

func main() {

	db := app.NewDB()
	//author
	authorRepository := repository.NewAuthorRepository()
	authorService := services.NewAuthorService(authorRepository, db)
	categoryController := controller.NewAuthorController(authorService)

	//Book
	bookRepository := repository.NewBookRepository()
	bookService := services.NewBookService(bookRepository, db)
	bookController := controller.NewBookController(bookService)

	router := app.NewRouter(categoryController, bookController)
	app := router
	app.Listen(":12345")

}
