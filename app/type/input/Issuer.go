package input

type (
	IssuerGetInput struct {
		ID uint64 `form:"id"`
	}
	IssuerAddInput struct {
		Code        string `json:"code" form:"code" binding:"required" example:"Issuer-add-code"`
		Name        string `json:"name" form:"name" binding:"required" example:"Issuer-a"`
		Description string `json:"description" form:"description" binding:"required" example:"Issuer-add-description"`
	}
	IssuerQueryInput struct {
		Page   int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size   int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		Name   string `form:"name" json:"name" example:"Issuer-a"`
		Code   string `form:"code" json:"code" example:"customer-a"`
		Status *uint8 `form:"status" json:"status" example:"1"`
	}
)
