package model

import "time"

type License struct {
	ID             uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`         // ID
	Code           string    `gorm:"column:code;" json:"code" example:"license-a"`                       // 代码, 唯一标识
	ProductID      uint64    `gorm:"column:product_id;" json:"product_id" example:"1"`                   //	产品ID
	CustomerID     uint64    `gorm:"column:customer_id;" json:"customer_id" example:"1"`                 //	客户ID
	ActivationCode string    `gorm:"column:activation_code;" json:"activation_code" example:"xxxx-zzzz"` // 激活码(Base64 URL安全编码)
	IssueAt        time.Time `gorm:"column:issue_at;" json:"issue_at" example:"2020-10-11T10:10:10"`     // 签发日期
	ExpireAt       time.Time `gorm:"column:expire_at;" json:"expire_at" example:"2020-10-11T10:10:10"`   // 到期日期
	Modules        string    `gorm:"column:modules;" json:"modules" example:"order"`                     // 授权模块
	MaxInstances   uint64    `gorm:"column:max_instances;default:0" json:"max_instances" example:"-1"`   // 最大实例数, -1 表示无限
	Status         *uint8    `gorm:"column:status;default:1" json:"status" example:"1"`                  // 状态: 1启用, 0禁用
	Remarks        string    `gorm:"column:remarks;" json:"remarks" example:"备注"`                        // 备注
	CreatedAt      time.Time `gorm:"column:created_at;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*License) TableName() string {
	return "license"
}
