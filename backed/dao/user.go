package dao

import (
	"errors"
	"myApp/forms"
	"myApp/global"
	"myApp/initialize"
	"myApp/models"
	"myApp/utils"
)

// 注册用户
func Register(registerData forms.RegisterForm) (*models.User, error) {
	// 检查用户名是否存在
	if errs := checkUserNameExists(registerData.UserName); errs != nil {
		return nil, errs
	}

	// 检查邮箱是否存在
	if errs := checkEmailExists(registerData.Email); errs != nil {
		return nil, errs
	}
	// 生成user_id
	int64ID, err := initialize.GetID()
	if err != nil {
		return nil, err
	}

	rsp := &models.User{
		Email:    registerData.Email,
		UserName: registerData.UserName,
		Password: utils.GenMd5(registerData.Password),
		UserID:   int64ID,
	}
	res := global.DB.Create(rsp)
	if res.Error != nil {
		return nil, res.Error
	}
	return rsp, nil
}

// 检查用户名是否存在
func checkUserNameExists(userName string) error {
	var count int64
	err := global.DB.Model(&models.User{}).
		Where("user_name = ?", userName).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户名已存在")
	}

	return nil
}

// 检查邮箱是否存在
func checkEmailExists(email string) error {
	var count int64
	err := global.DB.Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("邮箱已被注册")
	}

	return nil
}

// 用户登录
func Login(LoginData forms.PasswordLoginForm) (*models.User, error) {
	var user models.User
	// 检查用户名是否存在
	result := global.DB.Where("user_name = ?", LoginData.UserName).First(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	// 验证密码
	passErr := utils.VerifyPassword(user.Password, LoginData.Password)
	if passErr == false {
		return nil, errors.New("密码不正确")
	}
	return &user, nil
}

// 通过id查询user
func GetUserByID(id int64) (*models.User, error) {
	var user models.User
	err := global.DB.Where("user_id =?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
