package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
)

type CartUseCase struct {
	CartRepo interfaces.CartRepo
}

func NewCartUseCase(CartRepo interfaces.CartRepo) services.CartUsecase {
	return &CartUseCase{
		CartRepo: CartRepo,
	}
}

func (c *CartUseCase) AddCartItem(ctx context.Context, body request.Cartreq) error {
	//validate product (find product using product_id)
	product, err := c.CartRepo.FindProduct(ctx, uint(body.ProductId))

	if err != nil {
		return errors.New("Inavlid Product")
	}

	//check the quantity of the product and if out od stock return error

	if product.Qty_in_stock == 0 {
		return errors.New("The product is unavailable at the moment")
	}
	//Find the user cart with user id
	cart, err := c.CartRepo.FindCartByUserID(ctx, body.UserID)
	if err != nil {
		return errors.New("the cart belonfs to defferent user")
	} else if cart.Id == 0 {
		cartId, err := c.CartRepo.AddCart(ctx, body.UserID)
		if err != nil {
			return errors.New("failed create a cart for the user")
		}
		cart.Id = cartId
	}
	//add product_id and cart_id to a table cart items
	cartItem, err := c.CartRepo.FindCartIdandProductId(ctx, uint(cart.Id), uint(body.ProductId))
	if err != nil {
		return err
	} else if cartItem.Id != 0 {
		//return errors.New("product is allready save in cart")
		c.CartRepo.AddQuantity(ctx, cartItem, 1)
		return nil
	}
	fmt.Println(cart.Id)
	cartitem := domain.CartItem{
		CartId:    uint(cart.Id),
		ProductId: uint(body.ProductId),
	}
	if err := c.CartRepo.AddCartItem(ctx, cartitem); err != nil {
		return err
	}

	return nil

}

func (c *CartUseCase) RemoveItem(ctx context.Context, body request.Cartreq) error {
	//product validation
	product, err := c.CartRepo.FindProduct(ctx, uint(body.ProductId))
	if err != nil {
		return errors.New("invalid product")
	}
	if product.Id == 0 {
		return errors.New("now product is unavailable")
	}
	//find user_cart using user_id
	cart, err := c.CartRepo.FindCartByUserID(ctx, body.UserID)
	if err != nil {
		return errors.New("no cart found")
	} else if cart.Id == 0 {
		return errors.New("unable to remove product from the cart its empty")
	}
	cartItem, err := c.CartRepo.FindCartIdandProductId(ctx, uint(cart.Id), uint(body.ProductId))
	if err != nil {
		return err
	} else if cartItem.Id == 0 {
		return errors.New("no product with such id")
	}

	if err := c.CartRepo.RemoveCarItem(ctx, cartItem.Id); err != nil {
		return err
	}
	return nil

}

func (c *CartUseCase) FindUserCart(ctx context.Context, UserId int) (domain.Cart, error) {
	var cart domain.Cart
	cart, err := c.CartRepo.FindCartByUserID(ctx, UserId)
	if err != nil {
		return cart, err
	}
	return cart, err
}

func (c *CartUseCase) AddQuantity(ctx context.Context, body request.Addcount) error {
	//product validation
	product, err := c.CartRepo.FindProduct(ctx, uint(body.ProductId))
	if err != nil {
		return errors.New("invalid product")
	}
	if product.Id == 0 {
		return errors.New("now product is unavailable")
	}
	if body.Count < 0 {
		return errors.New("sorry can't enter a value less than zero")
	} else if body.Count > int(product.Qty_in_stock) {
		return errors.New("The product quntity limit exceeded")
	}

	cart, err := c.CartRepo.FindCartByUserID(ctx, body.UserID)
	if err != nil {
		return errors.New("user have no cart")
	}

	cartitem, err := c.CartRepo.FindCartIdandProductId(ctx, cart.Id, uint(body.ProductId))
	if err != nil {
		return err
	} else if cartitem.Id == 0 {
		return errors.New("product is not exist in your cart")
	}

	err = c.CartRepo.AddQuantity(ctx, cartitem, body.Count)
	if err != nil {
		return errors.New("unable to add more quantity")
	}

	return nil

}
func (c *CartUseCase) FindCartlistByCartID(ctx context.Context, cartID uint) ([]respondse.Cartres, error) {
	cartitems, err := c.CartRepo.FindCartlistByCartID(ctx, cartID)
	if err != nil {
		return cartitems, err
	}
	return cartitems, nil
}
