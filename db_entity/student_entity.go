package db_entity

import (
	"gorm.io/gorm"
	"learn-gin/app/constants"
	"time"
)

type Student struct {
	Id         int32           `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'id'"`
	Name       string          `json:"name" gorm:"column:name;type:varchar(255);comment:'名字'"`
	Age        int64           `json:"age" gorm:"column:age;comment:'年龄'"`
	Content    string          `json:"content" gorm:"column:content;type:varchar(255);comment:'描述'"`
	UpdateTime constants.HTime `json:"update_time" time_format:"unix" gorm:"column:update_time"`
	DelFlag    int64           `json:"del_flag" gorm:"column:del_flag;comment:'删除标识'"`
}

// 自定义表名
func (Student) TableName() string {
	return "student"
}

// 更新表更新时间为当前时间
func (u *Student) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("update_time", time.Now())
	return nil
}

// 新增表更新时间为当前时间
func (v *Student) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("update_time", time.Now())
	return nil
}
