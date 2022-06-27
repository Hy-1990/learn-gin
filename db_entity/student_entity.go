package db_entity

import "time"

type Student struct {
	Id         int32     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'id'"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(255);comment:'名字'"`
	Age        int64     `json:"age" gorm:"column:age;comment:'年龄'"`
	Content    string    `json:"content" gorm:"column:content;type:varchar(255);comment:'描述'"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
	DelFlag    int64     `json:"del_flag" gorm:"column:del_flag;comment:'删除标识'"`
}

// 自定义表名
func (Student) TableName() string {
	return "student"
}
