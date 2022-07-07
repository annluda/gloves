package status

import (
	"github.com/gin-gonic/gin"
	"gloves/app/controllers"
	userModel "gloves/app/models"
)

func backTo(c *gin.Context, currentUser *userModel.User) {
	back := c.DefaultPostForm("back", "")
	if back != "" {
		controllers.Redirect(c, back, true)
		return
	}

	controllers.RedirectRouter(c, "users.show", currentUser.ID)
}
