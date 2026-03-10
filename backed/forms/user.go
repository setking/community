package forms

// 用户
type PasswordLoginForm struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=10"`
}
type RegisterForm struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=10"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password"`
	Email           string `form:"email" json:"email" binding:"required,custom_email"`
	Role            string `form:"role" json:"role"`
}
