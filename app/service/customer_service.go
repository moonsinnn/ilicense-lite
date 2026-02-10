package service

import (
	"context"
	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/dao"
	"ilicense-lite/type/input"
	"ilicense-lite/type/model"
	"ilicense-lite/type/output"
)

type CustomerService struct {
	CustomerDao *dao.CustomerDao
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		CustomerDao: dao.NewCustomerDao(),
	}
}

func (this *CustomerService) CustomerGet(ctx context.Context, in *input.CustomerGetInput) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.CustomerDao.CustomerGet(ctx, in.ID)
}
func (this *CustomerService) CustomerAdd(ctx context.Context, in *input.CustomerAddInput) (interface{}, error) {
	m := &model.Customer{
		Name:    in.Name,
		Code:    in.Code,
		Contact: in.Contact,
		Phone:   in.Phone,
		Email:   in.Email,
		Address: in.Address,
	}
	if err := this.CustomerDao.CustomerAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *CustomerService) CustomerQuery(ctx context.Context, in *input.CustomerQueryInput) (interface{}, error) {
	items, total, err := this.CustomerDao.CustomerQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return &output.CustomerQueryOutput{Total: total, Items: items}, nil
}
