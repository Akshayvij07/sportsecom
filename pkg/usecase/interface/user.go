package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type UserUseCase interface {
	SignUp(ctx context.Context, user request.UserSign) (respondse.UserValue, error)
	Login(ctx context.Context, user request.Login) (string, error)
	OtpLogin(mobno string) (string, error)
	AddAdress(ctx context.Context, UserID int, address request.AddressReq, isDefault bool) error
	UpdateAdress(ctx context.Context, UserID int, address request.AddressReq) error
	VeiwAdress(ctx context.Context, UserID int) ([]respondse.Address, error)
	ChangePassword(ctx context.Context, User request.Password) error
	ListWishlist(ctx context.Context, userID uint) ([]respondse.Wishlist, error)
	RemoveFromWishList(ctx context.Context, wishList domain.WishList) error
	AddToWishList(ctx context.Context, wishList domain.WishList) error
	GetInvoice(ctx context.Context, UserID int) (Invoice respondse.Invoice, err error)
}
