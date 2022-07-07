package policies

import (
	"github.com/gin-gonic/gin"
	statusModel "gloves/app/models"
	"gloves/pkg/logger"
)

// StatusPolicyDestroy 是否有删除内容的权限
func StatusPolicyDestroy(c *gin.Context, currentUser *statusModel.User, status *statusModel.Status) bool {
	if currentUser.ID != status.UserID {
		logger.Infof("%s 没有权限删除 (ID: %d)", currentUser.Name, status.UserID)
		Unauthorized(c)
		return false
	}

	return true
}
