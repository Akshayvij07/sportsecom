package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type AdminRepository interface {
	FindAdmin(ctx context.Context, Email string) (domain.Admin, error)
	SaveAdmin(ctx context.Context, admin request.Admin) error
	FindAllUser(ctx context.Context, pagination request.Pagination) (users []respondse.UserValue, err error)
	BlockUser(body request.BlockUser, AdminId int) error
	UnBlockUser(id int) error
	FindUserbyId(ctx context.Context, userId int) (domain.Users, error)
}
