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
	r.GET("/student/selectOutline", controllers.StudentCtrl.SelectOutline)
	r.POST("/student/updateExec", controllers.StudentCtrl.UpdateExec)
	r.GET("/student/selectByNamespace", controllers.StudentCtrl.SelectByNamespace)
	r.GET("/student/getSql", controllers.StudentCtrl.GetSql)
	r.GET("/student/testRow", controllers.StudentCtrl.TestRow)
	r.GET("/student/testError", controllers.StudentCtrl.TestError)
}
