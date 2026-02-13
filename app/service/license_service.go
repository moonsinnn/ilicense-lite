package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"ilicense-lite/dao"
	"ilicense-lite/library/util"
	"ilicense-lite/type/input"
	"ilicense-lite/type/model"
	"ilicense-lite/type/output"

	"github.com/goccy/go-json"
)

type LicenseService struct {
	licenseDao  *dao.LicenseDao
	issuerDao   *dao.IssuerDao
	productDao  *dao.ProductDao
	customerDao *dao.CustomerDao
}

func NewLicenseService() *LicenseService {
	return &LicenseService{
		licenseDao:  dao.NewLicenseDao(),
		issuerDao:   dao.NewIssuerDao(),
		productDao:  dao.NewProductDao(),
		customerDao: dao.NewCustomerDao(),
	}
}
func (this *LicenseService) LicenseDeleteOne(ctx context.Context, in *input.LicenseDeleteOneInput) error {
	return this.licenseDao.LicenseDeleteOne(ctx, in.ID)
}
func (this *LicenseService) LicenseDelete(ctx context.Context, in *input.LicenseDeleteInput) error {
	return this.licenseDao.LicenseDelete(ctx, in.IDs)
}
func (this *LicenseService) LicenseGet(ctx context.Context, in *input.LicenseGetInput) (interface{}, error) {
	return this.licenseDao.LicenseGet(ctx, in.ID)
}
func (this *LicenseService) LicenseAdd(ctx context.Context, in *input.LicenseAddInput) (interface{}, error) {
	now := time.Now()
	if in.IssuerID <= 0 {
		in.IssuerID = 1
	}
	if in.Code == "" {
		in.Code = newLicenseCode(now)
	}
	expireAtTime, err := util.ParseDate(in.ExpireAt)
	if err != nil {
		return nil, err
	}
	m := &model.License{
		Code:         in.Code,
		ProductID:    in.ProductID,
		CustomerID:   in.CustomerID,
		IssuerID:     in.IssuerID,
		IssueAt:      now,
		ExpireAt:     expireAtTime,
		Modules:      in.Modules,
		MaxInstances: in.MaxInstances,
		Remarks:      in.Remarks,
	}

	activeCode, err := this.buildActivationCode(ctx, m)
	if err != nil {
		return nil, err
	}
	m.ActivationCode = activeCode
	if err := this.licenseDao.LicenseAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *LicenseService) LicenseRenew(ctx context.Context, in *input.LicenseRenewInput) (interface{}, error) {
	now := time.Now()
	license, err := this.licenseDao.LicenseGet(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	expireAtTime, err := util.ParseDate(in.ExpireAt)
	if err != nil {
		return nil, err
	}
	m := &model.License{
		Code:         newLicenseCode(now),
		ProductID:    license.ProductID,
		CustomerID:   license.CustomerID,
		IssuerID:     license.IssuerID,
		IssueAt:      now,
		ExpireAt:     expireAtTime,
		Modules:      license.Modules,
		MaxInstances: license.MaxInstances,
		Remarks:      in.Remarks,
	}

	activeCode, err := this.buildActivationCode(ctx, m)
	if err != nil {
		return nil, err
	}
	m.ActivationCode = activeCode
	if err := this.licenseDao.LicenseAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *LicenseService) LicenseQuery(ctx context.Context, in *input.LicenseQueryInput) (interface{}, error) {
	items, total, err := this.licenseDao.LicenseQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	return &output.LicenseQueryOutput{Total: total, Items: items}, nil
}

func (this *LicenseService) LicenseActivate(ctx context.Context, in *input.LicenseActivateInput) (interface{}, error) {
	if in.IssuerID <= 0 {
		in.IssuerID = 1
	}
	issuer, err := this.issuerDao.IssuerGet(ctx, in.IssuerID)
	if err != nil {
		return nil, err
	}
	result, err := ValidateActivationCode(in.Code, issuer.PublicKey)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ValidateActivationCode(
	activationCode string,
	publicKeyBase64 string,
) (interface{}, error) {

	// 1. 清理格式（去空格）
	cleaned := strings.ReplaceAll(activationCode, " ", "")
	cleaned = strings.TrimSpace(cleaned)
	if cleaned == "" {
		return nil, errors.New("activation code is empty")
	}

	// 2. URL-safe Base64 解码（无 padding）
	decoded, err := base64.RawURLEncoding.DecodeString(cleaned)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}

	// 3. 解析数据结构
	buf := bytes.NewReader(decoded)

	// 读取数据长度
	var dataLen int32
	if err := binary.Read(buf, binary.BigEndian, &dataLen); err != nil {
		return nil, err
	}
	if dataLen <= 0 || int(dataLen) > buf.Len() {
		return nil, errors.New("invalid activation payload length")
	}

	// 读取数据
	dataBytes := make([]byte, dataLen)
	if _, err := io.ReadFull(buf, dataBytes); err != nil {
		return nil, err
	}

	// 读取签名长度
	var sigLen int32
	if err := binary.Read(buf, binary.BigEndian, &sigLen); err != nil {
		return nil, err
	}
	if sigLen <= 0 || int(sigLen) > buf.Len() {
		return nil, errors.New("invalid activation signature length")
	}

	// 读取签名
	signature := make([]byte, sigLen)
	if _, err := io.ReadFull(buf, signature); err != nil {
		return nil, err
	}

	// 4. 加载公钥
	publicKey, err := util.LoadPublicKey(publicKeyBase64)
	if err != nil {
		return nil, err
	}

	// 5. 验证签名
	if err := util.VerifySignature(dataBytes, signature, publicKey); err != nil {
		return nil, errors.New("signature verification failed")
	}

	// 6. 解析 License JSON
	var info output.LicenseInfo
	if err := json.Unmarshal(dataBytes, &info); err != nil {
		return nil, err
	}
	if info.IssuerCode == "" || info.CustomerCode == "" || info.ProductCode == "" {
		return nil, errors.New("invalid license payload")
	}

	// 8. 校验有效期
	if info.ExpireAt.IsZero() || info.ExpireAt.Before(time.Now()) {
		return nil, errors.New("license expired")
	}

	return &output.LicenseActivateOutput{OK: true, LicenseInfo: &info}, nil
}

func newLicenseCode(now time.Time) string {
	return fmt.Sprintf("LIC-%d", now.UnixNano())
}

func (this *LicenseService) buildActivationCode(ctx context.Context, m *model.License) (string, error) {
	issuer, err := this.issuerDao.IssuerGet(ctx, m.IssuerID)
	if err != nil {
		return "", err
	}
	product, err := this.productDao.ProductGet(ctx, m.ProductID)
	if err != nil {
		return "", err
	}
	customer, err := this.customerDao.CustomerGet(ctx, m.CustomerID)
	if err != nil {
		return "", err
	}
	licenseData := &output.LicenseInfo{
		LicenseCode:  m.Code,
		CustomerCode: customer.Code,
		CustomerName: customer.Name,
		ProductCode:  product.Code,
		ProductName:  product.Name,
		IssuerCode:   issuer.Code,
		IssuerName:   issuer.Name,
		IssueAt:      m.IssueAt,
		ExpireAt:     m.ExpireAt,
		Modules:      m.Modules,
		MaxInstances: m.MaxInstances,
	}
	licenseDataBytes, err := json.Marshal(licenseData)
	if err != nil {
		return "", err
	}
	privateKey, err := util.GetPrivateKey(issuer.PrivateKey)
	if err != nil {
		return "", err
	}
	return util.GenerateActivationCode(string(licenseDataBytes), privateKey)
}
