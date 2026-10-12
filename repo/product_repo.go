package repo

import (
	"vbCodeWeb/model"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// ProductRepo 工具/产品数据访问。
type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

// List 分页查询产品。
func (r *ProductRepo) List(req *request.PageReq) ([]model.Product, int64, error) {
	db := r.db.Model(&model.Product{})
	if req.Keyword != "" {
		kw := "%" + req.Keyword + "%"
		db = db.Where("name LIKE ? OR tagline LIKE ? OR description LIKE ?", kw, kw, kw)
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
	var list []model.Product
	err := db.Order("featured DESC, click_count DESC, id DESC").
		Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&list).Error
	return list, total, err
}

// GetByID 按主键查询产品。
func (r *ProductRepo) GetByID(id uint) (*model.Product, error) {
	p := &model.Product{}
	if err := r.db.First(p, id).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// GetBySlug 按唯一标识查询产品。
func (r *ProductRepo) GetBySlug(slug string) (*model.Product, error) {
	p := &model.Product{}
	if err := r.db.Where("slug = ?", slug).First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// Create 新建产品。
func (r *ProductRepo) Create(p *model.Product) error {
	return r.db.Create(p).Error
}

// Update 按指定字段更新产品。
func (r *ProductRepo) Update(p *model.Product) error {
	return r.db.Model(&model.Product{}).Where("id = ?", p.ID).
		Select("name", "slug", "logo", "tagline", "description", "website_url",
			"category_id", "pricing", "platforms", "status", "featured").
		Updates(p).Error
}

// Delete 删除产品。
func (r *ProductRepo) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

// IncrClick 点击数 +1。
func (r *ProductRepo) IncrClick(id uint) error {
	return r.db.Model(&model.Product{}).Where("id = ?", id).
		UpdateColumn("click_count", gorm.Expr("click_count + 1")).Error
}
