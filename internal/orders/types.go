package orders

import (
	"context"

	repo "github.com/riyagarg2711/ecom-api-course/internal/adapters/postgresql/sqlc"
)

type orderItem struct {
	ProductID int64 `json:"productID"`
	Quantity int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerID"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(cts context.Context, tempOrder createOrderParams) (repo.Order,error)
}