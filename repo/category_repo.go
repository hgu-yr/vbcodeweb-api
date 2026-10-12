package repo

import (
	"vbCodeWeb/model"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// CategoryRepo 分类数据访问。
type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

// List 分页查询分类。
func (r *CategoryRepo) List(req *request.PageReq) ([]model.Category, int64, error) {
	db := r.db.Model(&model.Category{})
	if req.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	if req.Status != nil {
		db = db.Where("scope = ?", *req.Status)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var list []model.Category
	err := db.Order("sort DESC, id DESC").
		Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&list).Error
	return list, total, err
}

// GetByID 按主键查询分类。
func (r *CategoryRepo) GetByID(id uint) (*model.Category, error) {
	c := &model.Category{}
	if err := r.db.First(c, id).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// Create 新建分类。
func (r *CategoryRepo) Create(c *model.Category) error {
	return r.db.Create(c).Error
}

// Update 按指定字段更新分类。
func (r *CategoryRepo) Update(c *model.Category) error {
	return r.db.Model(&model.Category{}).Where("id = ?", c.ID).
		Select("name", "slug", "scope", "sort").
		Updates(c).Error
}

// Delete 删除分类。
func (r *CategoryRepo) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}
