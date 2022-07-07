package followers

import (
	"gloves/app/controllers"
	followerModel "gloves/app/models"
	"gloves/app/policies"
	"gloves/pkg/flash"
	"gloves/routes/named"

	"github.com/gin-gonic/gin"
)

// Store 关注用户
func Store(c *gin.Context, currentUser *followerModel.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	if ok := policies.UserPolicyFollow(c, currentUser, id); !ok {
		return
	}

	isFollowing := false
	if id != int(currentUser.ID) {
		isFollowing = followerModel.IsFollowing(int(currentUser.ID), id)
	}

	if !isFollowing {
		if err := followerModel.DoFollow(currentUser.ID, uint(id)); err != nil {
			flash.NewDangerFlash(c, "关注失败: "+err.Error())
		}
	}

	controllers.Redirect(c, named.G("users.show", id)+"?page=1", false)
}

// Destroy 取消关注用户
func Destroy(c *gin.Context, currentUser *followerModel.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	if ok := policies.UserPolicyFollow(c, currentUser, id); !ok {
		return
	}

	isFollowing := false
	if id != int(currentUser.ID) {
		isFollowing = followerModel.IsFollowing(int(currentUser.ID), id)
	}

	if isFollowing {
		if err := followerModel.DoUnFollow(currentUser.ID, uint(id)); err != nil {
			flash.NewDangerFlash(c, "取消关注失败: "+err.Error())
		}
	}

	controllers.Redirect(c, named.G("users.show", id)+"?page=1", false)
}
