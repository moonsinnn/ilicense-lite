package model

import "time"

type Customer struct {
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`         // ID
	Name      string    `gorm:"column:name;" json:"name" example:"客户-A"`                            // 名称
	Code      string    `gorm:"column:code;" json:"code" example:"customer-a"`                      // 代码, 唯一标识
	Contact   string    `gorm:"column:contact;" json:"contact" example:"联系人-A"`                     // 联系人
	Phone     string    `gorm:"column:phone;" json:"phone" example:"13433332222"`                   // 联系电话
	Email     string    `gorm:"column:email;" json:"email" example:"134@126.com"`                   // 联系电话
	Address   string    `gorm:"column:address;" json:"address" example:"北京"`                        // 联系电话
	Status    *uint8    `gorm:"column:status;default:1" json:"status" example:"1"`                  // 状态: 1启用, 0禁用
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*Customer) TableName() string {
	return "customer"
}
