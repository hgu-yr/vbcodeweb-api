package model

import (
	"time"

	"gorm.io/gorm"
)

// Submission 用户投稿。审核通过后转成 Product 或 MiniProgram。
type Submission struct {
	gorm.Model
	// Kind 1:工具 2:小程序
	Kind        int8   `gorm:"type:tinyint;not null;default:1;index;comment:类型 1工具 2小程序" json:"kind"`
	Name        string `gorm:"type:varchar(128);not null;comment:名称" json:"name"`
	Slug        string `gorm:"type:varchar(128);comment:唯一标识(可空,审核时可补)" json:"slug"`
	Tagline     string `gorm:"type:varchar(255);comment:一句话简介" json:"tagline"`
	Description string `gorm:"type:text;comment:详细介绍" json:"description"`

	// 工具相关
	WebsiteURL string `gorm:"type:varchar(255);comment:官网地址" json:"website_url"`
	Pricing    string `gorm:"type:varchar(32);comment:收费模式" json:"pricing"`
	Platforms  string `gorm:"type:varchar(255);comment:支持平台" json:"platforms"`

	// 小程序相关
	AppID       string `gorm:"type:varchar(64);comment:小程序AppID" json:"app_id"`
	Subject     string `gorm:"type:varchar(128);comment:开发主体/作者" json:"subject"`
	BuilderTool string `gorm:"type:varchar(64);comment:使用的 vibe coding 工具" json:"builder_tool"`
	GuideURL    string `gorm:"type:varchar(255);comment:引导/详情地址" json:"guide_url"`

	CategoryID  uint   `gorm:"index;comment:分类ID" json:"category_id"`
	LogoURL     string `gorm:"type:varchar(255);comment:logo/主图地址" json:"logo_url"`
	CoverURL    string `gorm:"type:varchar(255);comment:封面图地址" json:"cover_url"`
	MiniCodeURL string `gorm:"type:varchar(255);comment:小程序码地址" json:"mini_code_url"`

	Contact string `gorm:"type:varchar(128);comment:联系方式" json:"contact"`

	// Status 0:待审核 1:通过 2:拒绝
	Status      int8       `gorm:"type:tinyint;not null;default:0;index;comment:状态 0待审核 1通过 2拒绝" json:"status"`
	ReviewNote  string     `gorm:"type:varchar(255);comment:审核备注" json:"review_note"`
	PublishedID uint       `gorm:"comment:通过后生成的条目ID" json:"published_id"`
	IP          string     `gorm:"type:varchar(64);comment:提交者IP" json:"ip"`
	ReviewedAt  *time.Time `gorm:"comment:审核时间" json:"reviewed_at"`
}

// TableName 指定投稿表名。
func (Submission) TableName() string {
	return "submissions"
}

// TableComment 指定投稿表注释。
func (Submission) TableComment() string {
	return "用户投稿表"
}

// 投稿类型。
const (
	SubmissionKindTool = 1
	SubmissionKindMini = 2
)

// 投稿状态。
const (
	SubmissionPending  = 0
	SubmissionApproved = 1
	SubmissionRejected = 2
)
