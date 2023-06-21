package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
)

type CouponUseCase interface {
	CreateCoupon(ctx context.Context, coupon domain.Coupon) error
	UpdateCouponById(ctx context.Context, CouponId int, coupon request.Coupon) (domain.Coupon, error)
	DeleteCoupon(ctx context.Context, CouponId int) (err error)
	ViewCoupon(ctx context.Context, couponID int) (domain.Coupon, error)
	ViewCoupons(ctx context.Context) ([]domain.Coupon, error)
	ApplyCoupontoCart(ctx context.Context, userID int, Code string) (float64, error)
}
