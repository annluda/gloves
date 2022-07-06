package status

import (
	"gloves/app/controllers"
	statusModel "gloves/app/models/status"
	userModel "gloves/app/models/user"
	"gloves/app/policies"
	"gloves/pkg/flash"

	"github.com/gin-gonic/gin"
)

// Store 创建
func Store(c *gin.Context, currentUser *userModel.User) {
	content := c.DefaultPostForm("content", "")
	contentLen := len(content)

	if contentLen == 0 {
		flash.NewDangerFlash(c, "内容不能为空")
		backTo(c, currentUser)
		return
	}

	//if contentLen > 140 {
	//  flash.NewDangerFlash(c, "内容长度不能超过 140 个字")
	//  backTo(c, currentUser)
	//  return
	//}

	status := &statusModel.Status{
		Content: content,
		UserID:  currentUser.ID,
	}
	if err := status.Create(); err != nil {
		flash.NewDangerFlash(c, "发布失败")
		backTo(c, currentUser)
		return
	}

	flash.NewSuccessFlash(c, "发布成功")
	backTo(c, currentUser)
}

// Destroy 删除
func Destroy(c *gin.Context, currentUser *userModel.User) {
	statusID, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	status, err := statusModel.Get(statusID)
	if err != nil {
		flash.NewDangerFlash(c, "删除失败")
		backTo(c, currentUser)
		return
	}

	// 权限判断
	if ok := policies.StatusPolicyDestroy(c, currentUser, status); !ok {
		return
	}

	// 删除
	if err := statusModel.Delete(int(status.ID)); err != nil {
		flash.NewDangerFlash(c, "删除失败")
		backTo(c, currentUser)
		return
	}

	flash.NewSuccessFlash(c, "删除成功")
	backTo(c, currentUser)
}
