package code

import "fmt"

type Code struct {
	Code    int    `json:"code"`    // 响应码
	Message string `json:"message"` // 响应消息
}

func (c *Code) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}

func New(code int, message string) error {
	return &Code{Code: code, Message: message}
}
