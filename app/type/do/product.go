package do

import "time"

type Product struct {
	ID          uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`         // 产品ID
	Name        string    `gorm:"column:name;" json:"name" example:"运营平台"`                        // 产品名称
	Code        string    `gorm:"column:code;" json:"code" example:"platform"`                        // 产品代码, 唯一标识
	Description string    `gorm:"column:description;" json:"description" example:"运营平台描述"`      // 用户密码
	Status      *uint8    `gorm:"column:status;default:1" json:"status" example:"1"`                  // 状态: 1启用, 0禁用
	CreatedAt   time.Time `gorm:"column:created_at;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*Product) TableName() string {
	return "product"
}
