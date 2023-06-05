package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type Orderusecase interface {
	PlaceOrder(ctx context.Context, UserID, paymentMethodId int) (domain.Orders, error)
	UpdateOrderStatus(ctx context.Context, update request.Update) error
	AdminListorders(ctx context.Context, pagination request.Pagination) (orders []domain.Orders, err error)
	ListofOrderStatuses(ctx context.Context) ([]domain.OrderStatus, error)
	ReturnOrder(userId, orderId int) (float64, error)
	Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error)
	Listorders(ctx context.Context, userid int) ([]respondse.OrderResponse, error)
	CancelOrder(ctx context.Context, orderId, userId int) error
	UListorders(ctx context.Context, userid int) ([]respondse.OrderResponse, error) 
}
