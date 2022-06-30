package services

import (
	"fmt"
	"learn-gin/app/constants"
	"learn-gin/app/pojo/req"
	"learn-gin/app/pojo/rsp"
	"learn-gin/config/log"
	"learn-gin/config/mysql"
	"learn-gin/db_entity"
	"time"
)

type StudentService interface {
	AddStudent(req *req.StudentAddReq)
	AddOrUpdateStudent(student *db_entity.Student)
	SelectAll()
	DeleteById(id int32)
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
	_student.UpdateTime = constants.HTime{time.Now()}
	_student.DelFlag = 0
	_db.Create(&_student)
	return *rsp.SuccessMsg("添加成功")
}

//新增或者更新学生
func (t StudentImpl) AddOrUpdateStudent(student *db_entity.Student) rsp.ResponseMsg {
	log.Logger.Info("新增或者更新学生参数:", log.Any("Student", student))
	_db := mysql.GetDB()
	if student.Id != 0 {
		fmt.Println(student.UpdateTime)
		_db.Model(&student).Updates(student)
	} else {
		_db.Create(&student)
	}
	return *rsp.SuccessMsg("操作成功")
}

//查询所有学生
func (t StudentImpl) SelectAll() rsp.ResponseMsg {
	log.Logger.Info("查询所有学生")
	_db := mysql.GetDB()
	var _result []db_entity.Student
	_db.Where("del_flag = ?", 0).Find(&_result)
	return *rsp.SuccessMsg(_result)
}

//根据id删除学生
func (t StudentImpl) DeleteById(id int32) rsp.ResponseMsg {
	log.Logger.Info("根据id删除学生")
	_db := mysql.GetDB()
	_db.Delete(&db_entity.Student{}, id)
	return *rsp.SuccessMsg("删除成功")
}
