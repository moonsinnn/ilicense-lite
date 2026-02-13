package dao

import (
	"context"
	"errors"

	"ilicense-lite/bootstrap/client"
	"ilicense-lite/type/model"

	"gorm.io/gorm"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (*UserDao) UserGetByID(ctx context.Context, id uint64) (*model.User, error) {
	m := &model.User{ID: id}
	if err := client.MysqlDemo.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}

func (*UserDao) UserGetByUsername(ctx context.Context, username string) (*model.User, error) {
	m := &model.User{}
	if err := client.MysqlDemo.WithContext(ctx).Where("username = ?", username).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}

func (*UserDao) UserAdd(ctx context.Context, m *model.User) error {
	if err := client.MysqlDemo.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*UserDao) UserUpdate(ctx context.Context, m *model.User) error {
	if err := client.MysqlDemo.WithContext(ctx).Save(m).Error; err != nil {
		return err
	}
	return nil
}
