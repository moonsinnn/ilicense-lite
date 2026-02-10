package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/dao"
	"ilicense-lite/library/util"
	"ilicense-lite/type/input"
	"ilicense-lite/type/model"
	"ilicense-lite/type/output"
)

type IssuerService struct {
	IssuerDao *dao.IssuerDao
}

func NewIssuerService() *IssuerService {
	return &IssuerService{
		IssuerDao: dao.NewIssuerDao(),
	}
}

func (this *IssuerService) IssuerGet(ctx context.Context, in *input.IssuerGetInput) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.IssuerDao.IssuerGet(ctx, in.ID)
}

func (this *IssuerService) IssuerAdd(ctx context.Context, in *input.IssuerAddInput) (interface{}, error) {
	// 1. 生成 RSA 2048 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	publicKey := &privateKey.PublicKey

	// 2. 导出公钥（PKIX / X.509）
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	publicKeyBase64 := base64.StdEncoding.EncodeToString(pubBytes)

	// 3. 导出私钥（PKCS#8）
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privBytes)

	// 4. 加密私钥（关键步骤）
	encryptedPrivateKey, err := util.Encrypt(privateKeyBase64)
	if err != nil {
		return nil, err
	}

	m := &model.Issuer{
		Name:         in.Name,
		Code:         in.Code,
		Description:  in.Description,
		PublicKey:    publicKeyBase64,
		PrivateKey:   encryptedPrivateKey,
		KeyAlgorithm: "RSA",
		KeySize:      2048,
	}
	if err := this.IssuerDao.IssuerAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *IssuerService) IssuerQuery(ctx context.Context, in *input.IssuerQueryInput) (interface{}, error) {
	items, total, err := this.IssuerDao.IssuerQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return &output.IssuerQueryOutput{Total: total, Items: items}, nil
}
