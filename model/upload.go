package model

import (
	"gorm.io/gorm"
)

// Upload 用户上传的图片。
// 直接存库（LONGBLOB）而不是落盘：容器文件系统在重新部署后会丢，
// 存库才能保证投稿记录里的图片地址长期有效。
type Upload struct {
	gorm.Model
	Filename string `gorm:"type:varchar(255);comment:原始文件名" json:"filename"`
	MimeType string `gorm:"type:varchar(64);comment:MIME 类型" json:"mime_type"`
	Size     int64  `gorm:"type:bigint;not null;default:0;comment:字节数" json:"size"`
	Data     []byte `gorm:"type:longblob;comment:图片二进制" json:"-"`
}

// TableName 指定上传表名。
func (Upload) TableName() string {
	return "uploads"
}

// TableComment 指定上传表注释。
func (Upload) TableComment() string {
	return "用户上传图片表"
}
