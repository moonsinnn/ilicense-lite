package output

import "ilicense-lite/type/model"

type (
	CustomerQueryOutput struct {
		Total int64            `json:"total"`
		Items []model.Customer `json:"items"`
	}
)
