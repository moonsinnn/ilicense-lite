package service

import (
	"context"
	"fmt"

	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/dao"
	"ilicense-lite/library/token"
	"ilicense-lite/type/do"
	"ilicense-lite/type/request"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}
func (this *UserService) UserLogin(ctx context.Context, in *request.UserLoginRequest) (interface{}, error) {
	u, err := this.userDao.UserQuery(ctx, in.Name, "", in.Password)
	if err != nil {
		return nil, err
	}
	if t, err := token.GenerateJWT(fmt.Sprintf("%d", u.ID)); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}
func (this *UserService) UserGet(ctx context.Context, in *request.UserGetRequest) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.userDao.UserGet(ctx, in.ID)
}
func (this *UserService) UserAdd(ctx context.Context, in *request.UserAddRequest) (interface{}, error) {
	m := &do.User{Name: in.Name, Email: in.Email, Password: in.Password}
	if err := this.userDao.UserAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *UserService) UserQuery(ctx context.Context, in *request.UserQueryRequest) (interface{}, error) {
	return this.userDao.UserList(ctx)
}
