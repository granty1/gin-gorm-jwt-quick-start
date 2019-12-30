package v1

import (
	"github.com/gin-cli/model"
	"github.com/gin-cli/pkg/app"
	"github.com/gin-cli/pkg/e"
	"github.com/gin-gonic/gin"
)

//UserList return list of user.
func UserList(c *gin.Context) {
	rep := app.NewRep(c)
	rep.JSON(e.SUCCESS, model.GetUserList())
}
