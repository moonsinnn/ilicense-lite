package input

type (
	CustomerGetInput struct {
		ID uint64 `form:"id"`
	}
	CustomerAddInput struct {
		Code    string `json:"code" form:"code" binding:"required" example:"Customer-add-code"`
		Name    string `json:"name" form:"name" binding:"required" example:"Customer-a"`
		Contact string `json:"contact" form:"contact"  example:"Customer-add-contact"`
		Phone   string `json:"phone" form:"phone" example:"Customer-add-phone"`
		Email   string `json:"email" form:"email" example:"Customer-add-email"`
		Address string `json:"address" form:"address" example:"Customer-add-address"`
	}
	CustomerQueryInput struct {
		Page   int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size   int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		Name   string `form:"name" json:"name" example:"Customer-a"`
		Code   string `form:"code" json:"code" example:"customer-a"`
		Status *uint8 `form:"status" json:"status" example:"1"`
	}
)
