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

	"github.com/cockroachdb/cockroach-go/v2/crdb"
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

func (productUseCase *ProductUseCase) GetOneById(id string) (result *model_response.Response[*entity.Product]) {
	transaction, transactionErr := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
	if transactionErr != nil {
		errorMessage := fmt.Sprintf("transaction failed :%s", transactionErr)
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: errorMessage,
			Data:    nil,
		}

		return result
	}
	productFound, productFoundErr := productUseCase.ProductRepository.GetOneById(transaction, id)
	if productFoundErr != nil {
		errorMessage := fmt.Sprintf("product-service UseCase, DetailProduct is failed, GetProduct failed : %s", productFoundErr)
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: errorMessage,
			Data:    nil,
		}

		return result
	}
	errorMessage := fmt.Sprintf("product-service UseCase, DetailProduct is failed, product is not found by id %s", id)
	if productFound == nil {
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusNotFound,
			Message: errorMessage,
			Data:    nil,
		}

		return result
	}

	result = &model_response.Response[*entity.Product]{
		Code:    http.StatusOK,
		Message: "product-service UseCase, DetailProduct is succeed.",
		Data:    productFound,
	}

	return result
}

func (productUseCase *ProductUseCase) UpdateProduct(id string, request *model_request.ProductPatchOneByIdRequest) (result *model_response.Response[*entity.Product]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
		if err != nil {
			return err
		}

		foundProduct, err := productUseCase.ProductRepository.GetOneById(begin, id)
		if err != nil {
			return err
		}
		if foundProduct == nil {
			err = begin.Rollback()
			result = &model_response.Response[*entity.Product]{
				Code:    http.StatusNotFound,
				Message: "product-service UseCase Update Stock is failed, product is not found by id.",
				Data:    nil,
			}
			return err
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
			return err
		}

		err = begin.Commit()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusOK,
			Message: "ProductProductCase Update Stock is succeed.",
			Data:    patchedProduct,
		}
		return err
	})

	if beginErr != nil {
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusInternalServerError,
			Message: "ProductProductCase Update Stock  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}
	return result
}

func (productUseCase *ProductUseCase) ListProduct() (result *model_response.Response[[]*entity.Product]) {
	transaction, transactionErr := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
	if transactionErr != nil {
		errorMessage := fmt.Sprintf("transaction failed :%s", transactionErr)
		result = &model_response.Response[[]*entity.Product]{
			Code:    http.StatusNotFound,
			Message: errorMessage,
			Data:    nil,
		}

		return result
	}

	fetchproduct, fetchproductErr := productUseCase.ProductRepository.ListProducts(transaction)
	if fetchproductErr != nil {
		errorMessage := fmt.Sprintf("product-service UseCase, ListProduct is failed, Getproduct failed : %s", fetchproductErr)
		result = &model_response.Response[[]*entity.Product]{
			Code:    http.StatusNotFound,
			Message: errorMessage,
			Data:    nil,
		}
		return result
	}

	if fetchproduct.Data == nil {
		result = &model_response.Response[[]*entity.Product]{
			Code:    http.StatusNotFound,
			Message: "product-service UseCase, ListProduct is failed, data product is empty ",
			Data:    nil,
		}
		return result
	}

	result = &model_response.Response[[]*entity.Product]{
		Code:    http.StatusOK,
		Message: "product-service UseCase, ListProduct is succeed.",
		Data:    fetchproduct.Data,
	}
	return result
}

func (productUseCase *ProductUseCase) DeleteProduct(id string) (result *model_response.Response[*entity.Product]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := productUseCase.DatabaseConfig.ProductDB.Connection.Begin()
		if err != nil {
			return err
		}

		deletedproduct, deletedproductErr := productUseCase.ProductRepository.DeleteOneById(begin, id)
		if deletedproductErr != nil {
			err = begin.Rollback()
			result = &model_response.Response[*entity.Product]{
				Code:    http.StatusNotFound,
				Message: "product-service UseCase, DeleteProduct is failed, " + deletedproductErr.Error(),
				Data:    nil,
			}
			return err
		}
		if deletedproduct == nil {
			err = begin.Rollback()
			result = &model_response.Response[*entity.Product]{
				Code:    http.StatusNotFound,
				Message: "product-service UseCase, DeleteProduct is failed, product is not deleted by id, " + id,
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusOK,
			Message: "product-service UseCase DeleteProduct is succeed.",
			Data:    deletedproduct,
		}
		return err
	})

	if beginErr != nil {
		result = &model_response.Response[*entity.Product]{
			Code:    http.StatusInternalServerError,
			Message: "product-service UseCase DeleteProduct is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}
