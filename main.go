package main

import (
	"fmt"
	"github.com/gin-cli/config"
	"github.com/gin-cli/init/jwk"
	"github.com/gin-cli/init/log"
	"github.com/gin-cli/init/migrate"
	"github.com/gin-cli/init/router"
	"github.com/gin-cli/middleware"
	"github.com/gin-cli/model"
	"net/http"
	"time"
)

func main() {
	log.Init()
	jwk.Init()
	model.Init(config.Config.DB)
	migrate.RegisterTables()

	r := router.Init()
	//use logger middleware
	r.Use(middleware.Logger())

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Config.Sys.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Project start in http://127.0.0.1:%d success 🎉🎉🎉🎉\n",
		config.Config.Sys.Port)

	_ = s.ListenAndServe()
}
