package policies

import (
	"gloves/app/controllers"

	"github.com/gin-gonic/gin"
)

// Unauthorized : 无权限时
func Unauthorized(c *gin.Context) {
	controllers.RenderUnauthorized(c)
}
