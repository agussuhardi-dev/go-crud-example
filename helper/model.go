package helper

import (
	"agussuhardi/go-crud/model/domain"
	"agussuhardi/go-crud/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
