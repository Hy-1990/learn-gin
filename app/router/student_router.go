package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/app/controllers"
)

func StudentRouter(r *gin.Engine) {
	r.POST("/student/addOne", controllers.StudentCtrl.StudentAddOne)
}
