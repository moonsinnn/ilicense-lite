package output

import "ilicense-lite/type/model"

type (
	LicenseQueryOutput struct {
		Total int64           `json:"total"`
		Items []model.License `json:"items"`
	}
)
