package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	"github.com/Akshayvij07/ecommerce/pkg/usecase/mockusecase"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	userUseCase := mockusecase.NewMockUserUseCase(ctrl)
	userHandler := NewUserHandler(userUseCase)

	testData := []struct {
		name             string
		input            request.UserSign
		buildStub        func(userUseCase mockusecase.MockUserUseCase)
		expectedCode     int
		expectedResponse respondse.Response
		expectedData     respondse.UserValue
		expectedError    error
	}{
		{
			name: "successful",
			input: request.UserSign{
				Name:     "Akshay",
				Email:    "akshay@gmail.com",
				Phone:    "+919562461825",
				Password: "akshay@123",
			},
			buildStub: func(userUseCase mockusecase.MockUserUseCase) {
				userUseCase.EXPECT().SignUp(gomock.Any(), request.UserSign{
					Name:     "Akshay",
					Email:    "akshay@gmail.com",
					Phone:    "+919562461825",
					Password: "akshay@123",
				}).Times(1).
					Return(respondse.UserValue{
						ID:       1,
						Name:     "Akshay",
						Email:    "akshay@gmail.com",
						Password: "akshay@123",
						Created:  time.Now(),
					}, nil)
			},
			expectedCode: 201,
			expectedResponse: respondse.Response{
				StatusCode: 201,
				Message:    "User signup successful",
				Data: respondse.UserValue{
					ID:      1,
					Name:    "Akshay",
					Email:   "akshay@gmail.com",
					Created: time.Now(),
				},
				Errors: nil,
			},
			expectedData: respondse.UserValue{
				ID:      1,
				Name:    "Akshay",
				Email:   "akshay@gmail.com",
				Created: time.Now(),
			},
			expectedError: nil,
		},
		{
			name: "duplicate user",
			input: request.UserSign{
				Name:     "Akshay",
				Email:    "akshay@gmail.com",
				Phone:    "+919562461825",
				Password: "akshay@123",
			},
			buildStub: func(userUseCase mockusecase.MockUserUseCase) {
				userUseCase.EXPECT().SignUp(gomock.Any(), request.UserSign{
					Name:     "Akshay",
					Email:    "akshay@gmail.com",
					Phone:    "+919562461825",
					Password: "akshay@123",
				}).Times(1).
					Return(respondse.UserValue{}, errors.New("user is already exist "))
			},
			expectedCode: 400,
			expectedResponse: respondse.Response{
				StatusCode: 400,
				Message:    "unable create account",
				Data:       respondse.UserValue{},
				Errors:     "user already exits",
			},
			expectedData:  respondse.UserValue{},
			expectedError: errors.New("user already exists"),
		},
	}

	for _, tc := range testData {

		t.Run(tc.name, func(t *testing.T) {

			tc.buildStub(*userUseCase)

			server := gin.New()
			server.POST("/SignUp", userHandler.SignUp)

			jsonData, err := json.Marshal(&tc.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockReq, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)

			responseRec := httptest.NewRecorder()

			server.ServeHTTP(responseRec, mockReq)

			//validate
			assert.Equal(t, tc.expectedCode, responseRec.Code)
		})
	}
}
