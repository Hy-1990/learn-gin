package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/app/controllers"
)

func TestRouter(r *gin.Engine) {
	r.GET("/", controllers.TestCtrl.HelloWorld)
	r.GET("/test/:name", controllers.TestCtrl.TestParam)
	r.GET("/test1", controllers.TestCtrl.TestDefaultParam)
	r.POST("/testPost", controllers.TestCtrl.TestPost)
	r.POST("/testPost2", controllers.TestCtrl.TestPostBody)
}
