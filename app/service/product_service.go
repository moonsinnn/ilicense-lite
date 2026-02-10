package service

import (
	"context"
	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/dao"
	"ilicense-lite/type/do"
	"ilicense-lite/type/request"
)

type ProductService struct {
	productDao *dao.ProductDao
}

func NewProductService() *ProductService {
	return &ProductService{
		productDao: dao.NewProductDao(),
	}
}

func (this *ProductService) ProductGet(ctx context.Context, in *request.ProductGetRequest) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.productDao.ProductGet(ctx, in.ID)
}
func (this *ProductService) ProductAdd(ctx context.Context, in *request.ProductAddRequest) (interface{}, error) {
	m := &do.Product{Name: in.Name, Code: in.Code, Description: in.Description}
	if err := this.productDao.ProductAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *ProductService) ProductQuery(ctx context.Context, in *request.ProductQueryRequest) (interface{}, error) {
	items, total, err := this.productDao.ProductQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return &request.ProductQueryResponse{total, items}, nil
}
