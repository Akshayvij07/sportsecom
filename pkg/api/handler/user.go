package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Akshayvij07/ecommerce/pkg/api/utilityHandler"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
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
			Message:    "unable signUp",
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

	Address, err := cr.userUseCase.AddAdress(c, UserID, newAddress)
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
		Data:       Address,
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

	Address, err := cr.userUseCase.UpdateAdress(c, UserID, UpdatedAddress)
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
		Data:       Address,
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
