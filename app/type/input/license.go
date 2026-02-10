package input

type (
	LicenseGetInput struct {
		ID uint64 `form:"id"`
	}
	LicenseAddInput struct {
		Code         string `json:"code" form:"code" binding:"required" example:"License-add-code"`
		ProductID    uint64 `json:"product_id" form:"product_id" binding:"required" example:"Product-add-product"`
		CustomerID   uint64 `json:"customer_id" form:"customer_id" binding:"required" example:"Customer-add-customer"`
		ExpireAt     string `json:"expire_at" form:"expire_at" binding:"required" example:"Customer-add-expire-at"`
		Modules      string `json:"modules" form:"modules"  example:"License-add-modules"`
		MaxInstances uint64 `json:"max_instances" form:"max_instances" example:"1"`
		Remarks      string `json:"remarks" form:"remarks" example:"remark"`
	}
	LicenseQueryInput struct {
		Page   int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size   int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		Name   string `form:"name" json:"name" example:"License-a"`
		Code   string `form:"code" json:"code" example:"customer-a"`
		Status *uint8 `form:"status" json:"status" example:"1"`
	}
)
