package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/app/controllers"
)

func RedisRouter(r *gin.Engine) {
	r.GET("/redis/testRedisInit", controllers.RedisCtrl.TestRedisInit)
}
