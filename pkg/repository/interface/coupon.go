package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
)

type CouponRepo interface {
	AddCoupon(ctx context.Context, Coupon domain.Coupon) (err error)
	UpdateCouponById(ctx context.Context, CouponId int, coupon request.Coupon) (updatedCoupon domain.Coupon, err error)
	DeleteCoupon(ctx context.Context, CouponId int) (err error)
	ViewCoupons(ctx context.Context) ([]domain.Coupon, error)
	ViewCoupon(ctx context.Context, couponID int) (domain.Coupon, error)
	GetByCode(ctx context.Context, couponCode string) (coupon domain.Coupon, err error)
	UpdateCouponByCode(ctx context.Context, code string, coupon domain.Coupon) error
	ApplyCoupontoCart(ctx context.Context, userID int, Code string) (float64, error)
}
