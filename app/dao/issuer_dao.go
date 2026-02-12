package dao

import (
	"context"
	"errors"

	"ilicense-lite/type/input"

	"ilicense-lite/bootstrap/client"
	"ilicense-lite/type/model"

	"gorm.io/gorm"
)

type IssuerDao struct{}

func NewIssuerDao() *IssuerDao {
	return &IssuerDao{}
}
func (*IssuerDao) IssuerGet(ctx context.Context, id uint64) (*model.Issuer, error) {
	m := &model.Issuer{ID: id}
	if err := client.MysqlDemo.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}

func (*IssuerDao) IssuerDeleteOne(ctx context.Context, id uint64) error {
	if err := client.MysqlDemo.WithContext(ctx).Delete(&model.Issuer{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
func (*IssuerDao) IssuerDelete(ctx context.Context, ids []uint64) error {
	if err := client.MysqlDemo.WithContext(ctx).Delete(&model.Issuer{}, ids).Error; err != nil {
		return err
	}
	return nil
}
func (*IssuerDao) IssuerAdd(ctx context.Context, m *model.Issuer) error {
	if err := client.MysqlDemo.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*IssuerDao) IssuerList(ctx context.Context) ([]model.Issuer, error) {
	var items []model.Issuer
	if err := client.MysqlDemo.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (*IssuerDao) IssuerQuery(ctx context.Context, input *input.IssuerQueryInput) (items []model.Issuer, total int64, err error) {
	offset := (input.Page - 1) * input.Size
	tx := client.MysqlDemo.WithContext(ctx).Model(&model.Issuer{})
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
