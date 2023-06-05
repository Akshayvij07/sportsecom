package handler

import (
	"fmt"
	"net/http"

	"github.com/Akshayvij07/ecommerce/pkg/config"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase  services.OtpUseCase
	userUseCase services.UserUseCase
	cfg         config.Config
}

func NewOtpHandler(cfg config.Config, otpUseCase services.OtpUseCase, userUseCase services.UserUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase:  otpUseCase,
		userUseCase: userUseCase,
		cfg:         cfg,
	}
}

// SendOtp
// @Summary Send OTP to user's mobile
// @ID send-otp
// @Description Send OTP to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param user_mobile body  request.OTPreq true "User mobile number"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /otp/send [post]
func (cr *OtpHandler) SendOtp(c *gin.Context) {
	var phno request.OTPreq
	err := c.Bind(&phno)
	if err != nil {
		fmt.Println("e1")
		c.JSON(http.StatusUnprocessableEntity, respondse.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})

		return
	}

	Sid, err := cr.otpUseCase.SendOTP(c.Request.Context(), phno)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "creatingfailed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, respondse.Response{
		StatusCode: 201,
		Message:    "otp send",
		Data:       Sid,
		Errors:     nil,
	})
}

// ValidateOtp
// @Summary Validate the OTP to user's mobile
// @ID validate-otp
// @Description Validate the  OTP sent to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param otp body request.Otpverifier true "OTP sent to user's mobile number"
// @Success 200 {object} respondse.Response
// @Failure 400 {object} respondse.Response
// @Router /otp/verify [post]
func (cr *OtpHandler) ValidateOtp(c *gin.Context) {
	var otpDetails request.Otpverifier
	err := c.Bind(&otpDetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.otpUseCase.VerifyOTP(c, otpDetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	ss, err := cr.userUseCase.OtpLogin(otpDetails.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, respondse.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, respondse.Response{
		StatusCode: 200,
		Message:    "login successful",
		Data:       nil,
		Errors:     nil,
	})
}
