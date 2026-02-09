package do

type User struct {
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT;column:id;" json:"id" example:"1"` // 用户ID
	Name     string `gorm:"column:name;" json:"name" example:"bill"`                     // 用户名称
	Email    string `gorm:"column:email;" json:"email" example:"bill@126.com"`           // 用户邮箱
	Password string `gorm:"column:password;" json:"password" example:"xlasiefska"`       // 用户密码
	ClientID string `gorm:"column:client_id;" json:"client_id" example:"xasdfsdf"`       // 客户端ID
}

func (*User) TableName() string {
	return "user"
}
