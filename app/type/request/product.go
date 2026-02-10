package request

import "ilicense-lite/type/do"

type (
	ProductGetRequest struct {
		ID uint64 `form:"id"`
	}
	ProductAddRequest struct {
		Code        string `json:"code" form:"code" binding:"required" example:"product-add-code"`
		Name        string `json:"name" form:"name" binding:"required" example:"product-a"`
		Description string `json:"description" form:"description" binding:"required" example:"product-add-description"`
	}
	ProductQueryRequest struct {
		Page   int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size   int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		Name   string `form:"name" json:"name" example:"product-a"`
		Code   string `form:"code" json:"code" example:"customer-a"`
		Status *uint8 `form:"status" json:"status" example:"1"`
	}

	ProductQueryResponse struct {
		Total int64        `json:"total"`
		Items []do.Product `json:"items"`
	}
)
