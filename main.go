package main

import (
	"github.com/gin-gonic/gin"
	"learn-gin/app/router"
	"learn-gin/config/log"
	"learn-gin/config/toml"
)

func main() {
	log.InitLogger(toml.GetConfig().Log.Path, toml.GetConfig().Log.Level)
	log.Logger.Info("hahahah")
	log.Logger.Info("config", log.Any("config", toml.GetConfig()))

	r := gin.Default()
	router.InitRouter(r)

	r.Run(":8080")
}
