package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type CartUsecase interface {
	AddCartItem(ctx context.Context, body request.Cartreq) error
	RemoveItem(ctx context.Context, body request.Cartreq) error
	FindUserCart(ctx context.Context, UserId int) (domain.Cart, error)
	FindCartlistByCartID(ctx context.Context, cartID uint) ([]respondse.Cartres, error)
	AddQuantity(ctx context.Context, body request.Addcount) error
}
