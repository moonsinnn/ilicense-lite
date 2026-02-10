package output

import "ilicense-lite/type/model"

type (
	IssuerQueryOutput struct {
		Total int64          `json:"total"`
		Items []model.Issuer `json:"items"`
	}
)
