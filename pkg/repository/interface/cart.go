package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type CartRepo interface {
	FindProduct(ctx context.Context, id int) (product respondse.Product, err error)
	AddCart(ctx context.Context, UserId int) (uint, error)
	FindCartIdandProductId(cxt context.Context, Cart_id uint, Product_id uint) (domain.CartItem, error)
	AddCartItem(ctx context.Context, cartItem domain.CartItem) error
	FindCartByUserID(ctx context.Context, UserId int) (domain.Cart, error)
	RemoveCarItem(ctx context.Context, CartItemId uint) error
	AddQuantity(ctx context.Context, cartItemId domain.CartItem, qty int) error
	FindCartlistByCartID(ctx context.Context, cartID uint) (cartitems []respondse.Cartres, err error)
	FindProductBySku(ctx context.Context, sku string) (respondse.Product, error)
}
