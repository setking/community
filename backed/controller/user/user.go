package user

import (
	"errors"
	"myApp/utils"
	"net/http"
	"time"

	"myApp/dao"
	"myApp/forms"
	"myApp/global"
	"myApp/initialize"
	"myApp/middlewares"
	"myApp/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 错误信息汉化处理
func HandleValidateError(ctx *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": initialize.ReplaceString(errs.Translate(global.Trans)),
	})
}

// 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidateError(c, err)
		return
	}
	// 注册
	user, err := dao.Register(registerForm)
	if err != nil {
		zap.S().Errorf("[Register] 新建 【用户失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeInternalServerError)
		return
	}
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint64(user.UserID),
		UserName:    user.UserName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, // 过期时间
			Issuer:    "setking",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"user_name":  user.UserName,
		"token":      token,
		"expires_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

// 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	passwordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidateError(c, err)
		return
	}
	// 登录
	user, err := dao.Login(passwordLoginForm)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
		zap.S().Errorf("[Login] 用户 【登录失败】 %s", err.Error())
		return
	}
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint64(user.UserID),
		UserName:    user.UserName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, // 过期时间
			Issuer:    "setking",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"user_name":  user.UserName,
		"token":      token,
		"expires_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}
