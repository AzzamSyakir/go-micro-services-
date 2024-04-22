package repository

import (
	"database/sql"
	"fmt"
	"go-micro-services/src/order-service/entity"
	model_response "go-micro-services/src/order-service/model/response"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	orderRepository := &OrderRepository{}
	return orderRepository
}

func (orderRepository *OrderRepository) Order(begin *sql.Tx, orders *entity.Order) (result *model_response.Response[*model_response.OrderResponse], err error) {
	rows, queryErr := begin.Query(
		`INSERT INTO "orders"(id, user_id, total_price, total_paid, total_return, receipt_code, created_at, updated_at, deleted_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		orders.Id,
		orders.UserId,
		orders.TotalPrice,
		orders.TotalPaid,
		orders.TotalReturn,
		orders.ReceiptCode,
		orders.CreatedAt,
		orders.UpdatedAt,
		orders.DeletedAt,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	ordersResponse := &model_response.OrderResponse{
		Id:          orders.Id,
		UserId:      orders.UserId,
		Name:        orders.Name,
		TotalPrice:  orders.TotalPrice,
		TotalPaid:   orders.TotalPaid,
		TotalReturn: orders.TotalReturn,
		ReceiptCode: orders.ReceiptCode,
		CreatedAt:   orders.CreatedAt,
		UpdatedAt:   orders.UpdatedAt,
	}
	result = &model_response.Response[*model_response.OrderResponse]{
		Data: ordersResponse,
	}
	err = nil
	return result, err
}
func (orderRepository *OrderRepository) OrderProducts(begin *sql.Tx, orderProducts *entity.OrderProducts) (result *entity.OrderProducts, err error) {
	rows, queryErr := begin.Query(
		`INSERT INTO "order_products"(id, order_id, product_id, total_price, qty, created_at, updated_at, deleted_at) values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		orderProducts.Id,
		orderProducts.OrderId,
		orderProducts.ProductId,
		orderProducts.TotalPrice,
		orderProducts.Qty,
		orderProducts.CreatedAt,
		orderProducts.UpdatedAt,
		orderProducts.DeletedAt,
	)
	if queryErr != nil {
		fmt.Println("queryErr : ", queryErr)
		result = nil
		err = queryErr
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	result = orderProducts
	err = nil
	return result, err
}
