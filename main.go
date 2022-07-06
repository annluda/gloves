package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"

	"gloves/app/helpers"
	followerModel "gloves/app/models/follower"
	passwordResetModel "gloves/app/models/password_reset"
	statusModel "gloves/app/models/status"
	userModel "gloves/app/models/user"
	"gloves/config"
	"gloves/database"
	"gloves/routes"
	"gloves/routes/named"
)

func main() {

	// 初始化配置
	config.InitConfig()

	// gin config
	g := gin.New()
	setupGin(g)

	// db config
	db := database.InitDB()
	// db migrate
	db.AutoMigrate(
		&userModel.User{},
		&passwordResetModel.PasswordReset{},
		&statusModel.Status{},
		&followerModel.Follower{},
	)
	defer db.Close()

	// router register
	routes.Register(g)
	printRoute()

	// 启动
	fmt.Println("---server on----")
	if err := http.ListenAndServe(":1007", g); err != nil {
		log.Fatal("http server 启动失败", err)
	}
}

// 配置 gin
func setupGin(g *gin.Engine) {
	// 启动模式配置
	gin.SetMode(config.AppConfig.RunMode)

	// 项目静态文件配置
	g.Static("/"+config.AppConfig.PublicPath, config.AppConfig.PublicPath)
	// 网站logo
	g.StaticFile("/favicon.ico", config.AppConfig.PublicPath+"/favicon.ico")

	// 模板配置
	// 注册模板函数
	g.SetFuncMap(template.FuncMap{
		// 根据 laravel-mix 的 public/mix-manifest.json 生成静态文件 path
		"Mix": helpers.Mix,
		// 生成项目静态文件地址
		"Static": helpers.Static,
		// 获取命名路由的 path
		"Route":         named.G,
		"RelativeRoute": named.GR,
	})
	// 模板存储路径
	g.LoadHTMLGlob(config.AppConfig.ViewsPath + "/**/*")
}

// 打印命名路由
func printRoute() {
	// 只有非 release 时才可用该函数
	if config.AppConfig.RunMode == config.RunmodeRelease {
		return
	}

	named.PrintRoutes()
}
