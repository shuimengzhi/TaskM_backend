// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTaskTag = "tm_task_tag"

// TaskTag mapped from table <tm_task_tag>
type TaskTag struct {

	/*
		任务标签id

	*/
	TtID int32 `gorm:"column:tt_id;primaryKey;autoIncrement:true;comment:任务标签id" json:"tt_id"`
	/*
		任务标签名字

	*/
	TtName string `gorm:"column:tt_name;not null;comment:任务标签名字" json:"tt_name"`
	/*
		排序数字

	*/
	TtSort int32 `gorm:"column:tt_sort;not null;comment:排序数字" json:"tt_sort"`
	/*
		色号

	*/
	TtColor string `gorm:"column:tt_color;not null;comment:色号" json:"tt_color"`
	/*
		所属项目id

	*/
	TtProjectID int32 `gorm:"column:tt_project_id;not null;comment:所属项目id" json:"tt_project_id"`
}

// TableName TaskTag's table name
func (*TaskTag) TableName() string {
	return TableNameTaskTag
}
