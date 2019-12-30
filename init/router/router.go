package router

import (
	"github.com/gin-cli/middleware"
	"github.com/gin-cli/router/routers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	g := gin.Default()
	baseGroup := g.Group("/gin-cli")

	baseGroup.Use(middleware.Logger())

	routers.AuthRoutes(baseGroup)
	routers.UsersRoutes(baseGroup)

	return g
}
