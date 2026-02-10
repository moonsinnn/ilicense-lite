package model

import "time"

type Issuer struct {
	ID           uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`            // ID
	Name         string    `gorm:"column:name;" json:"name" example:"issuer-a"`                           // 名称
	Code         string    `gorm:"column:code;" json:"code" example:"issuer-a-code"`                      // 代码, 唯一标识
	Description  string    `gorm:"column:description;" json:"description" example:"issuer-a-description"` // 描述
	PublicKey    string    `gorm:"column:public_key;" json:"public_key" example:"xxx"`                    // 公钥
	PrivateKey   string    `gorm:"column:private_key;" json:"private_key" example:"xxx"`                  // 私钥
	KeyAlgorithm string    `gorm:"column:key_algorithm;" json:"key_algorithm" example:"xxx"`              // 加密算法
	KeySize      int       `gorm:"column:key_size;" json:"key_size" example:"10"`                         // 密钥长度
	Status       *uint8    `gorm:"column:status;default:1" json:"status" example:"1"`                     // 状态: 1启用, 0禁用
	CreatedAt    time.Time `gorm:"column:created_at;" json:"created_at" example:"2020-10-11T10:10:10"`    // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;" json:"updated_at" example:"2020-10-11T10:10:10"`    // 更新时间
}

func (*Issuer) TableName() string {
	return "issuer"
}
