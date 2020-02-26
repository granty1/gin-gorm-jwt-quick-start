package v1

import (
	"github.com/gin-cli/model"
	"github.com/gin-cli/pkg/app"
	"github.com/gin-cli/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

//UserList return list of user.
func UserList(c *gin.Context) {
	rep := app.NewRep(c)
	rep.JSON(e.SUCCESS, model.GetUserList())
}

//TestRoles data
func TestRoles(c *gin.Context) {
	data := []model.BaseRole{
		model.BaseRole{
			RoleId:   "1",
			RoleName: "Test",
		},
		model.BaseRole{
			RoleId:   "2",
			RoleName: "Train",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func TestPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": Page{
			Name:  "Page",
			Url:   "hello.com",
			Title: "title",
		},
	})
}

type Page struct {
	Name  string
	Url   string
	Title string
}

func TestPages(c *gin.Context) {
	data := []Page{
		Page{
			Name:  "Krakend",
			Url:   "https://krakend.io",
			Title: "krakend page",
		},
		Page{
			Name:  "Grant",
			Url:   "grant.io",
			Title: "grant page",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
