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

type AdminHandler struct {
	AdminUsecase services.AdminUsecase
}

func NewAdminHandler(AdmUsecase services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		AdminUsecase: AdmUsecase,
	}
}

// @Summary SaveAdmin
// @ID SaveAdmin
// @Description Save admin with details
// @Tags Admin
// @Accept json
// @Produce json
// @Param inputs body domain.Admin{}  true "Input Field"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /admin/signup [post]
func (cr *AdminHandler) SaveAdmin(c *gin.Context) {
	var admin request.Admin

	err := c.Bind(&admin)
	fmt.Println(admin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, respondse.Response{
			StatusCode: 422,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.AdminUsecase.SaveAdmin(c.Request.Context(), admin)
	fmt.Println(admin)
	fmt.Println("testa1")
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Unable to Signup admin",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	fmt.Println("testa2")
	c.JSON(http.StatusCreated, respondse.Response{
		StatusCode: 201,
		Message:    "Signup Succesfully",
		Data:       nil,
		Errors:     nil,
	})
}

// @Summary LoginAdmin
// @ID LogInAdmin
// @Description Login admin with details
// @Tags Admin
// @Accept json
// @Produce json
// @Param inputs body request.AdminLogin{}  true "Input Field"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /admin/login [post]
func (cr *AdminHandler) LoginAdmin(c *gin.Context) {
	var admin request.AdminLogin

	err := c.Bind(&admin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, respondse.Response{
			StatusCode: 422,
			Message:    "Unable to Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ss, err := cr.AdminUsecase.LoginAdmin(c.Request.Context(), admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Failed to Login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuth", ss, 3600*24*1, "", "", false, true)
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "Succesfully Logged In",
		Data:       ss,
		Errors:     nil,
	})
}

// AdminLogout
// @Summary AdminLogout
// @ID AdminLogout
// @Description Logout as a user exit from the ecommerce site
// @Tags Admin
// @Success 200 "Success"
// @Failure 400 "Failed"
// @Router /admin/logout [post]
func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "Logout Successfully",
		Data:       nil,
		Errors:     nil,
	})
}

// BlockUser
// @Summary Admin can block a user
// @ID block-user
// @Description Admin can block a  user
// @Tags Admin
// @Accept json
// @Produce json
// @Param input body request.BlockUser{} true "inputs"
// @Success 200 {object} respondse.Response
// @Failure 401 {object} respondse.Response
// @Failure 422 {object} respondse.Response
// @Failure 500 {object} respondse.Response
// @Router /admin/block [patch]
func (cr *AdminHandler) BlockUser(c *gin.Context) {
	var body request.BlockUser
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, respondse.Response{
			StatusCode: 422,
			Message:    "Unable to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminId, err := utilityHandler.GEtAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Can't find admin",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.AdminUsecase.BlockUser(body, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Can't Block user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "User Blocked",
		Data:       nil,
		Errors:     nil,
	})

}

// UnblockUser
// @Summary Admin can unblock a blocked user
// @ID unblock-user
// @Description Admin can unblock a blocked user
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be unblocked"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Failure 500 {object} respondse.Response
// @Router /admin/unblock/{user_id} [patch]
func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	paramsId := c.Param("user_id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.AdminUsecase.UnBlockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "cant unblock user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "user unblocked",
		Data:       nil,
		Errors:     nil,
	})
}

// @Summary FindAllUsers
// @Id FindAllUsers
// @Discription list of users.
// @tags Admin
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param perPage query int false "Number of items to retrieve per page"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /admin/findall [get]
func (cr *AdminHandler) FindAllUser(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "invalid pagenumber",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	pagePer, err := strconv.Atoi(c.Query("perPage"))
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "Invalid perPage",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	list := request.Pagination{
		Page:    uint(page),
		PerPage: uint(pagePer),
	}

	users, err := cr.AdminUsecase.FindAllUser(c.Request.Context(), list)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "user not found",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "List of users",
		Data:       users,
		Errors:     nil,
	})
	return
}

// FindUserByID
// @Summary Admin can fetch a specific  user details using id
// @ID find-user-by-id
// @Description Admin can fetch a specific user details using user id
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be fetched"
// @Success 200 {object} respondse.Response
// @Failure 422 {object} respondse.Response
// @Failure 500 {object} respondse.Response
// @Router /admin/finduser/{user_id} [get]
func (cr *AdminHandler) FindUserByID(c *gin.Context) {
	paramsID := c.Param("user_id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, respondse.Response{StatusCode: 422, Message: "failed to parse user id", Data: nil, Errors: err.Error()})
		return
	}
	user, err := cr.AdminUsecase.FindUserbyId(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, respondse.Response{StatusCode: 500, Message: "failed fetch user", Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200, Message: "Successfully fetched user details", Data: user, Errors: nil,
	})

}
