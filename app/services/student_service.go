package services

import (
	"learn-gin/app/pojo/req"
	"learn-gin/app/pojo/rsp"
	"learn-gin/config/log"
	"learn-gin/config/mysql"
	"learn-gin/db_entity"
	"time"
)

type StudentService interface {
	AddStudent(req *req.StudentAddReq)
}

type StudentImpl struct {
}

// 添加学生
func (t StudentImpl) AddStudent(req *req.StudentAddReq) rsp.ResponseMsg {
	log.Logger.Info("添加学生参数:", log.Any("StudentReq", req))
	_db := mysql.GetDB()
	var _student db_entity.Student
	_student.Name = req.Name
	_student.Age = req.Age
	_student.Content = req.Content
	_student.UpdateTime = time.Now()
	_student.DelFlag = 0
	_db.Create(&_student)
	return *rsp.SuccessMsg("添加成功")
}
