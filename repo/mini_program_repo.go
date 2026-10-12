package repo

import (
	"vbCodeWeb/model"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// MiniProgramRepo 小程序案例数据访问。
type MiniProgramRepo struct {
	db *gorm.DB
}

func NewMiniProgramRepo(db *gorm.DB) *MiniProgramRepo {
	return &MiniProgramRepo{db: db}
}

// List 分页查询小程序案例。
func (r *MiniProgramRepo) List(req *request.PageReq) ([]model.MiniProgram, int64, error) {
	db := r.db.Model(&model.MiniProgram{})
	if req.Keyword != "" {
		kw := "%" + req.Keyword + "%"
		db = db.Where("name LIKE ? OR description LIKE ? OR builder_tool LIKE ?", kw, kw, kw)
	}
	if req.CategoryID > 0 {
		db = db.Where("category_id = ?", req.CategoryID)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Featured != nil {
		db = db.Where("featured = ?", *req.Featured)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var list []model.MiniProgram
	err := db.Order("featured DESC, click_count DESC, id DESC").
		Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&list).Error
	return list, total, err
}

// GetByID 按主键查询小程序案例。
func (r *MiniProgramRepo) GetByID(id uint) (*model.MiniProgram, error) {
	m := &model.MiniProgram{}
	if err := r.db.First(m, id).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// GetBySlug 按唯一标识查询小程序案例。
func (r *MiniProgramRepo) GetBySlug(slug string) (*model.MiniProgram, error) {
	m := &model.MiniProgram{}
	if err := r.db.Where("slug = ?", slug).First(m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Create 新建小程序案例。
func (r *MiniProgramRepo) Create(m *model.MiniProgram) error {
	return r.db.Create(m).Error
}

// Update 按指定字段更新小程序案例。
func (r *MiniProgramRepo) Update(m *model.MiniProgram) error {
	return r.db.Model(&model.MiniProgram{}).Where("id = ?", m.ID).
		Select("name", "slug", "app_id", "subject", "cover", "mini_code_img",
			"description", "builder_tool", "category_id", "guide_url", "status", "featured").
		Updates(m).Error
}

// Delete 删除小程序案例。
func (r *MiniProgramRepo) Delete(id uint) error {
	return r.db.Delete(&model.MiniProgram{}, id).Error
}

// IncrClick 查看次数 +1。
func (r *MiniProgramRepo) IncrClick(id uint) error {
	return r.db.Model(&model.MiniProgram{}).Where("id = ?", id).
		UpdateColumn("click_count", gorm.Expr("click_count + 1")).Error
}
