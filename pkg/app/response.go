package app

import (
	"github.com/gin-cli/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Gin struct {
	g *gin.Context
}

func NewRep(g *gin.Context) Gin {
	return Gin{g:g}
}

func (g Gin) Response(httpCode, responseCode int, message string, data interface{}) {
	g.g.JSON(httpCode, Response{
		Code:    responseCode,
		Message: message,
		Data:    data,
	})
}

//JSON means response success.
func (g Gin) JSON(code int, data interface{}) {
	g.Response(http.StatusOK, code, e.Message(code), data)
}

func (g Gin) Ok(code int) {
	g.Response(http.StatusOK, code, e.Message(code), nil)
}

func (g Gin) Fail() {
	g.Response(http.StatusInternalServerError, e.ERROR, e.Message(e.ERROR), nil)
}
