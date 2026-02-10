package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"ilicense-lite/bootstrap/logger"
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

func (this *LicenseService) LicenseGet(ctx context.Context, in *input.LicenseGetInput) (interface{}, error) {
	logger.ServiceLogger.WithContext(ctx).Infof("********%+v", "test")
	return this.licenseDao.LicenseGet(ctx, in.ID)
}
func (this *LicenseService) LicenseAdd(ctx context.Context, in *input.LicenseAddInput) (interface{}, error) {
	if in.IssuerID <= 0 {
		in.IssuerID = 1
	}
	if in.Code == "" {
		in.Code = fmt.Sprintf("LIC-%d", time.Now().UnixMilli())
	}
	expireAtTime, err := util.ParseDate(in.ExpireAt)
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

	issuer, err := this.issuerDao.IssuerGet(ctx, in.IssuerID)
	if err != nil {
		return nil, err
	}
	product, err := this.productDao.ProductGet(ctx, in.ProductID)
	if err != nil {
		return nil, err
	}
	customer, err := this.customerDao.CustomerGet(ctx, in.CustomerID)
	if err != nil {
		return nil, err
	}
	// 2. 构建License数据
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
		return nil, err
	}
	privateKey, err := util.GetPrivateKey(issuer.PrivateKey)
	if err != nil {
		return nil, err
	}
	activeCode, err := util.GenerateActivationCode(string(licenseDataBytes), privateKey)
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
	license, err := this.licenseDao.LicenseGet(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	issuer, err := this.issuerDao.IssuerGet(ctx, license.IssuerID)
	if err != nil {
		return nil, err
	}
	product, err := this.productDao.ProductGet(ctx, license.ProductID)
	if err != nil {
		return nil, err
	}
	customer, err := this.customerDao.CustomerGet(ctx, license.CustomerID)
	if err != nil {
		return nil, err
	}
	expireAtTime, err := util.ParseDate(in.ExpireAt)
	if err != nil {
		return nil, err
	}
	m := &model.License{
		Code:         fmt.Sprintf("LIC-%d", time.Now().UnixMilli()),
		ProductID:    license.ProductID,
		CustomerID:   license.CustomerID,
		IssuerID:     license.IssuerID,
		IssueAt:      time.Now(),
		ExpireAt:     expireAtTime,
		Modules:      license.Modules,
		MaxInstances: license.MaxInstances,
		Remarks:      in.Remarks,
	}

	// 2. 构建License数据
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
		return nil, err
	}
	privateKey, err := util.GetPrivateKey(issuer.PrivateKey)
	if err != nil {
		return nil, err
	}
	activeCode, err := util.GenerateActivationCode(string(licenseDataBytes), privateKey)
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
	log.Printf("清理后激活码长度: %d", len(cleaned))

	// 2. URL-safe Base64 解码（无 padding）
	decoded, err := base64.RawURLEncoding.DecodeString(cleaned)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}
	log.Printf("解码后长度: %d bytes", len(decoded))

	// 3. 解析数据结构
	buf := bytes.NewReader(decoded)

	// 读取数据长度
	var dataLen int32
	if err := binary.Read(buf, binary.BigEndian, &dataLen); err != nil {
		return nil, err
	}
	log.Printf("数据长度: %d bytes", dataLen)

	// 读取数据
	dataBytes := make([]byte, dataLen)
	if _, err := io.ReadFull(buf, dataBytes); err != nil {
		return nil, err
	}
	jsonData := string(dataBytes)
	log.Printf("License数据: %s", jsonData)

	// 读取签名长度
	var sigLen int32
	if err := binary.Read(buf, binary.BigEndian, &sigLen); err != nil {
		return nil, err
	}
	log.Printf("签名长度: %d bytes", sigLen)

	// 读取签名
	signature := make([]byte, sigLen)
	if _, err := io.ReadFull(buf, signature); err != nil {
		return nil, err
	}
	log.Printf("实际读取签名: %d bytes", len(signature))

	// 4. 加载公钥
	publicKey, err := util.LoadPublicKey(publicKeyBase64)
	if err != nil {
		return nil, err
	}

	// 5. 验证签名
	if err := util.VerifySignature(dataBytes, signature, publicKey); err != nil {
		log.Printf("签名验证失败: %v", err)
		return nil, errors.New("signature verification failed")
	}

	// 6. 解析 License JSON
	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataBytes, &dataMap); err != nil {
		return nil, err
	}

	logger.ServiceLogger.Infof("----------%+v", dataMap)
	issueAt, _ := util.ParseDate(dataMap["issue_at"].(string))
	expireAt, _ := util.ParseDate(dataMap["expire_at"].(string))
	// 7. 组装返回对象
	info := &output.LicenseInfo{
		CustomerCode: dataMap["customer_code"].(string),
		CustomerName: dataMap["customer_name"].(string),
		ProductCode:  dataMap["product_code"].(string),
		ProductName:  dataMap["product_name"].(string),
		IssuerCode:   dataMap["issuer_code"].(string),
		IssuerName:   dataMap["issuer_name"].(string),
		IssueAt:      issueAt,
		ExpireAt:     expireAt,
		Modules:      dataMap["modules"].(string),
		MaxInstances: uint64(dataMap["max_instances"].(float64)), // JSON number
	}

	// 8. 校验有效期
	if info.ExpireAt.Before(time.Now()) {
		return nil, errors.New("license expired")
	}

	return &output.LicenseActivateOutput{OK: true, LicenseInfo: info}, nil
}
