package routers

import (
	v1 "github.com/gin-cli/router/api/v1"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(baseRouters *gin.RouterGroup) {
	authGroup := baseRouters.Group("/auth")
	{
		authGroup.POST("/login", v1.Login)
	}
}
