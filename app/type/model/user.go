package model

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`
	Username  string    `gorm:"column:username;uniqueIndex" json:"username" example:"admin"`
	Password  string    `gorm:"column:password" json:"-"`
	Name      string    `gorm:"column:name" json:"name" example:"管理员"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email" example:"admin@example.com"`
	Avatar    string    `gorm:"column:avatar" json:"avatar" example:""`
	Status    *uint8    `gorm:"column:status;default:1" json:"status" example:"1"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at" example:"2020-10-11T10:10:10"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at" example:"2020-10-11T10:10:10"`
}

func (*User) TableName() string {
	return "user"
}
