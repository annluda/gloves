package sessions

import (
	"gloves/app/auth"
	"gloves/app/controllers"
	userRequest "gloves/app/requests/user"
	"gloves/pkg/flash"

	"github.com/gin-gonic/gin"
)

// Create 登录界面
func Create(c *gin.Context) {
	controllers.Render(c, "sessions/create.html", gin.H{
		//"back": c.Query("back"),
		"back": "/",
	})
}

// Store 登录 (创建新会话)
func Store(c *gin.Context) {
	// 验证参数并且获取用户
	userLoginForm := &userRequest.UserLoginForm{
		Name:     c.PostForm("name"),
		Password: c.PostForm("password"),
	}
	user, errors := userLoginForm.ValidateAndGetUser(c)

	if len(errors) != 0 || user == nil {
		flash.SaveValidateMessage(c, errors)
		controllers.RedirectToLoginPage(c)
		return
	}

	// 用户是否激活
	//if !user.IsActivated() {
	//	flash.NewWarningFlash(c, "你的账号未激活，请检查邮箱中的注册邮件进行激活。")
	//	controllers.RedirectRouter(c, "root")
	//	return
	//}

	auth.Login(c, user)
	flash.NewSuccessFlash(c, "登录成功！")

	// 返回上次访问的页面
	back := c.Query("back")
	if back != "" {
		controllers.Redirect(c, back, true)
		return
	}

	controllers.RedirectRouter(c, "users.show", user.ID)
}

// Destroy 登出 (销毁会话)
func Destroy(c *gin.Context) {
	auth.Logout(c)
	flash.NewSuccessFlash(c, "已成功退出！")
	controllers.RedirectToLoginPage(c)
}
