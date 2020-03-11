package routers

import (
	v1 "github.com/gin-cli/router/api/v1/jwk"
	"github.com/gin-gonic/gin"
)

func JWKRoutes(baseRoute *gin.RouterGroup) {
	baseRoute.GET("/public", v1.TestPublicKey)
	baseRoute.GET("/private", v1.TestPrivateKey)
}
