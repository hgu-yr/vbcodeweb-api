package repo

import (
	"vbCodeWeb/model"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// SubmissionRepo 用户投稿数据访问。
type SubmissionRepo struct {
	db *gorm.DB
}

func NewSubmissionRepo(db *gorm.DB) *SubmissionRepo {
	return &SubmissionRepo{db: db}
}

// Create 新建投稿。
func (r *SubmissionRepo) Create(s *model.Submission) error {
	return r.db.Create(s).Error
}

// GetByID 按主键查询投稿。
func (r *SubmissionRepo) GetByID(id uint) (*model.Submission, error) {
	s := &model.Submission{}
	if err := r.db.First(s, id).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// List 分页查询投稿。
func (r *SubmissionRepo) List(req *request.PageReq) ([]model.Submission, int64, error) {
	db := r.db.Model(&model.Submission{})
	if req.Keyword != "" {
		kw := "%" + req.Keyword + "%"
		db = db.Where("name LIKE ? OR description LIKE ? OR contact LIKE ?", kw, kw, kw)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var list []model.Submission
	err := db.Order("status ASC, id DESC").
		Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&list).Error
	return list, total, err
}

// CountPending 统计待审核数量。
func (r *SubmissionRepo) CountPending() (int64, error) {
	var n int64
	err := r.db.Model(&model.Submission{}).Where("status = ?", model.SubmissionPending).Count(&n).Error
	return n, err
}

// Save 保存投稿变更。
func (r *SubmissionRepo) Save(s *model.Submission) error {
	return r.db.Save(s).Error
}

// Delete 删除投稿。
func (r *SubmissionRepo) Delete(id uint) error {
	return r.db.Delete(&model.Submission{}, id).Error
}
