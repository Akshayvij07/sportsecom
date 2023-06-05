package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

type ProductUsecase interface {
	Addcategory(ctx context.Context, req request.Category) (respondse.Category, error)
	UpdatCategory(ctx context.Context, category request.Category, id int) (respondse.Category, error)
	DeleteCatagory(ctx context.Context, Id int) error
	Listallcatagory(ctx context.Context) ([]respondse.Category, error)
	ShowCatagory(ctx context.Context, Id int) (respondse.Category, error)
	SaveProduct(ctx context.Context, product request.Product) (respondse.Product, error)
	UpdateProduct(ctx context.Context, id int, product request.Product) (respondse.Product, error)
	DeleteProduct(ctx context.Context, id int) error
	ViewAllProducts(ctx context.Context, pagination request.Pagination) (products []respondse.Product, err error)
	VeiwProduct(ctx context.Context, id int) (respondse.Product, error)
}
