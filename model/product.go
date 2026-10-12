package model

import (
	"gorm.io/gorm"
)

// Product vibe coding 工具/产品表。
type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(128);not null;comment:产品名称" json:"name"`
	Slug        string `gorm:"type:varchar(128);uniqueIndex;not null;comment:唯一标识" json:"slug"`
	Logo        string `gorm:"type:varchar(255);comment:Logo地址" json:"logo"`
	Tagline     string `gorm:"type:varchar(255);comment:一句话简介" json:"tagline"`
	Description string `gorm:"type:text;comment:详细介绍" json:"description"`
	WebsiteURL  string `gorm:"type:varchar(255);comment:官网地址" json:"website_url"`
	CategoryID  uint   `gorm:"index;comment:分类ID" json:"category_id"`
	Pricing     string `gorm:"type:varchar(32);comment:收费模式 免费/付费/免费增值" json:"pricing"`
	Platforms   string `gorm:"type:varchar(255);comment:支持平台(逗号分隔)" json:"platforms"`
	// Status 1:正常 0:下线
	Status int8 `gorm:"type:tinyint;not null;default:1;comment:状态 1正常 0下线" json:"status"`
	// Featured 1:精选 0:普通
	Featured   int8  `gorm:"type:tinyint;not null;default:0;comment:是否精选 1是 0否" json:"featured"`
	ClickCount int64 `gorm:"type:bigint;not null;default:0;comment:点击次数" json:"click_count"`
}

// TableName 指定产品表名。
func (Product) TableName() string {
	return "products"
}

// TableComment 指定产品表注释。
func (Product) TableComment() string {
	return "vibe coding 工具/产品表"
}
