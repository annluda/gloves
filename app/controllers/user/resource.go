package user

import (
	"github.com/gin-gonic/gin"
	followerModel "gloves/app/models"

	"gloves/routes/named"

	"gloves/app/controllers"
	userRequest "gloves/app/requests/user"
	"gloves/app/services"
	viewmodels "gloves/app/view_models"
	"gloves/pkg/flash"
	"gloves/pkg/pagination"
)

// Index 用户列表
func Index(c *gin.Context, currentUser *followerModel.User) {
	defaultPageLine := 10

	allUserCount, err := followerModel.AllCount()
	if err != nil {
		flash.NewDangerFlash(c, "获取用户数据失败: "+err.Error())
		controllers.Redirect(c, named.G("users.index")+"?page=1", false)
		return
	}
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, defaultPageLine, allUserCount)

	if currentPage > pageTotalCount {
		controllers.Redirect(c, named.G("users.index")+"?page=1", false)
		return
	}

	users := services.UserListService(offset, limit)

	controllers.Render(c, "user/index.html",
		pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount, gin.H{
			"users": users,
		}))
}

// Create 创建用户页面
func Create(c *gin.Context) {
	controllers.Render(c, "user/create.html", gin.H{})
}

// Show 用户详情
func Show(c *gin.Context, currentUser *followerModel.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 如果要看的就是当前用户，那么就不用再去数据库中获取了
	user := currentUser
	if id != int(currentUser.ID) {
		user, err = followerModel.UserGet(id)
	}

	if err != nil || user == nil {
		controllers.Render404(c)
		return
	}

	// 获取分页参数
	statusesAllLength, _ := followerModel.GetUserAllStatusCount(int(user.ID))
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, 10, statusesAllLength)
	if currentPage > pageTotalCount {
		controllers.Redirect(c, named.G("users.show", id)+"?page=1", false)
		return
	}

	// 获取用户的内容
	statuses, _ := followerModel.GetUserStatus(int(user.ID), offset, limit)
	statusesViewModels := make([]*viewmodels.StatusViewModel, 0)
	for _, s := range statuses {
		statusesViewModels = append(statusesViewModels, viewmodels.NewStatusViewModelSerializer(s))
	}
	// 获取关注/粉丝
	followingsLength, _ := followerModel.FollowingsCount(id)
	followersLength, _ := followerModel.FollowersCount(id)
	isFollowing := false
	if id != int(currentUser.ID) {
		isFollowing = followerModel.IsFollowing(int(currentUser.ID), id)
	}

	controllers.Render(c, "user/show.html",
		pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount, gin.H{
			"userData":         viewmodels.NewUserViewModelSerializer(user),
			"statuses":         statusesViewModels,
			"statusesLength":   statusesAllLength,
			"followingsLength": followingsLength,
			"followersLength":  followersLength,
			"isFollowing":      isFollowing,
		}))
}

// Store 保存用户
func Store(c *gin.Context) {
	// 验证参数和创建用户
	userCreateForm := &userRequest.UserCreateForm{
		Name: c.PostForm("name"),
		//Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	user, errors := userCreateForm.ValidateAndSave()

	if len(errors) != 0 || user == nil {
		flash.SaveValidateMessage(c, errors)
		controllers.RedirectRouter(c, "users.create")
		return
	}

	controllers.RedirectRouter(c, "root")
}
