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
	customerDao *dao.CustomerDao
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		customerDao: dao.NewCustomerDao(),
	}
}
func (this *CustomerService) CustomerDeleteOne(ctx context.Context, in *input.CustomerDeleteOneInput) error {
	return this.customerDao.CustomerDeleteOne(ctx, in.ID)
}
func (this *CustomerService) CustomerDelete(ctx context.Context, in *input.CustomerDeleteInput) error {
	return this.customerDao.CustomerDelete(ctx, in.IDs)
}
func (this *CustomerService) CustomerGet(ctx context.Context, in *input.CustomerGetInput) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.customerDao.CustomerGet(ctx, in.ID)
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
	if err := this.customerDao.CustomerAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *CustomerService) CustomerQuery(ctx context.Context, in *input.CustomerQueryInput) (interface{}, error) {
	items, total, err := this.customerDao.CustomerQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return &output.CustomerQueryOutput{Total: total, Items: items}, nil
}
