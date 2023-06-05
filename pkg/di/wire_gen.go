// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Akshayvij07/ecommerce/pkg/api"
	"github.com/Akshayvij07/ecommerce/pkg/api/handler"
	"github.com/Akshayvij07/ecommerce/pkg/config"
	"github.com/Akshayvij07/ecommerce/pkg/db"
	"github.com/Akshayvij07/ecommerce/pkg/repository"
	"github.com/Akshayvij07/ecommerce/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	otpUseCase := usecase.NewOtpUseCase(cfg)
	otpHandler := handler.NewOtpHandler(cfg, otpUseCase, userUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUsecase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUsecase)
	productRepo := repository.NewproductRepository(gormDB)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewproductHandler(productUsecase)
	cartRepo:=repository.NewCartRepository(gormDB)
	cartUsecase:=usecase.NewCartUseCase(cartRepo)
	cartHandler:=handler.NewCartHandler(cartUsecase)
	orderRepo:=repository.NewOrderepository(gormDB)
	orderUsecase:=usecase.NewOrderUseCase(orderRepo,cartRepo)
	orderHandler:=handler.NewOrderHandler(orderUsecase)


	serverHTTP := http.NewServerHTTP(userHandler,otpHandler,adminHandler,productHandler,cartHandler,orderHandler)


	return serverHTTP, nil
}
