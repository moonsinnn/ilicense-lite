package dao

import (
	"context"
	"errors"
	"ilicense-lite/type/input"

	"ilicense-lite/bootstrap/client"
	"ilicense-lite/type/model"

	"gorm.io/gorm"
)

type LicenseDao struct{}

func NewLicenseDao() *LicenseDao {
	return &LicenseDao{}
}
func (*LicenseDao) LicenseGet(ctx context.Context, id uint64) (*model.License, error) {
	m := &model.License{ID: id}
	if err := client.MysqlDemo.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
func (*LicenseDao) LicenseAdd(ctx context.Context, m *model.License) error {
	if err := client.MysqlDemo.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*LicenseDao) LicenseList(ctx context.Context) ([]model.License, error) {
	var items []model.License
	if err := client.MysqlDemo.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (*LicenseDao) LicenseQuery(ctx context.Context, input *input.LicenseQueryInput) (items []model.License, total int64, err error) {
	offset := (input.Page - 1) * input.Size
	tx := client.MysqlDemo.WithContext(ctx).Model(&model.License{})
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
