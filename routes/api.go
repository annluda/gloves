package routes

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"gloves/app/controllers"
)

func registerApi(g *gin.Engine) {
	// 头像
	g.GET("/avatar/:id", avatar)
}

// 头像
func avatar(c *gin.Context) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	imageBase64 := ""
	if id == 0 {

	}
	imageBuffer, _ := base64.StdEncoding.DecodeString(imageBase64)
	c.Data(200, "image/jpeg", imageBuffer)
}
