package model

import (
	"time"

	"gorm.io/gorm"
)

// Article 资讯/教程/榜单文章表，用于 SEO 与公众号导流内容承接。
type Article struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null;comment:标题" json:"title"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex;not null;comment:唯一标识" json:"slug"`
	Summary     string `gorm:"type:varchar(512);comment:摘要" json:"summary"`
	Content     string `gorm:"type:longtext;comment:正文(Markdown)" json:"content"`
	Cover       string `gorm:"type:varchar(255);comment:封面图地址" json:"cover"`
	CategoryID  uint   `gorm:"index;comment:分类ID" json:"category_id"`
	// Status 1:发布 0:草稿
	Status      int8       `gorm:"type:tinyint;not null;default:0;comment:状态 1发布 0草稿" json:"status"`
	ViewCount   int64      `gorm:"type:bigint;not null;default:0;comment:阅读量" json:"view_count"`
	PublishedAt *time.Time `gorm:"comment:发布时间" json:"published_at"`
}

// TableName 指定文章表名。
func (Article) TableName() string {
	return "articles"
}

// TableComment 指定文章表注释。
func (Article) TableComment() string {
	return "文章表"
}
