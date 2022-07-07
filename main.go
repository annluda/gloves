package main

import (
	"fmt"
	"gloves/app/models"
	"gloves/database"
	"gloves/pkg/logger"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"gloves/app/helpers"
	"gloves/config"
	"gloves/routes"
	"gloves/routes/named"
)

func main() {

	// 初始化配置
	config.InitConfig()
	logger.Init()
	database.InitDB()
	models.Migrate()

	// gin
	g := setupGin()

	// router register
	routes.Register(g)
	//printRoute()

	// 启动
	fmt.Println("gloves started.")
	if err := http.ListenAndServe(":1007", g); err != nil {
		logger.Fatalf("http server 启动失败", err)
	}
}

// 配置 gin
func setupGin() *gin.Engine {
	// 启动模式配置
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

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

	return g
}

// 打印命名路由
func printRoute() {
	// 只有非 release 时才可用该函数
	if config.AppConfig.RunMode == config.RunmodeRelease {
		return
	}

	named.PrintRoutes()
}
