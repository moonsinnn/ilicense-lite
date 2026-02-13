package service

import (
	"context"
	"errors"
	"strconv"

	"ilicense-lite/dao"
	token2 "ilicense-lite/library/token"
	"ilicense-lite/type/input"
	"ilicense-lite/type/model"
	"ilicense-lite/type/output"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{userDao: dao.NewUserDao()}
}

func toUserProfileOutput(m *model.User) *output.UserProfileOutput {
	if m == nil {
		return nil
	}
	return &output.UserProfileOutput{
		ID:        m.ID,
		Username:  m.Username,
		Name:      m.Name,
		Email:     m.Email,
		Avatar:    m.Avatar,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (s *UserService) UserRegister(ctx context.Context, in *input.UserRegisterInput) (interface{}, error) {
	exists, err := s.userDao.UserGetByUsername(ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, errors.New("用户名已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	m := &model.User{
		Username: in.Username,
		Password: string(hashedPassword),
		Name:     in.Name,
		Email:    in.Email,
	}
	if err := s.userDao.UserAdd(ctx, m); err != nil {
		return nil, err
	}
	return toUserProfileOutput(m), nil
}

func (s *UserService) UserLogin(ctx context.Context, in *input.UserLoginInput) (interface{}, error) {
	m, err := s.userDao.UserGetByUsername(ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, errors.New("用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(in.Password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	token, err := token2.GenerateJWT(strconv.FormatUint(m.ID, 10))
	if err != nil {
		return nil, err
	}

	return &output.UserLoginOutput{Token: token, User: toUserProfileOutput(m)}, nil
}

func (s *UserService) UserProfileGet(ctx context.Context, userID uint64) (interface{}, error) {
	m, err := s.userDao.UserGetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, errors.New("用户不存在")
	}
	return toUserProfileOutput(m), nil
}

func (s *UserService) UserProfileUpdate(ctx context.Context, userID uint64, in *input.UserProfileUpdateInput) (interface{}, error) {
	m, err := s.userDao.UserGetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, errors.New("用户不存在")
	}

	if in.Name != "" {
		m.Name = in.Name
	}
	if in.Email != "" {
		m.Email = in.Email
	}
	if in.Avatar != "" {
		m.Avatar = in.Avatar
	}

	if err := s.userDao.UserUpdate(ctx, m); err != nil {
		return nil, err
	}
	return toUserProfileOutput(m), nil
}

func (s *UserService) UserPasswordUpdate(ctx context.Context, userID uint64, in *input.UserPasswordUpdateInput) error {
	m, err := s.userDao.UserGetByID(ctx, userID)
	if err != nil {
		return err
	}
	if m == nil {
		return errors.New("用户不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(in.OldPassword)); err != nil {
		return errors.New("旧密码不正确")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hashedPassword)
	return s.userDao.UserUpdate(ctx, m)
}
