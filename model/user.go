package model

import (
	"gorm.io/gorm"
)

// User 管理员账号表。
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);uniqueIndex;not null;comment:用户名" json:"username"`
	Password string `gorm:"type:varchar(128);not null;comment:密码(bcrypt加密)" json:"-"`
	Nickname string `gorm:"type:varchar(64);comment:昵称" json:"nickname"`
	Role     string `gorm:"type:varchar(32);not null;default:admin;comment:角色" json:"role"`
	// Status 1:正常 0:禁用
	Status int8 `gorm:"type:tinyint;not null;default:1;comment:状态 1正常 0禁用" json:"status"`
}

// TableName 指定用户表名。
func (User) TableName() string {
	return "users"
}

// TableComment 指定用户表注释。
func (User) TableComment() string {
	return "管理员账号表"
}
