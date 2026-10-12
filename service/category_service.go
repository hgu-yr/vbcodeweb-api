package service

import (
	"errors"

	"vbCodeWeb/model"
	"vbCodeWeb/repo"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// CategoryService 分类业务逻辑。
type CategoryService struct {
	categoryRepo repo.CategoryRepo
}

func NewCategoryService(categoryRepo repo.CategoryRepo) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

// List 分页查询分类。
func (s *CategoryService) List(req *request.PageReq) ([]model.Category, int64, error) {
	req.Normalize()
	return s.categoryRepo.List(req)
}

// GetByID 查询分类详情。
func (s *CategoryService) GetByID(id uint) (*model.Category, error) {
	return s.categoryRepo.GetByID(id)
}

// Create 新建分类。
func (s *CategoryService) Create(req *request.CategoryReq) (*model.Category, error) {
	c := &model.Category{
		Name:  req.Name,
		Slug:  req.Slug,
		Scope: req.Scope,
		Sort:  req.Sort,
	}
	if c.Scope == 0 {
		c.Scope = 1
	}
	if err := s.categoryRepo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

// Update 更新分类。
func (s *CategoryService) Update(req *request.CategoryReq) (*model.Category, error) {
	if req.ID == 0 {
		return nil, errors.New("id 不能为空")
	}
	c := &model.Category{
		Model: gorm.Model{ID: req.ID},
		Name:  req.Name,
		Slug:  req.Slug,
		Scope: req.Scope,
		Sort:  req.Sort,
	}
	if err := s.categoryRepo.Update(c); err != nil {
		return nil, err
	}
	return s.categoryRepo.GetByID(req.ID)
}

// Delete 删除分类。
func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
}
