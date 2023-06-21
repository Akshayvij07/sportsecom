package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type CartDataBase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepo {
	return &CartDataBase{
		DB: DB,
	}
}

func (c *CartDataBase) AddCart(ctx context.Context, UserId int) (uint, error) {
	var cart domain.Cart
	query := `INSERT INTO carts(user_id,discount,total_price)VALUES($1,$2,$3)
	RETURNING id`
	if c.DB.Raw(query, UserId, 0, 0).Scan(&cart).Error != nil {
		return 0, errors.New("failed create the cart for the user")
	}
	fmt.Println("cart", cart)
	return cart.Id, nil
}

func (c *CartDataBase) FindProduct(ctx context.Context, id uint) (respondse.Product, error) {
	var product respondse.Product
	Query := `SELECT p.id,p.product_name as name,p.description,p.brand,p.prize,p.category_id,p.qty_in_stock,c.category_name,p.created_at,p.updated_at FROM products p 
	JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(Query, id).Scan(&product).Error
	return product, err
}

func (c *CartDataBase) FindUserCart(ctx context.Context, UserId int) (domain.Cart, error) {
	var cart domain.Cart
	query := `SELECT * from carts WHERE user_id = ?`
	if c.DB.Raw(query, UserId).Scan(&cart).Error != nil {
		return cart, errors.New("failed to find a cart with this user id")
	}
	fmt.Println("find cart", cart)
	return cart, nil

}

func (c *CartDataBase) FindCartIdandProductId(cxt context.Context, Cart_id uint, Product_id uint) (domain.CartItem, error) {
	var cartItem domain.CartItem
	Query := `SELECT * FROM cart_items WHERE cart_id = $1 AND product_id = $2`
	if c.DB.Raw(Query, Cart_id, Product_id).Scan(&cartItem).Error != nil {
		return cartItem, errors.New("cant find any cart_item with this cart_id and product_id")
	}
	return cartItem, nil

}

func (c *CartDataBase) AddCartItem(ctx context.Context, cartItem domain.CartItem) error {

	Query := `INSERT INTO cart_items(cart_id,product_id,qty)VALUES($1,$2,$3)`
	if c.DB.Raw(Query, cartItem.CartId, cartItem.ProductId, 1).Scan(&cartItem).Error != nil {
		return errors.New("can't add this item")
	}
	return nil
}

func (c *CartDataBase) FindCartByUserID(ctx context.Context, UserId int) (domain.Cart, error) {
	var cart domain.Cart
	Find := `SELECT * FROM carts WHERE user_id = ?`
	if c.DB.Raw(Find, UserId).Scan(&cart).Error != nil {
		return cart, errors.New("no cart found using this user_id")
	}
	return cart, nil
}
func (c *CartDataBase) RemoveCarItem(ctx context.Context, CartItemId uint) error {
	Query := `DELETE FROM cart_items WHERE id = $1`
	if c.DB.Exec(Query, CartItemId).Error != nil {
		return errors.New("failed to remove item from the cart")
	}
	return nil
}

func (c *CartDataBase) AddQuantity(ctx context.Context, cartItem domain.CartItem, qty int) error {
	Query := `UPDATE cart_items SET qty = qty+$1 WHERE id = $2`
	if c.DB.Exec(Query, qty, cartItem.Id).Error != nil {
		return errors.New("failed to add quantity")
	}
	return nil
}
func (c *CartDataBase) FindCartlistByCartID(ctx context.Context, cartID uint) (cartitems []respondse.Cartres, err error) {

	query := `SELECT ci.product_id, p.product_name, ci.qty, p.prize, p.qty_in_stock,
	(p.prize * ci.qty) AS sub_total FROM cart_items ci 
	INNER JOIN products p ON ci.product_id = p.id 

	
	WHERE ci.cart_id = ?;
	
	`
	if c.DB.Raw(query, cartID).Scan(&cartitems).Error != nil {
		return cartitems, errors.New("failed to show cartitems")
	}
	return cartitems, err
}
