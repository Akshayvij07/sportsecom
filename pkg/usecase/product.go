package usecase

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/Akshayvij07/ecommerce/pkg/utilityHandler"
)

type ProductUsecase struct {
	ProductRepo interfaces.ProductRepo
}

func NewProductUsecase(ProductRepo interfaces.ProductRepo) services.ProductUsecase {
	return &ProductUsecase{
		ProductRepo: ProductRepo,
	}
}

func (P *ProductUsecase) Addcategory(ctx context.Context, req request.Category) (respondse.Category, error) {
	addcatagory, err := P.ProductRepo.Addcategory(ctx, req)
	return addcatagory, err
}

func (P *ProductUsecase) UpdatCategory(ctx context.Context, category request.Category, id int) (respondse.Category, error) {
	updatedcatagory, err := P.ProductRepo.UpdatCategory(ctx, category, id)
	return updatedcatagory, err
}
func (P *ProductUsecase) DeleteCatagory(ctx context.Context, Id int) error {
	err := P.ProductRepo.DeleteCatagory(ctx, Id)
	return err
}

func (p *ProductUsecase) Listallcatagory(ctx context.Context) ([]respondse.Category, error) {
	Allcatagory, err := p.ProductRepo.Listallcatagory(ctx)

	return Allcatagory, err
}

func (p *ProductUsecase) ShowCatagory(ctx context.Context, Id int) (respondse.Category, error) {

	yourcategory, err := p.ProductRepo.ShowCatagory(ctx, Id)

	return yourcategory, err

}

func (p *ProductUsecase) SaveProduct(ctx context.Context, product request.Product) (respondse.Product, error) {

	sku:=utilityHandler.GenerateSKU()
	newproduct, err := p.ProductRepo.SaveProduct(ctx, product,sku)

	return newproduct, err
}

func (p *ProductUsecase) UpdateProduct(ctx context.Context, id int, product request.Product) (respondse.Product, error) {
	updateproduct, err := p.ProductRepo.UpdateProduct(ctx, id, product)
	return updateproduct, err

}

func (p *ProductUsecase) DeleteProduct(ctx context.Context, id int) error {

	err := p.ProductRepo.DeleteProduct(ctx, id)

	return err
}

func (p *ProductUsecase) ViewAllProducts(ctx context.Context, pagination request.Pagination) (products []respondse.Product, err error) {
	allProducts, err := p.ProductRepo.ViewAllProducts(ctx, pagination)
	return allProducts, err
}

func (p *ProductUsecase) VeiwProduct(ctx context.Context, id int) (respondse.Product, error) {
	product, err := p.ProductRepo.ViewProduct(ctx, id)
	return product, err
}
