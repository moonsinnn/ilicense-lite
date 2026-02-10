package output

import (
	"time"

	"ilicense-lite/type/model"
)

type (
	LicenseQueryOutput struct {
		Total int64           `json:"total"`
		Items []model.License `json:"items"`
	}

	LicenseInfo struct {
		IssuerCode   string    `json:"issuer_code"`
		IssuerName   string    `json:"issuer_name"`
		CustomerCode string    `json:"customer_code"`
		CustomerName string    `json:"customer_name"`
		ProductCode  string    `json:"product_code"`
		ProductName  string    `json:"product_name"`
		LicenseCode  string    `json:"license_code"`
		IssueAt      time.Time `json:"issue_at"`
		ExpireAt     time.Time `json:"expire_at"`
		Modules      string    `json:"modules"`
		MaxInstances uint64    `json:"max_instances"`
	}

	LicenseActivateOutput struct {
		*LicenseInfo
		OK bool `json:"ok"`
	}
)
