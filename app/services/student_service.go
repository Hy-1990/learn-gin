package services

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
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
	SelectOutline()
	UpdateExec(req req.StudentUpdateExecReq)
	SelectByNamespace(age int64)
	TestRow()
	TestError()
	TestTransaction()
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

//查询所有学生简述信息
func (t StudentImpl) SelectOutline() rsp.ResponseMsg {
	log.Logger.Info("查询所有学生简述信息")
	_db := mysql.GetDB()
	var _result []constants.StudentOutline
	_db.Raw("select id,name,age from student where del_flag = 0").Scan(&_result)
	return *rsp.SuccessMsg(_result)
}

//使用exec进行数据更新
func (t StudentImpl) UpdateExec(req req.StudentUpdateExecReq) rsp.ResponseMsg {
	log.Logger.Info("使用exec模式数据更新")
	_db := mysql.GetDB()
	_db.Exec("UPDATE student SET name = ? WHERE id = ?", req.Name, req.Id)
	return *rsp.SuccessMsg("更新成功")
}

//使用命名参数查询
func (t StudentImpl) SelectByNamespace(age int64) rsp.ResponseMsg {
	log.Logger.Info("使用命名参数查询")
	_db := mysql.GetDB()
	var students []db_entity.Student
	_db.Where("age > @name and del_flag = @name2", sql.Named("name", age), sql.Named("name2", 0)).Find(&students)
	return *rsp.SuccessMsg(students)
}

//获取Sql语句
func (t StudentImpl) GetSql() rsp.ResponseMsg {
	log.Logger.Info("获取sql语句")
	_db := mysql.GetDB()
	_sql := _db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&db_entity.Student{}).Where("id > ?", 0).Limit(2).Order("age desc").Find(&[]db_entity.Student{})
	})
	fmt.Printf("sql is > %s\n", _sql)
	return *rsp.SuccessMsg("获取成功")
}

//测试row遍历方式
func (t StudentImpl) TestRow() rsp.ResponseMsg {
	log.Logger.Info("测试row遍历方式")
	_db := mysql.GetDB()
	var (
		_id   int32
		_name string
		_age  int64
	)
	_rows, _err := _db.Raw("select id,name,age from student where del_flag = 0").Rows()
	if _err != nil {
		log.Logger.Panic("执行sql异常", log.Any("error", _err.Error()))
	}
	defer _rows.Close()
	for _rows.Next() {
		_rows.Scan(&_id, &_name, &_age)
		fmt.Printf("student -> id=%v,name=%v,age=%v\n", _id, _name, _age)
	}
	return *rsp.SuccessMsg("测试成功")
}

//测试gorm异常
func (t StudentImpl) TestError() rsp.ResponseMsg {
	log.Logger.Info("测试gorm异常")
	_db := mysql.GetDB()
	var _student db_entity.Student
	if _err := _db.Where("del_flag = 1").First(&_student).Error; _err != nil {
		if errors.Is(_err, gorm.ErrRecordNotFound) {
			fmt.Println("error is ErrRecordNotFound")
		}
		log.Logger.Panic("error -> ", log.Any("error", _err))
	}
	log.Logger.Debug("student -> ", log.Any("student", _student))
	return *rsp.SuccessMsg("测试成功")
}

//测试事务效果
func (t StudentImpl) TestTransaction() rsp.ResponseMsg {
	log.Logger.Info("测试事务效果")
	_db := mysql.GetDB()
	_db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&db_entity.Student{
			Name: "张飞", Age: 200,
		})
		var _student db_entity.Student
		if _err := tx.Where("del_flag = 1").First(&_student).Error; _err != nil {
			return _err
		}
		fmt.Println(_student)
		return nil
	})
	return *rsp.SuccessMsg("测试成功")
}
