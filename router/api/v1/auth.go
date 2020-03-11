package v1

import (
	"github.com/gin-cli/init/log"
	"github.com/gin-cli/middleware"
	"github.com/gin-cli/model"
	"github.com/gin-cli/pkg/app"
	"github.com/gin-cli/pkg/e"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Plain    bool   `json:"plain"`
}

func Login(c *gin.Context) {
	rep := app.NewRep(c)
	var req LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		rep.Ok(e.INVALID_PARAMS)
		return
	}

	u, ok := model.GetUser(req.Username, req.Password)
	if !ok {
		rep.Ok(e.LOGIN_INFO_ERROR)
		return
	}
	claims := middleware.NewClaims(u.UUID, u.Username, u.RoleID)
	if req.Plain {
		rep.JSON(e.SUCCESS, claims)
	} else {
		token, err := claims.GenerateToken()
		if err != nil {
			log.Log.Error(err)
			rep.Ok(e.ERROR_AUTH_TOKEN)
			return
		}

		rep.JSON(e.SUCCESS, token)
	}
}
