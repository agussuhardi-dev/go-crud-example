package controller

import (
	"agussuhardi/go-crud/helper"
	"agussuhardi/go-crud/model/web"
	"agussuhardi/go-crud/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse, err := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    201,
		Status:  "OK",
		Message: "Category has been created",
		Data:    categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	id := param.ByName("id")
	categoryUpdateRequest.Id, err = strconv.Atoi(id)
	helper.PanicIfError(err)

	categoryUpdateRequest.Name = param.ByName("name")
	categoryResponse, err := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Category has been updated",
		Data:    categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	id_, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id_)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Category has been deleted",
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	id_, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	categoryResponse, err := controller.CategoryService.FindById(request.Context(), id_)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Category has been found",
		Data:    categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Category has been found",
		Data:    categoryResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)

}
