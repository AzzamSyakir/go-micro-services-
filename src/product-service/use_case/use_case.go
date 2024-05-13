package use_case

import (
	"fmt"
	"go-micro-services/src/product-service/config"
	"go-micro-services/src/product-service/entity"
	model_request "go-micro-services/src/product-service/model/request/controller"
	model_response "go-micro-services/src/product-service/model/response"
	"go-micro-services/src/product-service/repository"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

type ProductUseCase struct {
	DatabaseConfig    *config.DatabaseConfig
	ProductRepository *repository.ProductRepository
}

func NewProductUseCase(
	databaseConfig *config.DatabaseConfig,
	productRepository *repository.ProductRepository,

) *ProductUseCase {
	productUseCase := &ProductUseCase{
		DatabaseConfig:    databaseConfig,
		ProductRepository: productRepository,
	}
	return productUseCase
}
func (productUseCase *ProductUseCase) CreateProduct(request *model_request.CreateProduct) (result *model_response.Response[*entity.Product], err error) {
	begin, err := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusBadRequest,
			Message: "ProductUseCase CreateProduct is failed, begin fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	if request.Name.String == "" || request.Price.Int64 == 0 || request.Stock.Int64 == 0 {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusBadRequest,
			Message: "ProductUseCase CreateProduct is failed, Please input data correctly, data cannot be empty",
			Data:    nil,
		}
		return result, rollback
	}
	firstLetter := strings.ToUpper(string(request.Name.String[0]))
	rand.Seed(time.Now().UnixNano())
	randomDigits := rand.Intn(900) + 100
	sku := fmt.Sprintf("%s%d", firstLetter, randomDigits)

	currentTime := null.NewTime(time.Now(), true)
	newproduct := &entity.Product{
		Id:         null.NewString(uuid.NewString(), true),
		Name:       request.Name,
		Sku:        null.NewString(sku, true),
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryId: request.CategoryId,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
		DeletedAt:  null.NewTime(time.Time{}, false),
	}

	createdProduct, err := productUseCase.ProductRepository.CreateProduct(begin, newproduct)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusBadRequest,
			Message: "ProductUseCase CreateProduct is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	commit := begin.Commit()
	result = &model_response.Response[*entity.Product]{
		Code:    http.StatusBadRequest,
		Message: "ProductUseCase CreateProduct is success",
		Data:    createdProduct,
	}
	return result, commit
}

func (productUseCase *ProductUseCase) GetOneById(id string) (result *model_response.Response[*entity.Product], err error) {
	begin, err := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, DetailProduct, begin failed, " + err.Error(),
			Data:    nil,
		}

		return result, rollback
	}
	productFound, err := productUseCase.ProductRepository.GetOneById(begin, id)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, DetailProduct is failed, GetProduct failed, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	if productFound == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, DetailProduct is failed, product is not found by id" + id,
			Data:    nil,
		}
		return result, rollback
	}

	commit := begin.Commit()
	result = &model_response.Response[*entity.Product]{
		Code:    http.StatusOK,
		Message: "product-service UseCase, DetailProduct is succeed.",
		Data:    productFound,
	}

	return result, commit
}

func (productUseCase *ProductUseCase) UpdateProduct(id string, request *model_request.ProductPatchOneByIdRequest) (result *model_response.Response[*entity.Product], err error) {
	begin, err := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, UpdateProduct fail begin is failed," + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	foundProduct, err := productUseCase.ProductRepository.GetOneById(begin, id)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase Update Product is failed, product is not found by id" + id,
			Data:    nil,
		}
		return result, rollback
	}
	if foundProduct == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase Update Product is failed, product is not found by id, " + id,
			Data:    nil,
		}
		return result, rollback
	}

	if request.Name.Valid {
		foundProduct.Name = request.Name
	}
	if request.Stock.Valid {
		foundProduct.Stock = request.Stock
	}
	if request.Price.Valid {
		foundProduct.Price = request.Price
	}
	if request.CategoryId.Valid {
		foundProduct.CategoryId = request.CategoryId
	}
	foundProduct.UpdatedAt = null.NewTime(time.Now(), true)

	patchedProduct, err := productUseCase.ProductRepository.PatchOneById(begin, id, foundProduct)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, Query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	commit := begin.Commit()
	result = &model_response.Response[*entity.Product]{
		Code:    http.StatusNotFound,
		Message: "product-service UseCase Update Product is succes.",
		Data:    patchedProduct,
	}
	return result, commit
}

func (productUseCase *ProductUseCase) ListProduct() (result *model_response.Response[[]*entity.Product], err error) {
	begin, beginErr := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
	if beginErr != nil {
		rollback := begin.Rollback()
		errorMessage := fmt.Sprintf("begin failed :%s", beginErr)
		result = &model_response.Response[[]*entity.Product]{
			Code:    http.StatusNotFound,
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}

	fetchproduct, fetchproductErr := productUseCase.ProductRepository.ListProducts(begin)
	if fetchproductErr != nil {
		rollback := begin.Rollback()
		errorMessage := fmt.Sprintf("product-service UseCase, ListProduct is failed, Getproduct failed : %s", fetchproductErr)
		result = &model_response.Response[[]*entity.Product]{
			Code:    http.StatusNotFound,
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}

	if fetchproduct.Data == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[[]*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, ListProduct is failed, data product is empty ",
			Data:    nil,
		}
		return result, rollback
	}
	commit := begin.Commit()
	result = &model_response.Response[[]*entity.Product]{
		Code:    http.StatusOK,
		Message: "product-service UseCase, ListProduct is succeed.",
		Data:    fetchproduct.Data,
	}
	return result, commit
}

func (productUseCase *ProductUseCase) DeleteProduct(id string) (result *model_response.Response[*entity.Product], err error) {
	begin, err := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, DeleteProduct is failed, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	deletedproduct, deletedproductErr := productUseCase.ProductRepository.DeleteOneById(begin, id)
	if deletedproductErr != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, DeleteProduct is failed, " + deletedproductErr.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	if deletedproduct == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, DeleteProduct is failed, product is not deleted by id, " + id,
			Data:    nil,
		}
		return result, rollback
	}
	rollback := begin.Commit()
	result = &model_response.Response[*entity.Product]{
		Code:    http.StatusOK,
		Message: "product-service UseCase DeleteProduct is succeed.",
		Data:    deletedproduct,
	}
	return result, rollback
}
