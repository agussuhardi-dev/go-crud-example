package main

import (
	"agussuhardi/go-crud/application"
	"agussuhardi/go-crud/controller"
	"agussuhardi/go-crud/helper"
	"agussuhardi/go-crud/repository"
	"agussuhardi/go-crud/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	db := application.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories", categoryController.Update)
	router.DELETE("/api/categories", categoryController.Delete)

	service := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := service.ListenAndServe()
	helper.PanicIfError(err)

}
