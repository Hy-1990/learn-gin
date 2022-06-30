package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/app/controllers"
)

func StudentRouter(r *gin.Engine) {
	r.POST("/student/addOne", controllers.StudentCtrl.StudentAddOne)
	r.POST("/student/addOrUpdateStudent", controllers.StudentCtrl.AddOrUpdateStudent)
	r.GET("/student/selectAll", controllers.StudentCtrl.SelectAll)
	r.DELETE("/student/deleteById", controllers.StudentCtrl.DeleteById)
}
