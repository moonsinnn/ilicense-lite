package output

import "ilicense-lite/type/model"

type (
	ProductQueryOutput struct {
		Total int64           `json:"total"`
		Items []model.Product `json:"items"`
	}
)
