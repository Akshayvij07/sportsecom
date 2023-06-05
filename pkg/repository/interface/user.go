package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type UserRepository interface {
	UserSignup(ctx context.Context, user request.UserSign) (userData respondse.UserValue, err error)
	UserLogin(ctx context.Context, Email string) (domain.Users, error)
	FindUser(ctx context.Context,UsersId int)(domain.Users, error)
	UpdatePassword(ctx context.Context, UserID int,Password string) error
	OtpLogin(mbnum string) (int, error)
	AddAdress(ctx context.Context, UserID int, address request.AddressReq) (domain.Address, error)
	UpdateAdress(ctx context.Context, UserID int, address request.AddressReq) (domain.Address, error)
	VeiwAdress(ctx context.Context, UserID int) (domain.Address, error)
}
