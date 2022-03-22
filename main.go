package main

import (
	"github.com/sugiantodenny01/bookstoreApp/app"
	"github.com/sugiantodenny01/bookstoreApp/controller"
	"github.com/sugiantodenny01/bookstoreApp/repository"
	"github.com/sugiantodenny01/bookstoreApp/services"
)

func main() {

	db := app.NewDB()
	conf := app.GetConfig()

	//author
	authorRepository := repository.NewAuthorRepository()
	authorService := services.NewAuthorService(authorRepository, db)
	categoryController := controller.NewAuthorController(authorService)

	//Book
	bookRepository := repository.NewBookRepository()
	bookService := services.NewBookService(bookRepository, db, conf.Port_App)
	bookController := controller.NewBookController(bookService)

	//sales
	salesRepository := repository.NewSalesRepository()
	salesService := services.NewSalesService(salesRepository, db)
	salesController := controller.NewSalesController(salesService)

	router := app.NewRouter(categoryController, bookController, salesController)
	app := router
	app.Listen(conf.Port_App)

}
