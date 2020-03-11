package routers

import (
	"github.com/gin-cli/middleware"
	v1 "github.com/gin-cli/router/api/v1/base_user"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	roles = []Role{
		Role{
			Id:   1,
			Name: "Administrator",
		},
		Role{
			Id:   2,
			Name: "Manual User",
		},
	}

	persons = []Person{
		Person{
			Name:   "Grant",
			Id:     1,
			RoleId: 1,
		},
		Person{
			Name:   "Kral",
			Id:     2,
			RoleId: 2,
		},
	}
)

//UsersRoutes to register routers about user.
func UsersRoutes(baseRoute *gin.RouterGroup) {
	baseRoute.Use(middleware.JWT()).GET("/roles", v1.TestRoles)
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

	baseRoute.GET("/user/:name", func(context *gin.Context) {
		var p Person
		err := context.ShouldBindUri(&p)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		for _, v := range persons {
			if v.Name == p.Name {
				context.JSON(http.StatusOK, v)
				return
			}
		}

		context.JSON(http.StatusNotFound, nil)
	})

	baseRoute.GET("/role/:id", func(context *gin.Context) {
		var r Role
		err := context.ShouldBindUri(&r)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		for _, v := range roles {
			if v.Id == r.Id {
				context.JSON(http.StatusOK, v)
				return
			}
		}
		context.JSON(http.StatusNotFound, nil)
	})
}

type Person struct {
	Name   string `json:"name" uri:"name" binding:"required"`
	Id     int    `json:"id" uri:"id"`
	RoleId int    `json:"role_id" uri:"role_id"`
}

type Role struct {
	Id   int    `json:"id" uri:"id"`
	Name string `json:"name" uri:"name"`
}
