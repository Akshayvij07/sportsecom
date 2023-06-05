package handler

import (
	"fmt"
	"net/http"

	"github.com/Akshayvij07/ecommerce/pkg/api/utilityHandler"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartUsecase services.CartUsecase
}

func NewCartHandler(CartUsecase services.CartUsecase) *CartHandler {
	return &CartHandler{
		CartUsecase: CartUsecase,
	}
}

// AddToCart godoc
// @Summary api for adding product to user cart
// @Description user can add stock of product to user cart
// Security ApiKeyAuth
// @Tags UsersCart
// @Accept json
// @Produce json
// @Param inputs body request.Cartreq{}  true "Input Field info"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /cart/add/item [post]
func (c CartHandler) AddItemToCart(ctx *gin.Context) {
	var body request.Cartreq
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't bind the body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	body.UserID, err = utilityHandler.GetUserIdFromContext(ctx)
	fmt.Println(body, err)
	err = c.CartUsecase.AddCartItem(ctx, body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "unable to add item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "Successfully added to cart",
		Data:       body,
		Errors:     nil,
	})

}

// RemoveFromCart godoc
// @Summary api for removing product from user cart
// @Description user can reduce stock of product to user cart
// Security ApiKeyAuth
// @Tags UsersCart
// @Accept json
// @Produce json
// @Param inputs body request.Cartreq{}  true "Input Field info"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /cart/remove/item [delete]
func (c *CartHandler) RemoveItem(ctx *gin.Context) {
	var body request.Cartreq

	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Unable to bind the body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	body.UserID, err = utilityHandler.GetUserIdFromContext(ctx)
	err = c.CartUsecase.RemoveItem(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Unable to remove the product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "Successfully removed",
		Data:       body.UserID,
		Errors:     nil,
	})
}

// AddQuntity
// @Summary User can delete a item
// @ID Add-Qantity
// @Description user can delete their cartitems by id
// @Tags UsersCart
// @Accept json
// @Produce json
// @Param input body request.Addcount{} true "Input Field"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /cart/Addcount [put]
func (c *CartHandler) Addcount(ctx *gin.Context) {

	var body request.Addcount

	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "faild to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	body.UserID, err = utilityHandler.GetUserIdFromContext(ctx)
	fmt.Println(body, err)
	err = c.CartUsecase.AddQuantity(ctx, body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Unable to Add count",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "count added",
		Data:       body,
		Errors:     nil,
	})

}

// viewCart godoc
// @summary api for get all cart item of user
// @description user can see all productItem that stored in cart
// @security ApiKeyAuth
// @ID Cart
// @Tags  UsersCart
// @Router /cart/viewcart [get]
// @Success 200 {object} respondse.Response{} "successfully got user cart items"
// @Failure 500 {object} respondse.Response{} "faild to get cart items"
func (c *CartHandler) ViewCartItems(ctx *gin.Context) {

	userID, err := utilityHandler.GetUserIdFromContext(ctx)
	cart, err := c.CartUsecase.FindUserCart(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Unable to get cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if cart.Id == 0 {
		ctx.JSON(http.StatusOK, respondse.Response{
			StatusCode: 200,
			Message:    "you are not add any products to cart",
			Data:       nil,
			Errors:     nil,
		})
		return
	}

	cartitems, err := c.CartUsecase.FindCartlistByCartID(ctx, cart.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Unable to get cartitems",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	if cartitems == nil {
		ctx.JSON(http.StatusOK, respondse.Response{
			StatusCode: 200,
			Message:    "sorry no products in your cart",
			Data:       nil,
			Errors:     nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "your carts here",
		Data:       cartitems,
		Errors:     nil,
	})
}
