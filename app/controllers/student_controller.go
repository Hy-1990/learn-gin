package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"learn-gin/app/pojo/req"
	"learn-gin/app/services"
	"learn-gin/config/log"
	"learn-gin/db_entity"
	"net/http"
	"strconv"
)

type StudentController struct {
}

//新增一个学生
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

//新增或者修改一个学生信息
func (s StudentController) AddOrUpdateStudent(context *gin.Context) {
	var addOrUpdateStudent db_entity.Student
	log.Logger.Info("AddOrUpdateStudent接口")
	if err := context.ShouldBindJSON(&addOrUpdateStudent); err != nil {
		log.Logger.Panic("参数异常")
	}

	if _, err := json.Marshal(addOrUpdateStudent); err != nil {
		log.Logger.Panic("参数解析异常")
	}
	_rsp := services.StudentServ.AddOrUpdateStudent(&addOrUpdateStudent)
	context.JSON(http.StatusOK, _rsp)
}

//查询所有学生
func (s StudentController) SelectAll(context *gin.Context) {
	log.Logger.Info("SelectAll接口")
	_rsp := services.StudentServ.SelectAll()
	context.JSON(http.StatusOK, _rsp)
}

//根据id删除学生
func (s StudentController) DeleteById(context *gin.Context) {
	log.Logger.Info("DeleteById接口")
	_id := context.Query("id")
	_a, _ := strconv.ParseInt(_id, 10, 64)
	_rsp := services.StudentServ.DeleteById(int32(_a))
	context.JSON(http.StatusOK, _rsp)
}
