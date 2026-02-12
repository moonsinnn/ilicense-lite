package dao

import (
	"context"
	"errors"
	"ilicense-lite/type/input"

	"ilicense-lite/bootstrap/client"
	"ilicense-lite/type/model"

	"gorm.io/gorm"
)

type ProductDao struct{}

func NewProductDao() *ProductDao {
	return &ProductDao{}
}
func (*ProductDao) ProductGet(ctx context.Context, id uint64) (*model.Product, error) {
	m := &model.Product{ID: id}
	if err := client.MysqlDemo.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
func (*ProductDao) ProductDeleteOne(ctx context.Context, id uint64) error {
	if err := client.MysqlDemo.WithContext(ctx).Delete(&model.Product{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
func (*ProductDao) ProductDelete(ctx context.Context, ids []uint64) error {
	if err := client.MysqlDemo.WithContext(ctx).Delete(&model.Product{}, ids).Error; err != nil {
		return err
	}
	return nil
}
func (*ProductDao) ProductAdd(ctx context.Context, m *model.Product) error {
	if err := client.MysqlDemo.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*ProductDao) ProductList(ctx context.Context) ([]model.Product, error) {
	var items []model.Product
	if err := client.MysqlDemo.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (*ProductDao) ProductQuery(ctx context.Context, input *input.ProductQueryInput) (items []model.Product, total int64, err error) {
	offset := (input.Page - 1) * input.Size
	tx := client.MysqlDemo.WithContext(ctx).Model(&model.Product{})
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
