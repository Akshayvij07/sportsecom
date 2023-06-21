//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/Akshayvij07/ecommerce/pkg/api"
	handler "github.com/Akshayvij07/ecommerce/pkg/api/handler"
	config "github.com/Akshayvij07/ecommerce/pkg/config"
	db "github.com/Akshayvij07/ecommerce/pkg/db"
	repository "github.com/Akshayvij07/ecommerce/pkg/repository"
	usecase "github.com/Akshayvij07/ecommerce/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewproductRepository,
		repository.NewCartRepository,
		repository.NewOrderepository,
		repository.NewCouponrepo,
		usecase.NewUserUseCase,
		usecase.NewOtpUseCase,
		usecase.NewAdminUseCase,
		usecase.NewProductUsecase,
		usecase.NewCartUseCase,
		usecase.NewOrderUseCase,
		usecase.NewCouponUseCase,
		handler.NewUserHandler,
		handler.NewOtpHandler,
		handler.NewAdminHandler,
		handler.NewproductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewCoupenHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
