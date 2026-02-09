package dao

import (
	"context"

	"ilicense-lite/bootstrap/client"
	"ilicense-lite/type/do"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}
func (*UserDao) UserGet(ctx context.Context, id int64) (*do.User, error) {
	m := &do.User{ID: id}
	if err := client.MysqlDemo.WithContext(ctx).First(m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
func (*UserDao) UserAdd(ctx context.Context, m *do.User) error {
	if err := client.MysqlDemo.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (*UserDao) UserList(ctx context.Context) ([]do.User, error) {
	var items []do.User
	if err := client.MysqlDemo.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (*UserDao) UserQuery(ctx context.Context, name, email, password string) (*do.User, error) {
	var m do.User
	db := client.MysqlDemo.WithContext(ctx).Model(&do.User{})
	if name != "" {
		db.Where("name = ?", name)
	}
	if email != "" {
		db.Where("email = ?", email)
	}
	if password != "" {
		db.Where("password = ?", password)
	}
	if err := db.Take(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
