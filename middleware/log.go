package middleware

import (
	"bytes"
	"github.com/gin-cli/init/log"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"time"
)

//Logger as middleware for gin-cli
func Logger() gin.HandlerFunc {
	log := log.Log

	return func(c *gin.Context) {
		// request time
		start := time.Now()
		// request path
		path := c.Request.URL.Path
		// request ip
		clientIP := c.ClientIP()
		// method
		method := c.Request.Method
		// read request content
		req, _ := httputil.DumpRequest(c.Request, true)
		log.Infof(`| %s | %s | %s | %5s | %s\n`,
			`Request :`, method, clientIP, path, string(req))
		// replace writer
		cusWriter := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = cusWriter
		// handle request
		c.Next()
		// ending time
		end := time.Now()
		//execute time
		latency := end.Sub(start)
		statusCode := c.Writer.Status()

		log.Infof(`| %s | %3d | %13v | %s \n`,
			`Response:`,
			statusCode,
			latency,
			cusWriter.body.String())
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
