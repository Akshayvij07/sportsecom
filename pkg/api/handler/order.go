package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/Akshayvij07/ecommerce/pkg/utilityHandler"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUseCase services.Orderusecase
}

func NewOrderHandler(OrderUseCase services.Orderusecase) *OrderHandler {
	return &OrderHandler{
		OrderUseCase: OrderUseCase,
	}
}

// Place Order
// @Summary Buy all items from the user's cart
// @ID buyAll
// @Description This endpoint allows a user to purchase all items in their cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param payment_id path string true "payment_id"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /order/place_order/{payment_id} [post]
func (cr *OrderHandler) CashonDElivery(ctx *gin.Context) {

	paramsId := ctx.Param("payment_id")
	PaymentMethodId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	UserID, err := utilityHandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.OrderUseCase.PlaceOrder(ctx, UserID, PaymentMethodId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "orderplaced",
		Data:       order,
		Errors:     nil,
	})
}

func (c *OrderHandler) RazorpayCheckout(ctx *gin.Context) {
	UserID, err := utilityHandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := ctx.Param("payment_id")
	payment_id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	razorPayOrder, err := c.OrderUseCase.Razorpay(ctx, UserID, payment_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	fmt.Println("herer")
	ctx.HTML(http.StatusOK, "razor.html", razorPayOrder)
}

func (cr *OrderHandler) RazorpayVerify(ctx *gin.Context) {
	// "razorpay_payment_id": response.razorpay_payment_id,
	// "razorpay_order_id": response.razorpay_order_id,
	// "razorpay_signature": response.razorpay_signature,
	// "payment_id": payment_id,

	razorPayPaymentId := ctx.Request.PostFormValue("razorpay_payment_id")
	razorPayOrderId := ctx.Request.PostFormValue("razorpay_order_id")
	razorpay_signature := ctx.Request.PostFormValue("razorpay_signature")
	paramsId := ctx.Request.PostFormValue("payment_id")

	userId, err := utilityHandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paymentid, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	body := request.RazorPayRequest{
		RazorPayPaymentId:  razorPayPaymentId,
		RazorPayOrderId:    razorPayOrderId,
		Razorpay_signature: razorpay_signature,
	}

	err = cr.OrderUseCase.VerifyRazorPay(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    " faild to veify razorpay order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.OrderUseCase.PlaceOrder(ctx, userId, paymentid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "orderplaced",
		Data:       order,
		Errors:     nil,
	})
	cr.OrderUseCase.RemoveCartItems(ctx, uint(userId))

}

// CancelOrder
// @Summary Cancels a specific order for the currently logged in user
// @ID cancel-order
// @Description Endpoint for cancelling an order associated with a user
// @Tags UserOrder
// @Accept json
// @Produce json
// @Param orderId path int true "ID of the order to be cancelled"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /order/cancel/{orderId} [patch]
func (cr *OrderHandler) CancelOrder(ctx *gin.Context) {
	UserID, err := utilityHandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := ctx.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.OrderUseCase.CancelOrder(ctx, orderId, UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't cancel order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "order canceld",
		Data:       nil,
		Errors:     nil,
	})
}

// listorder
// @Summary to get order details
// @ID view-order-by-id
// @Description retrieving the details of a specific order identified by its order ID.
// @Tags UserOrder
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} respondse.Response "Successfully fetched order details"
// @Failure 400 {object} respondse.Response "Failed to fetch order details"
// @Router /order/view/{order_id} [get]
func (cr *OrderHandler) ListOrder(ctx *gin.Context) {
	UserID, err := utilityHandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := ctx.Param("order_id")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.OrderUseCase.Listorder(ctx, UserID, orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       order,
		Errors:     nil,
	})
}

// ListAllOrders
// @Summary for geting all order list
// @ID List-all-orders
// @Description Endpoint for getting all orders
// @Tags UserOrder
// @Accept json
// @Produce json
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /order/listall [get]
func (cr *OrderHandler) ListAllOrders(ctx *gin.Context) {
	UserID, err := utilityHandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	fmt.Println(UserID)
	orders, err := cr.OrderUseCase.UListorders(ctx, UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       orders,
		Errors:     nil,
	})
}

// ReturnOrder
// @Summary Return a specific order for the currently logged in user
// @ID return-order
// @Description Endpoint for Returning an order associated with a user
// @Tags UserOrder
// @Accept json
// @Produce json
// @Param orderId path int true "ID of the order to be Returned"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /order/return/{orderId} [patch]
func (cr *OrderHandler) ReturnOrder(ctx *gin.Context) {
	UserID, err := utilityHandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := ctx.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	returnAmount, err := cr.OrderUseCase.ReturnOrder(UserID, orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't return order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "order returnd ",
		Data:       returnAmount,
		Errors:     nil,
	})
}

// ListAllOrderStatuses
// @Summary for geting all order status list
// @ID List-all-orderStatus
// @Description Endpoint for getting all orderStatuses
// @Tags AdminOrder
// @Accept json
// @Produce json
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /admin/order/Status [get]
func (cr *OrderHandler) Statuses(ctx *gin.Context) {
	status, err := cr.OrderUseCase.ListofOrderStatuses(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't List the order_statuses",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "order_statuses ",
		Data:       status,
		Errors:     nil,
	})

}

// @Summary FindAllorders In admin side
// @Id FindAllordersInshop
// @Discription list of orders.
// @tags AdminOrder
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param perPage query int false "Number of items to retrieve per page"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /admin/order/Allorders [get]
func (cr *OrderHandler) AllOrders(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {

		page = 1
	}

	perPage, err := strconv.Atoi(ctx.Query("perPage"))
	if err != nil {

		perPage = 10
	}

	ListofOrders := request.Pagination{
		Page:    uint(page),
		PerPage: uint(perPage),
	}

	orders, err := cr.OrderUseCase.AdminListorders(ctx, ListofOrders)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't List the orders",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "orders ",
		Data:       orders,
		Errors:     nil,
	})
}

// @Summary Updateorderstatus
// @ID Order_status
// @Description update the order statuses by every orderid.
// @Tags AdminOrder
// @Accept json
// @Produce json
// @Param   inputs   body request.Update{}  true  "Input Field"
// @Success 200 {object} respondse.Response
// @Failure 422 {object} respondse.Response
// @Router /admin/order/UpdateStatus [patch]
func (cr *OrderHandler) UpdateOrderStatus(ctx *gin.Context) {
	var Update request.Update
	err := ctx.Bind(&Update)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.OrderUseCase.UpdateOrderStatus(ctx, Update)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't do something went wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "changed the orderstatus ",
		Data:       nil,
		Errors:     nil,
	})

}
