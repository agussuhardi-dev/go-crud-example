package service

import (
	"agussuhardi/go-crud/helper"
	"agussuhardi/go-crud/model/domain"
	"agussuhardi/go-crud/model/web"
	"agussuhardi/go-crud/repository"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category), nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error) {

	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category), nil

}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	service.CategoryRepository.Delete(ctx, tx, category)
	return nil
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	return helper.ToCategoryResponse(category), nil
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	var categoryResponses []web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}
	return categoryResponses
}
