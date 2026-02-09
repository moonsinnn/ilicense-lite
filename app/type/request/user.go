package request

type (
	UserLoginRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	UserGetRequest struct {
		ID int64 `form:"id"`
	}
	UserAddRequest struct {
		Name     string `json:"name" form:"name" binding:"required" example:"bill"`           // 用户名称
		Email    string `json:"email" form:"email" binding:"required" example:"bill@126.com"` // 用户邮箱
		Password string `json:"password" form:"password" binding:"required" example:"111111"` // 用户密码
	}
	UserQueryRequest struct {
		Page int64 `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size int64 `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
	}
	UserSignBackRequest struct {
		Code  string `json:"code" form:"code" binding:"required" example:"0"`
		State int64  `json:"state" form:"state" binding:"required" example:"0"`
	}
)
