package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/Akshayvij07/ecommerce/pkg/utilityHandler"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary User Signup
// @ID UserSignup
// @Description Create a new user with the specified details
// @Tags Users
// @Accept json
// @Productejson
// @Param input body request.UserSign{}  true "Input Field"
// @Success 200 {object} respondse.Response
// @Faliure 400 {object} respondse.Response
// @Router /signup [post]
func (cr *UserHandler) SignUp(c *gin.Context) {

	var user request.UserSign
	err := c.Bind(&user)
	fmt.Println(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, respondse.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	UserValue, err := cr.userUseCase.SignUp(c.Request.Context(), user)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Unable to SignUP",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, respondse.Response{
		StatusCode: 201,
		Message:    "user signUp successfully",
		Data:       UserValue,
		Errors:     nil,
	})
}

// LoginWithEmail
// @Summary User Login
// @ID UserLogin
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Product json
// @Param input body request.Login  true "Input Field"
// @Success 200 {object} respondse.Response
// @Faliure 400 {object} respondse.Response
// @Router /login [post]
func (cr *UserHandler) Login(c *gin.Context) {

	var user request.Login
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})

		return
	}

	ss, err := cr.userUseCase.Login(c.Request.Context(), user)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*1, "", "", false, true)
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "succesfully logged In",
		Data:       nil,
		Errors:     nil,
	})
}

// @Summary updatePassword_for_user
// @ID Update_Password
// @Description Update user Password.
// @Tags EditUsers
// @Accept json
// @Produce json
// @Param   inputs   body     request.Password{} true  "Input Field"
// @Success 200 {object} respondse.Response
// @Faliure 400 {object} respondse.Response
// @Router /UpdatePassword [patch]
func (cr *UserHandler) ChangePassword(c *gin.Context) {
	var user request.Password
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})

		return
	}
	user.UserID, err = utilityHandler.GetUserIdFromContext(c)

	//fmt.Println(user, err)
	err = cr.userUseCase.ChangePassword(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "unable change password",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, respondse.Response{
		StatusCode: 201,
		Message:    "password changed  successfully",
		Data:       nil,
		Errors:     nil,
	})

}

// Home
// @Summary HomePage
// @ID Homepage
// @Description landing page for users
// @Tags Users
// @Success 200 "success"
// @Failure 400 "failed"
// @Router /home [GET]
func (cr *UserHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "Welcome to Homepage",
		Data:       nil,
		Errors:     nil,
	})
}

// Edit
// @Summary EditProfile
// @ID EditProfile
// Description User can Edit Profile
// @Tags Users
// @Success 200 "success"
func (cr *UserHandler) EditUser(c *gin.Context) {

}

// UserLogout
// @Summary User Login
// @ID UserLogout
// @Description Logout as a user exit from the ecommerce site
// @Tags Users
// @Success 200 "success"
// @Failure 400 "Failed"
// @Router /logout [post]
func (cr *UserHandler) LogOut(c *gin.Context) {
	c.SetCookie("UserAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "Logout Succesfully ",
		Data:       nil,
		Errors:     nil,
	})
}

// @Summary AddAdrress_for_user
// @ID Add_Adress
// @Description Create a new user with the specified details.
// @Tags UsersAddress
// @Accept json
// @Produce json
// @Param   inputs   body     request.AddressReq{} true  "Input Field"
// @Success 200 {object} respondse.Response
// @Failure 422 {object} respondse.Response
// @Router /SaveAddress [post]
func (cr *UserHandler) AddAdress(c *gin.Context) {
	var newAddress request.AddressReq
	err := c.Bind(&newAddress)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, respondse.Response{
				StatusCode: 400,
				Message:    "failed to read request body",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	UserID, err := utilityHandler.GetUserIdFromContext(c)
	if newAddress.IsDefault == nil {
		newAddress.IsDefault = new(bool)
	}

	err = cr.userUseCase.AddAdress(c, UserID, newAddress, *newAddress.IsDefault)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant add this address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "SuccessFully Added YOUR Address",
		Data:       nil,
		Errors:     nil,
	})
}

// @Summary updateAdrress_for_user
// @ID Update_Adress
// @Description Update user Adresses.
// @Tags UsersAddress
// @Accept json
// @Produce json
// @Param   inputs   body     request.AddressReq{} true  "Input Field"
// @Success 200 {object} respondse.Response
// @Failure 422 {object} respondse.Response
// @Router /UpdateAddress [patch]
func (cr *UserHandler) UpdateAdress(c *gin.Context) {
	var UpdatedAddress request.AddressReq
	err := c.Bind(&UpdatedAddress)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, respondse.Response{
				StatusCode: 400,
				Message:    "failed to read request body",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	UserID, err := utilityHandler.GetUserIdFromContext(c)

	err = cr.userUseCase.UpdateAdress(c, UserID, UpdatedAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant update this address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "SuccessFully updated your address",
		Data:       nil,
		Errors:     nil,
	})
}

