package routers

import (
	"github.com/gin-cli/middleware"
	"github.com/gin-cli/router/api/v1/base_user"
	"github.com/gin-gonic/gin"
)

//UsersRoutes to register routers about user.
func UsersRoutes(baseRoute *gin.RouterGroup) {
	userGroup := baseRoute.Group("/user").Use(middleware.JWT())
	{
		userGroup.GET("/list", v1.UserList)
	}
}
