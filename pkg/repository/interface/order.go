package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type OrderRepo interface {
	OrderAll(ctx context.Context, UserID uint, paymentMethodId int) (domain.Orders, error)
	UpdateOrderStatus(ctx context.Context, update request.Update) error
	AdminListorders(ctx context.Context, pagination request.Pagination) (orders []domain.Orders, err error)
	ListofOrderStatuses(ctx context.Context) (status []domain.OrderStatus, err error)
	ReturnOrder(userId, orderId int) (float64, error)
	Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error)
	Listorders(ctx context.Context) ([]respondse.OrderResponse, error)
	CancelOrder(ctx context.Context, orderId, userId int) error
	UListorders(ctx context.Context, UserId int) ([]respondse.UserOrderResponse, error)
	RemoveCarItems(ctx context.Context, CartItemId uint) error
	FindCartByUserID(ctx context.Context, UserId int) (domain.Cart, error)
	FindWalletByUserID(ctx context.Context, UserId int) (domain.Wallet, error)
	SaveWallet(ctx context.Context, UserId int) (uint, error)
}