// viewAdress godoc
// @summary api for get address of user
// @description user can see their Adress
// @security ApiKeyAuth
// @Id User_Address
// @Tags  UsersAddress
// @Router /viewAddress [get]
// @Success 200 {object} respondse.Response{} "successfully get Address"
// @Failure 500 {object} respondse.Response{} "faild to get Address"
func (c *UserHandler) VeiwAddress(ctx *gin.Context) {

	userID, err := utilityHandler.GetUserIdFromContext(ctx)
	Adress, err := c.userUseCase.VeiwAdress(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "something Went Wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	{
		ctx.JSON(http.StatusOK, respondse.Response{
			StatusCode: 200,
			Message:    "this your Address",
			Data:       Adress,
			Errors:     nil,
		})
		return
	}
}

// AddToWishList godoc
// @summary api to add a product to wish list
// @descritpion user can add product to wish list
// @security ApiKeyAuth
// @id AddToWishList
// @tags Wishlist
// @Param product_id path string true "product_id"
// @Router /Addwishlist/{product_id} [post]
// @Success 200 {object} respondse.Response{} "successfully added product to wishlist"
// @Failure 400 {object} respondse.Response{} "invalid input"
func (c *UserHandler) AddToWishList(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	productID, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID, err := utilityHandler.GetUserIdFromContext(ctx)

	var wishList = domain.WishList{
		ProductID: uint(productID),
		UserID:    uint(userID),
	}
	err = c.userUseCase.AddToWishList(ctx, wishList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "something Went Wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	{
		ctx.JSON(http.StatusOK, respondse.Response{
			StatusCode: 200,
			Message:    "successfully add to whishlist",
			Data:       wishList,
			Errors:     nil,
		})
		return
	}
}

// RemoveFromWishList godoc
// @summary api to remove a product from wish list
// @descritpion user can remove a product from wish list
// @security ApiKeyAuth
// @id RemoveFromWishList
// @tags Wishlist
// @Param product_id path string true "product_id"
// @Router /Removewishlist/{product_id} [delete]
// @Success 200 {object} respondse.Response{} "successfully removed product item from wishlist"
// @Failure 400 {object} respondse.Response{} "invalid input"
func (c *UserHandler) RemoveFromWishList(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	productID, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "some err to convert",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID, err := utilityHandler.GetUserIdFromContext(ctx)

	var wishList = domain.WishList{
		ProductID: uint(productID),
		UserID:    uint(userID),
	}

	if err := c.userUseCase.RemoveFromWishList(ctx, wishList); err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "something Went Wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	{
		ctx.JSON(http.StatusOK, respondse.Response{
			StatusCode: 200,
			Message:    "successfully remove from whishlist",
			Data:       wishList,
			Errors:     nil,
		})
		return
	}
}

// GetWishListI godoc
// @summary api get all wish list items of user
// @descritpion user get all wish list items
// @security ApiKeyAuth
// @id GetWishListI
// @tags Wishlist
// @Router /wishlist [get]
// @Success 200 "Successfully wish list items got"
// @Success 200 "Wish list is empty"
// @Failure 400  "faild to get user wish list items"
func (u *UserHandler) GetWishList(ctx *gin.Context) {

	userID, err := utilityHandler.GetUserIdFromContext(ctx)

	wishList, err := u.userUseCase.ListWishlist(ctx, uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "something Went Wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	{
		ctx.JSON(http.StatusOK, respondse.Response{
			StatusCode: 200,
			Message:    "successfully get the whishlist",
			Data:       wishList,
			Errors:     nil,
		})
		return
	}
}

// ViewInvoice
// @Summary User can view Invoice
// @ID Get-Invoice
// @Description Admin can view the sales report
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /invoice [get]
func (cr *UserHandler) GetInvoice(ctx *gin.Context) {

	userID, err := utilityHandler.GetUserIdFromContext(ctx)
	Invo := utilityHandler.GenerateInvoiceNumber()

	Invoice, err := cr.userUseCase.GetInvoice(ctx, userID)
	Invoice.InvoiceNumber = Invo
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "Invoice",
		Data:       Invoice,
		Errors:     nil,
	})

}
