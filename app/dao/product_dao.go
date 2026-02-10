package dao

import (
	"context"
	"errors"
	"ilicense-lite/type/request"

	"ilicense-lite/bootstrap/client"
	"ilicense-lite/type/do"

	"gorm.io/gorm"
)

type ProductDao struct{}

func NewProductDao() *ProductDao {
	return &ProductDao{}
}
func (*ProductDao) ProductGet(ctx context.Context, id uint64) (*do.Product, error) {
	m := &do.Product{ID: id}
	if err := client.MysqlDemo.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
func (*ProductDao) ProductAdd(ctx context.Context, m *do.Product) error {
	if err := client.MysqlDemo.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*ProductDao) ProductList(ctx context.Context) ([]do.Product, error) {
	var items []do.Product
	if err := client.MysqlDemo.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (*ProductDao) ProductQuery(ctx context.Context, input *request.ProductQueryRequest) (items []do.Product, total int64, err error) {
	offset := (input.Page - 1) * input.Size
	tx := client.MysqlDemo.WithContext(ctx).Model(&do.Product{})
	// 动态条件
	if input.Name != "" {
		tx = tx.Where("name LIKE ?", input.Name+"%")
	}

	if input.Code != "" {
		tx = tx.Where("code = ?", input.Code)
	}

	// status 用指针判断
	if input.Status != nil {
		tx = tx.Where("status = ?", *input.Status)
	}

	// 总数
	if err = tx.Count(&total).Error; err != nil {
		return
	}

	err = tx.
		Order("id DESC").
		Offset(offset).
		Limit(input.Size).
		Find(&items).Error
	return
}
