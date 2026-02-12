package service

import (
	"context"
	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/dao"
	"ilicense-lite/type/input"
	"ilicense-lite/type/model"
	"ilicense-lite/type/output"
)

type ProductService struct {
	productDao *dao.ProductDao
}

func NewProductService() *ProductService {
	return &ProductService{
		productDao: dao.NewProductDao(),
	}
}
func (this *ProductService) ProductDeleteOne(ctx context.Context, in *input.ProductDeleteOneInput) error {
	return this.productDao.ProductDeleteOne(ctx, in.ID)
}
func (this *ProductService) ProductDelete(ctx context.Context, in *input.ProductDeleteInput) error {
	return this.productDao.ProductDelete(ctx, in.IDs)
}
func (this *ProductService) ProductGet(ctx context.Context, in *input.ProductGetInput) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.productDao.ProductGet(ctx, in.ID)
}
func (this *ProductService) ProductAdd(ctx context.Context, in *input.ProductAddInput) (interface{}, error) {
	m := &model.Product{
		Name:        in.Name,
		Code:        in.Code,
		Description: in.Description,
	}
	if err := this.productDao.ProductAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *ProductService) ProductQuery(ctx context.Context, in *input.ProductQueryInput) (interface{}, error) {
	items, total, err := this.productDao.ProductQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return &output.ProductQueryOutput{Total: total, Items: items}, nil
}
