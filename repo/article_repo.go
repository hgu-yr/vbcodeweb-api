package repo

import (
	"vbCodeWeb/model"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// ArticleRepo 文章数据访问。
type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{db: db}
}

// List 分页查询文章。
func (r *ArticleRepo) List(req *request.PageReq) ([]model.Article, int64, error) {
	db := r.db.Model(&model.Article{})
	if req.Keyword != "" {
		kw := "%" + req.Keyword + "%"
		db = db.Where("title LIKE ? OR summary LIKE ?", kw, kw)
	}
	if req.CategoryID > 0 {
		db = db.Where("category_id = ?", req.CategoryID)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var list []model.Article
	err := db.Order("id DESC").
		Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&list).Error
	return list, total, err
}

// GetByID 按主键查询文章。
func (r *ArticleRepo) GetByID(id uint) (*model.Article, error) {
	a := &model.Article{}
	if err := r.db.First(a, id).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// GetBySlug 按唯一标识查询文章。
func (r *ArticleRepo) GetBySlug(slug string) (*model.Article, error) {
	a := &model.Article{}
	if err := r.db.Where("slug = ?", slug).First(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// Create 新建文章。
func (r *ArticleRepo) Create(a *model.Article) error {
	return r.db.Create(a).Error
}

// Update 按指定字段更新文章。
func (r *ArticleRepo) Update(a *model.Article) error {
	return r.db.Model(&model.Article{}).Where("id = ?", a.ID).
		Select("title", "slug", "summary", "content", "cover", "category_id", "status", "published_at").
		Updates(a).Error
}

// Delete 删除文章。
func (r *ArticleRepo) Delete(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}

// IncrView 阅读量 +1。
func (r *ArticleRepo) IncrView(id uint) error {
	return r.db.Model(&model.Article{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}
