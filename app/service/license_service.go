package service

import (
	"context"
	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/dao"
	"ilicense-lite/library/util"
	"ilicense-lite/type/input"
	"ilicense-lite/type/model"
	"ilicense-lite/type/output"
	"time"
)

type LicenseService struct {
	LicenseDao *dao.LicenseDao
}

func NewLicenseService() *LicenseService {
	return &LicenseService{
		LicenseDao: dao.NewLicenseDao(),
	}
}

func (this *LicenseService) LicenseGet(ctx context.Context, in *input.LicenseGetInput) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.LicenseDao.LicenseGet(ctx, in.ID)
}
func (this *LicenseService) LicenseAdd(ctx context.Context, in *input.LicenseAddInput) (interface{}, error) {
	expireAtTime, err := util.ParseExpireAt(in.ExpireAt)
	if err != nil {
		return nil, err
	}
	m := &model.License{
		Code:         in.Code,
		ProductID:    in.ProductID,
		CustomerID:   in.CustomerID,
		IssueAt:      time.Now(),
		ExpireAt:     expireAtTime,
		Modules:      in.Modules,
		MaxInstances: in.MaxInstances,
		Remarks:      in.Remarks,
	}
	if err := this.LicenseDao.LicenseAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *LicenseService) LicenseQuery(ctx context.Context, in *input.LicenseQueryInput) (interface{}, error) {
	items, total, err := this.LicenseDao.LicenseQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return &output.LicenseQueryOutput{total, items}, nil
}
