package routers

import (
	"github.com/gin-cli/middleware"
	v1 "github.com/gin-cli/router/api/v1/base_user"
	"github.com/gin-gonic/gin"
	"net/http"
)

//UsersRoutes to register routers about user.
func UsersRoutes(baseRoute *gin.RouterGroup) {
	baseRoute.GET("/roles", v1.TestRoles)
	baseRoute.GET("/pages", v1.TestPages)
	baseRoute.GET("/page", v1.TestPage)
	baseRoute.GET("/person/:id/:name", func(context *gin.Context) {
		var p Person
		err := context.ShouldBindUri(&p)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusOK, p)
	})

	userGroup := baseRoute.Group("/user").Use(middleware.JWT())
	{
		userGroup.GET("/list", v1.UserList)
	}
}

type Person struct {
	Name string `json:"name" uri:"name" binding:"required"`
	Id   int    `json:"id" uri:"id" binding:"required"`
}
