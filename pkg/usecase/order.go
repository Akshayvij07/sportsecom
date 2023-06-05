package usecase

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
)

type Orderusecase struct {
	cartRepo  interfaces.CartRepo
	orderRepo interfaces.OrderRepo
}

func NewOrderUseCase(orderRepo interfaces.OrderRepo, cartRepo interfaces.CartRepo) services.Orderusecase {
	return &Orderusecase{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (c *Orderusecase) PlaceOrder(ctx context.Context, UserID, paymentMethodId int) (domain.Orders, error) {
	order, err := c.orderRepo.OrderAll(ctx, uint(UserID), paymentMethodId)
	return order, err
}

func (c *Orderusecase) CancelOrder(ctx context.Context, orderId, userId int) error {
	err := c.orderRepo.CancelOrder(ctx, orderId, userId)
	return err
}

func (c *Orderusecase) Listorders(ctx context.Context, userid int) ([]respondse.OrderResponse, error) {
	var orders []respondse.OrderResponse
	orders, err := c.orderRepo.Listorders(ctx)
	return orders, err
}
func (c *Orderusecase) UListorders(ctx context.Context, userid int) ([]respondse.OrderResponse, error) {
	var orders []respondse.OrderResponse
	orders, err := c.orderRepo.UListorders(ctx, userid)
	return orders, err
}

func (c *Orderusecase) Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error) {
	order, err = c.orderRepo.Listorder(ctx, Orderid, UserId)
	return order, err
}

func (c *Orderusecase) ReturnOrder(userId, orderId int) (float64, error) {
	total, err := c.orderRepo.ReturnOrder(userId, orderId)
	return total, err
}

func (c *Orderusecase) ListofOrderStatuses(ctx context.Context) ([]domain.OrderStatus, error) {
	var status []domain.OrderStatus
	status, err := c.orderRepo.ListofOrderStatuses(ctx)
	return status, err
}

func (c *Orderusecase) AdminListorders(ctx context.Context, pagination request.Pagination) (orders []domain.Orders, err error) {
	orders, err = c.orderRepo.AdminListorders(ctx, pagination)
	return orders, err
}

func (c *Orderusecase) UpdateOrderStatus(ctx context.Context, update request.Update) error {
	err := c.orderRepo.UpdateOrderStatus(ctx, update)
	return err
}
