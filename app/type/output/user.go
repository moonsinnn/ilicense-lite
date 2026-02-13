package output

import "time"

type UserProfileOutput struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginOutput struct {
	Token string             `json:"token"`
	User  *UserProfileOutput `json:"user"`
}
