package service

import (
	"errors"

	"vbCodeWeb/model"
	"vbCodeWeb/repo"
	"vbCodeWeb/request"

	"gorm.io/gorm"
)

// MiniProgramService 小程序案例业务逻辑。
type MiniProgramService struct {
	miniProgramRepo repo.MiniProgramRepo
}

func NewMiniProgramService(miniProgramRepo repo.MiniProgramRepo) *MiniProgramService {
	return &MiniProgramService{miniProgramRepo: miniProgramRepo}
}

// List 分页查询小程序案例。
func (s *MiniProgramService) List(req *request.PageReq) ([]model.MiniProgram, int64, error) {
	req.Normalize()
	return s.miniProgramRepo.List(req)
}

// GetByID 查询小程序案例详情。
func (s *MiniProgramService) GetByID(id uint) (*model.MiniProgram, error) {
	return s.miniProgramRepo.GetByID(id)
}

// GetBySlug 按唯一标识查询小程序案例。
func (s *MiniProgramService) GetBySlug(slug string) (*model.MiniProgram, error) {
	return s.miniProgramRepo.GetBySlug(slug)
}

// Create 新建小程序案例。
func (s *MiniProgramService) Create(req *request.MiniProgramReq) (*model.MiniProgram, error) {
	m := toMiniProgram(req)
	if m.Status == 0 {
		m.Status = 1
	}
	if err := s.miniProgramRepo.Create(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Update 更新小程序案例。
func (s *MiniProgramService) Update(req *request.MiniProgramReq) (*model.MiniProgram, error) {
	if req.ID == 0 {
		return nil, errors.New("id 不能为空")
	}
	if err := s.miniProgramRepo.Update(toMiniProgram(req)); err != nil {
		return nil, err
	}
	return s.miniProgramRepo.GetByID(req.ID)
}

// Delete 删除小程序案例。
func (s *MiniProgramService) Delete(id uint) error {
	return s.miniProgramRepo.Delete(id)
}

// IncrClick 查看次数 +1。
func (s *MiniProgramService) IncrClick(id uint) error {
	return s.miniProgramRepo.IncrClick(id)
}

func toMiniProgram(req *request.MiniProgramReq) *model.MiniProgram {
	return &model.MiniProgram{
		Model:       gorm.Model{ID: req.ID},
		Name:        req.Name,
		Slug:        req.Slug,
		AppID:       req.AppID,
		Subject:     req.Subject,
		Cover:       req.Cover,
		MiniCodeImg: req.MiniCodeImg,
		Description: req.Description,
		BuilderTool: req.BuilderTool,
		CategoryID:  req.CategoryID,
		GuideURL:    req.GuideURL,
		Status:      req.Status,
		Featured:    req.Featured,
	}
}
