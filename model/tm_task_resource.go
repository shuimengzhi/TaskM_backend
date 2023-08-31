// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTaskResource = "tm_task_resource"

// TaskResource mapped from table <tm_task_resource>
type TaskResource struct {

	/*
		资源id

	*/
	TrID int32 `gorm:"column:tr_id;primaryKey;autoIncrement:true;comment:资源id" json:"tr_id"`
	/*
		绝对路径包含文件名

	*/
	TrAbsolutePath string `gorm:"column:tr_absolute_path;not null;comment:绝对路径包含文件名" json:"tr_absolute_path"`
	/*
		文件名

	*/
	TrFileName string `gorm:"column:tr_file_name;not null;comment:文件名" json:"tr_file_name"`
	/*
		任务id

	*/
	TrTaskID int32 `gorm:"column:tr_task_id;not null;comment:任务id" json:"tr_task_id"`
	/*
		创建时间

	*/
	TrCreateTime int32 `gorm:"column:tr_create_time;not null;comment:创建时间" json:"tr_create_time"`
}

// TableName TaskResource's table name
func (*TaskResource) TableName() string {
	return TableNameTaskResource
}