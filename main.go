package main

import (
	"github.com/gin-gonic/gin"
	"learn-gin/app/router"
	"learn-gin/config/log"
	"learn-gin/config/toml"
	"net/http"
	"time"
)

func main() {
	log.InitLogger(toml.GetConfig().Log.Path, toml.GetConfig().Log.Level)
	log.Logger.Info("hahahah")
	log.Logger.Info("config", log.Any("config", toml.GetConfig()))

	r := gin.Default()
	r.Use(router.Cors())
	r.Use(router.Recovery)
	router.InitRouter(r)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if nil != err {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
}
