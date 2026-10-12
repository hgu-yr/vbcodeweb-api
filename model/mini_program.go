package model

import (
	"gorm.io/gorm"
)

// MiniProgram 用 vibe coding 开发的微信小程序案例表。
type MiniProgram struct {
	gorm.Model
	Name        string `gorm:"type:varchar(128);not null;comment:小程序名称" json:"name"`
	Slug        string `gorm:"type:varchar(128);uniqueIndex;not null;comment:唯一标识" json:"slug"`
	AppID       string `gorm:"type:varchar(64);comment:小程序AppID" json:"app_id"`
	Subject     string `gorm:"type:varchar(128);comment:开发主体/作者" json:"subject"`
	Cover       string `gorm:"type:varchar(255);comment:封面图地址" json:"cover"`
	MiniCodeImg string `gorm:"type:varchar(255);comment:小程序码图片地址" json:"mini_code_img"`
	Description string `gorm:"type:text;comment:介绍" json:"description"`
	BuilderTool string `gorm:"type:varchar(64);comment:使用的 vibe coding 工具" json:"builder_tool"`
	CategoryID  uint   `gorm:"index;comment:分类ID" json:"category_id"`
	GuideURL    string `gorm:"type:varchar(255);comment:引导/详情地址" json:"guide_url"`
	// Status 1:正常 0:下线
	Status int8 `gorm:"type:tinyint;not null;default:1;comment:状态 1正常 0下线" json:"status"`
	// Featured 1:精选 0:普通
	Featured   int8  `gorm:"type:tinyint;not null;default:0;comment:是否精选 1是 0否" json:"featured"`
	ClickCount int64 `gorm:"type:bigint;not null;default:0;comment:查看次数" json:"click_count"`
}

// TableName 指定小程序案例表名。
func (MiniProgram) TableName() string {
	return "mini_programs"
}

// TableComment 指定小程序案例表注释。
func (MiniProgram) TableComment() string {
	return "vibe coding 小程序案例表"
}
