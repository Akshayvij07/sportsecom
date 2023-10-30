package usecase

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	"github.com/Akshayvij07/ecommerce/pkg/repository/mockrepo"

	"github.com/Akshayvij07/ecommerce/pkg/helper/request"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      request.UserSign
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(request.UserSign)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}
func EqCreateUserParams(arg request.UserSign, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestUserSignup(t *testing.T) {

	constTime := time.Now()

	testData := []struct {
		name           string
		input          request.UserSign
		buildStub      func(userRepo *mockrepo.MockUserRepository, user request.UserSign)
		expectedOutput respondse.UserValue
		expectedError  error
	}{
		{
			name: "FailedToSaveUserOnDatabase",
			input: request.UserSign{
				Name:     "Akshay",
				Email:    "akshay@gmail.com",
				Phone:    "9562461825",
				Password: "akshay@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository, user request.UserSign) {

				userRepo.EXPECT().UserSignup(gomock.Any(), EqCreateUserParams(user, user.Password)).Times(1).
					Return(respondse.UserValue{}, errors.New("error on database"))
			},
			expectedOutput: respondse.UserValue{},
			expectedError:  errors.New("error on database"),
		},
		{
			name: "SuccessSignup",
			input: request.UserSign{
				Name:     "Akshay",
				Email:    "akshay@gmail.com",
				Phone:    "9562461825",
				Password: "akshay@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository, user request.UserSign) {
				userRepo.EXPECT().UserSignup(gomock.Any(), EqCreateUserParams(user, user.Password)).Times(1).
					Return(respondse.UserValue{
						ID:       1,
						Name:     "Akshay",
						Email:    "akshay@gmail.com",
						Password: "hashed password",
						Created:  constTime,
					}, nil)
			},
			expectedOutput: respondse.UserValue{
				ID:       1,
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Password: "hashed password",
				Created:  constTime,
			},
			expectedError: nil,
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRepo := mockrepo.NewMockUserRepository(ctrl)
			userUseCase := NewUserUseCase(userRepo)
			tt.buildStub(userRepo, tt.input)

			user, err := userUseCase.SignUp(context.TODO(), tt.input)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, user, tt.expectedOutput)
		})
	}

}

func TestLoginWithEmail(t *testing.T) {
	//NewController from gomock package returns a new controller for testing
	ctrl := gomock.NewController(t)

	// NewMockUserRepository creates a new mockRepo instance
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo)

	// testData is a slice of struct which holds multiple test cases
	testData := []struct {
		name              string
		input             request.Login
		buildStub         func(userRepo *mockrepo.MockUserRepository)
		isExpectingOutput bool
		expectedError     error
	}{
		{
			name: "get details from database",
			input: request.Login{
				Email:    "akshay123@gmail.com",
				Password: "akshay@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				userRepo.EXPECT().UserLogin(gomock.Any(), "akshay123@gmail.com").Times(1).
					Return(domain.Users{}, errors.New("no user found"))
			},
			isExpectingOutput: false,
			expectedError:     errors.New("no user found"),
		},

		{
			name: "blocked user",
			input: request.Login{
				Email:    "akshay123@gmail.com",
				Password: "akshay@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				userRepo.EXPECT().UserLogin(gomock.Any(), "akshay123@gmail.com").Times(1).
					Return(domain.Users{
						ID:        1,
						Email:     "akshay123@gmail.com",
						Password:  "akshay@123",
						IsBlocked: true,
					}, errors.New("user is blocked"))
			},
			isExpectingOutput: false,
			expectedError:     errors.New("user is blocked"),
		},

		{
			name: "blocked user",
			input: request.Login{
				Email:    "akshay123@gmail.com",
				Password: "akshay@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				userRepo.EXPECT().UserLogin(gomock.Any(), "akshay123@gmail.com").Times(1).
					Return(domain.Users{
						ID:        1,
						Email:     "akshay123@gmail.com",
						Password:  "akshay@123",
						IsBlocked: true,
					}, errors.New("user is blocked"))
			},
			isExpectingOutput: false,
			expectedError:     errors.New("user is blocked"),
		},

		{
			name: "successfull logged into sportsecom",
			input: request.Login{
				Email:    "akshay123@gmail.com",
				Password: "akshay@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte("akshay@123"), 10)
				if err != nil {
					t.Fatalf("")
				}
				userRepo.EXPECT().UserLogin(gomock.Any(), "akshay123@gmail.com").Times(1).
					Return(domain.Users{
						ID:        1,
						Email:     "akshay123@gmail.com",
						Password:  string(hashedPassword),
						IsBlocked: false,
					}, nil)
			},
			isExpectingOutput: true,
			expectedError:     nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(userRepo)
			tokenString, actualErr := userUseCase.Login(context.TODO(), tt.input)

			if tt.expectedError == nil {
				assert.Nil(t, actualErr)
			} else {
				assert.Equal(t, tt.expectedError, actualErr)
			}

			if tt.isExpectingOutput {
				assert.NotEmpty(t, tokenString)
			} else {
				assert.Empty(t, tokenString)
			}

		})
	}
}
