package service

import (
	"errors"

	"vbCodeWeb/model"
	"vbCodeWeb/repo"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// ProductService 工具/产品业务逻辑。
type ProductService struct {
	productRepo repo.ProductRepo
}

func NewProductService(productRepo repo.ProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

// List 分页查询产品。
func (s *ProductService) List(req *request.PageReq) ([]model.Product, int64, error) {
	req.Normalize()
	return s.productRepo.List(req)
}

// GetByID 查询产品详情。
func (s *ProductService) GetByID(id uint) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}

// GetBySlug 按唯一标识查询产品。
func (s *ProductService) GetBySlug(slug string) (*model.Product, error) {
	return s.productRepo.GetBySlug(slug)
}

// Create 新建产品。
func (s *ProductService) Create(req *request.ProductReq) (*model.Product, error) {
	p := toProduct(req)
	if p.Status == 0 {
		p.Status = 1
	}
	if err := s.productRepo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Update 更新产品。
func (s *ProductService) Update(req *request.ProductReq) (*model.Product, error) {
	if req.ID == 0 {
		return nil, errors.New("id 不能为空")
	}
	if err := s.productRepo.Update(toProduct(req)); err != nil {
		return nil, err
	}
	return s.productRepo.GetByID(req.ID)
}

// Delete 删除产品。
func (s *ProductService) Delete(id uint) error {
	return s.productRepo.Delete(id)
}

// IncrClick 点击数 +1。
func (s *ProductService) IncrClick(id uint) error {
	return s.productRepo.IncrClick(id)
}

func toProduct(req *request.ProductReq) *model.Product {
	return &model.Product{
		Model:       gorm.Model{ID: req.ID},
		Name:        req.Name,
		Slug:        req.Slug,
		Logo:        req.Logo,
		Tagline:     req.Tagline,
		Description: req.Description,
		WebsiteURL:  req.WebsiteURL,
		CategoryID:  req.CategoryID,
		Pricing:     req.Pricing,
		Platforms:   req.Platforms,
		Status:      req.Status,
		Featured:    req.Featured,
	}
}
