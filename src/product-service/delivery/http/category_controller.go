package http

import (
	"encoding/json"
	model_request "go-micro-services/src/product-service/model/request/controller"
	"go-micro-services/src/product-service/model/response"
	"go-micro-services/src/product-service/use_case"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryController struct {
	CategoryUseCase *use_case.CategoryUseCase
}

func NewCategoryController(categoryUseCase *use_case.CategoryUseCase) *CategoryController {
	CategoryController := &CategoryController{
		CategoryUseCase: categoryUseCase,
	}
	return CategoryController
}
func (categoryController *CategoryController) GetCategory(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]
	foundCategory, foundCategoryErr := categoryController.CategoryUseCase.GetOneById(id)
	if foundCategoryErr == nil {
		response.NewResponse(writer, foundCategory)
	}
}

func (CategoryController *CategoryController) UpdateCategory(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	request := &model_request.CategoryRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}
	result := CategoryController.CategoryUseCase.UpdateCategory(id, request)

	response.NewResponse(writer, result)
}

func (CategoryController *CategoryController) CreateCategory(writer http.ResponseWriter, reader *http.Request) {

	request := &model_request.CategoryRequest{}

	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, "Failed to decode request body: "+decodeErr.Error(), http.StatusBadRequest)
		return
	}

	result := CategoryController.CategoryUseCase.CreateCategory(request)

	response.NewResponse(writer, result)
}

func (CategoryController *CategoryController) ListCategories(writer http.ResponseWriter, reader *http.Request) {
	Category, CategoryErr := CategoryController.CategoryUseCase.ListCategories()
	if CategoryErr == nil {
		response.NewResponse(writer, Category)
	}
}

func (CategoryController *CategoryController) DeleteCategory(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	result := CategoryController.CategoryUseCase.DeleteCategory(id)

	response.NewResponse(writer, result)
}
