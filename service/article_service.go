package service

import (
	"errors"

	"vbCodeWeb/model"
	"vbCodeWeb/repo"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// ArticleService 文章业务逻辑。
type ArticleService struct {
	articleRepo repo.ArticleRepo
}

func NewArticleService(articleRepo repo.ArticleRepo) *ArticleService {
	return &ArticleService{articleRepo: articleRepo}
}

// List 分页查询文章。
func (s *ArticleService) List(req *request.PageReq) ([]model.Article, int64, error) {
	req.Normalize()
	return s.articleRepo.List(req)
}

// GetByID 查询文章详情。
func (s *ArticleService) GetByID(id uint) (*model.Article, error) {
	return s.articleRepo.GetByID(id)
}

// GetBySlug 按唯一标识查询文章。
func (s *ArticleService) GetBySlug(slug string) (*model.Article, error) {
	return s.articleRepo.GetBySlug(slug)
}

// Create 新建文章。
func (s *ArticleService) Create(req *request.ArticleReq) (*model.Article, error) {
	a := toArticle(req)
	if err := s.articleRepo.Create(a); err != nil {
		return nil, err
	}
	return a, nil
}

// Update 更新文章。
func (s *ArticleService) Update(req *request.ArticleReq) (*model.Article, error) {
	if req.ID == 0 {
		return nil, errors.New("id 不能为空")
	}
	if err := s.articleRepo.Update(toArticle(req)); err != nil {
		return nil, err
	}
	return s.articleRepo.GetByID(req.ID)
}

// Delete 删除文章。
func (s *ArticleService) Delete(id uint) error {
	return s.articleRepo.Delete(id)
}

// IncrView 阅读量 +1。
func (s *ArticleService) IncrView(id uint) error {
	return s.articleRepo.IncrView(id)
}

func toArticle(req *request.ArticleReq) *model.Article {
	return &model.Article{
		Model:       gorm.Model{ID: req.ID},
		Title:       req.Title,
		Slug:        req.Slug,
		Summary:     req.Summary,
		Content:     req.Content,
		Cover:       req.Cover,
		CategoryID:  req.CategoryID,
		Status:      req.Status,
		PublishedAt: req.PublishedAt,
	}
}
