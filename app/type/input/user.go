package input

type (
	UserRegisterInput struct {
		Username string `json:"username" form:"username" binding:"required" example:"admin"`
		Password string `json:"password" form:"password" binding:"required" example:"admin123456"`
		Name     string `json:"name" form:"name" binding:"required" example:"管理员"`
		Email    string `json:"email" form:"email" binding:"required" example:"admin@example.com"`
	}

	UserLoginInput struct {
		Username string `json:"username" form:"username" binding:"required" example:"admin"`
		Password string `json:"password" form:"password" binding:"required" example:"admin123456"`
	}

	UserProfileUpdateInput struct {
		Name   string `json:"name" form:"name" example:"管理员"`
		Email  string `json:"email" form:"email" example:"admin@example.com"`
		Avatar string `json:"avatar" form:"avatar" example:""`
	}

	UserPasswordUpdateInput struct {
		OldPassword string `json:"old_password" form:"old_password" binding:"required"`
		NewPassword string `json:"new_password" form:"new_password" binding:"required"`
	}
)
