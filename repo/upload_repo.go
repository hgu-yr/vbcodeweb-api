package repo

import (
	"vbCodeWeb/model"

	"gorm.io/gorm"
)

// UploadRepo 用户上传图片数据访问。
type UploadRepo struct {
	db *gorm.DB
}

func NewUploadRepo(db *gorm.DB) *UploadRepo {
	return &UploadRepo{db: db}
}

// Create 保存一条上传记录。
func (r *UploadRepo) Create(u *model.Upload) error {
	return r.db.Create(u).Error
}

// GetByID 按主键读取上传记录（含二进制）。
func (r *UploadRepo) GetByID(id uint) (*model.Upload, error) {
	u := &model.Upload{}
	if err := r.db.First(u, id).Error; err != nil {
		return nil, err
	}
	return u, nil
}
