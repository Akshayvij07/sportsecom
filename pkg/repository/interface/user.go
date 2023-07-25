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
	FindUser(ctx context.Context, UsersId int) (domain.Users, error)
	UpdatePassword(ctx context.Context, UserID int, Password string) error
	OtpLogin(mbnum string) (int, error)
	AddAdress(ctx context.Context, UserID int, AddressReq request.AddressReq) (address domain.Address, err error)
	UpdateAddress(ctx context.Context, address domain.Address) error
	VeiwAdress(ctx context.Context, UserID int) ([]respondse.Address, error)
	FindAddressByUserDetails(ctx context.Context, address request.AddressReq, userID int) (domain.Address, error)
	SaveUserAddress(ctx context.Context, userAddress domain.UserAddress) error
	UpdateUserAdress(ctx context.Context, userAddress domain.UserAddress) error
	FindUserAddressById(ctx context.Context, userId int) (domain.Address, error)
	FindWishListItem(ctx context.Context, productID, userID uint) (domain.WishList, error)
	FindAllWishListItemsByUserID(ctx context.Context, userID uint) ([]respondse.Wishlist, error)
	SaveWishListItem(ctx context.Context, wishList domain.WishList) error
	RemoveWishListItem(ctx context.Context, wishList domain.WishList) error
	FindProduct(ctx context.Context, id uint) (respondse.Product, error)
	GetInvoice(ctx context.Context, UserId int) (respondse.Invoice, error)
}
