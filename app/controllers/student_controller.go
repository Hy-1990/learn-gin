package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"learn-gin/app/pojo/req"
	"learn-gin/app/services"
	"learn-gin/config/log"
	"net/http"
)

type StudentController struct {
}

func (s StudentController) StudentAddOne(context *gin.Context) {
	var addStudentReq req.StudentAddReq
	log.Logger.Info("StudentAddOne接口")
	// 将前端穿过来的json数据绑定存储在这个实体类中，BindJSON()也能使用
	if err := context.ShouldBindJSON(&addStudentReq); err != nil {
		log.Logger.Panic("参数异常")
	}

	if _, err := json.Marshal(addStudentReq); err != nil {
		log.Logger.Panic("参数解析异常")
	}
	_rsp := services.StudentServ.AddStudent(&addStudentReq)
	context.JSON(http.StatusOK, _rsp)
}
