package dao

import (
	"context"
	"errors"
	"ilicense-lite/type/input"

	"ilicense-lite/bootstrap/client"
	"ilicense-lite/type/model"

	"gorm.io/gorm"
)

type CustomerDao struct{}

func NewCustomerDao() *CustomerDao {
	return &CustomerDao{}
}
func (*CustomerDao) CustomerGet(ctx context.Context, id uint64) (*model.Customer, error) {
	m := &model.Customer{ID: id}
	if err := client.MysqlDemo.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
func (*CustomerDao) CustomerAdd(ctx context.Context, m *model.Customer) error {
	if err := client.MysqlDemo.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*CustomerDao) CustomerList(ctx context.Context) ([]model.Customer, error) {
	var items []model.Customer
	if err := client.MysqlDemo.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (*CustomerDao) CustomerQuery(ctx context.Context, input *input.CustomerQueryInput) (items []model.Customer, total int64, err error) {
	offset := (input.Page - 1) * input.Size
	tx := client.MysqlDemo.WithContext(ctx).Model(&model.Customer{})
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
