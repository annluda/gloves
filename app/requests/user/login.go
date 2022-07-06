package user

import (
	userModel "gloves/app/models/user"
	"gloves/pkg/flash"

	"github.com/gin-gonic/gin"
)

type UserLoginForm struct {
	Name     string
	Password string
}

func (u *UserLoginForm) ValidateAndGetUser(c *gin.Context) (user *userModel.User, errors []string) {

	// 通过用户名获取用户，并且判断密码是否正确
	user, err := userModel.GetByName(u.Name)
	if err != nil {
		errors = append(errors, "用户名不对哦")
		return nil, errors
	}

	if err := user.Compare(u.Password); err != nil {
		flash.NewDangerFlash(c, "密码不对哦")
		return nil, errors
	}

	return user, []string{}
}
