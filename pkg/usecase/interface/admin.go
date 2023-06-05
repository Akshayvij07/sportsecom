package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type AdminUsecase interface {
	SaveAdmin(ctx context.Context, admin request.Admin) error
	LoginAdmin(ctx context.Context, admin request.AdminLogin) (string, error)
	FindAllUser(ctx context.Context, pagination request.Pagination) (users []respondse.UserValue, err error)
	BlockUser(body request.BlockUser, adminId int) error
	UnBlockUser(id int) error
	FindUserbyId(ctx context.Context, userID int) (domain.Users, error)
}
