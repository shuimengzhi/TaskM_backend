// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "tm_user"

// User mapped from table <tm_user>
type User struct {
	UID       int32  `gorm:"column:u_id;primaryKey;autoIncrement:true;comment:用户id" json:"u_id"`   // 用户id
	UAccount  string `gorm:"column:u_account;not null;comment:账户" json:"u_account"`                // 账户
	UNickName string `gorm:"column:u_nick_name;not null;comment:昵称" json:"u_nick_name"`            // 昵称
	UPassword string `gorm:"column:u_password;not null;comment:密码" json:"u_password"`              // 密码
	UStatus   bool   `gorm:"column:u_status;not null;default:1;comment:1:正常 2:禁用" json:"u_status"` // 1:正常 2:禁用
	UAvatar   string `gorm:"column:u_avatar;not null;comment:头像地址" json:"u_avatar"`                // 头像地址
	URoleType int8   `gorm:"column:u_role_type;not null;comment:1:管理员 2:普通人" json:"u_role_type"`   // 1:管理员 2:普通人
	/*
		创建时间
	*/
	UCreateTime int32 `gorm:"column:u_create_time;not null;comment:创建时间" json:"u_create_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}