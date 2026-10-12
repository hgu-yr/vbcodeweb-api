package model

import (
	"gorm.io/gorm"
)

// Category 分类表，统一管理工具/小程序/文章的分类。
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(64);not null;comment:分类名称" json:"name"`
	Slug string `gorm:"type:varchar(64);uniqueIndex;not null;comment:唯一标识" json:"slug"`
	// Scope 1:工具 2:小程序 3:文章
	Scope int8 `gorm:"type:tinyint;not null;default:1;comment:归属 1工具 2小程序 3文章" json:"scope"`
	Sort  int  `gorm:"type:int;not null;default:0;comment:排序(越大越靠前)" json:"sort"`
}

// TableName 指定分类表名。
func (Category) TableName() string {
	return "categories"
}

// TableComment 指定分类表注释。
func (Category) TableComment() string {
	return "分类表"
}
