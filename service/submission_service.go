package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"
	"time"

	"vbCodeWeb/model"
	"vbCodeWeb/repo"
	"vbCodeWeb/request"
)

// SubmissionService 用户投稿业务逻辑。
type SubmissionService struct {
	submissionRepo repo.SubmissionRepo
	productRepo    repo.ProductRepo
	miniRepo       repo.MiniProgramRepo
	uploadSvc      UploadService
}

func NewSubmissionService(
	submissionRepo repo.SubmissionRepo,
	productRepo repo.ProductRepo,
	miniRepo repo.MiniProgramRepo,
	uploadSvc UploadService,
) *SubmissionService {
	return &SubmissionService{
		submissionRepo: submissionRepo,
		productRepo:    productRepo,
		miniRepo:       miniRepo,
		uploadSvc:      uploadSvc,
	}
}

// Submit 用户提交作品：保存图片后落库为「待审核」。
func (s *SubmissionService) Submit(
	req *request.SubmissionReq,
	image, code *multipart.FileHeader,
	ip string,
) (*model.Submission, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("名称不能为空")
	}
	kind := req.Kind
	if kind != model.SubmissionKindMini {
		kind = model.SubmissionKindTool
	}

	sub := &model.Submission{
		Kind:        kind,
		Name:        name,
		Slug:        normalizeSlug(req.Slug),
		Tagline:     strings.TrimSpace(req.Tagline),
		Description: strings.TrimSpace(req.Description),
		WebsiteURL:  strings.TrimSpace(req.WebsiteURL),
		Pricing:     strings.TrimSpace(req.Pricing),
		Platforms:   strings.TrimSpace(req.Platforms),
		AppID:       strings.TrimSpace(req.AppID),
		Subject:     strings.TrimSpace(req.Subject),
		BuilderTool: strings.TrimSpace(req.BuilderTool),
		GuideURL:    strings.TrimSpace(req.GuideURL),
		CategoryID:  req.CategoryID,
		Contact:     strings.TrimSpace(req.Contact),
		IP:          ip,
		Status:      model.SubmissionPending,
	}

	imgURL, err := s.uploadSvc.Save(image)
	if err != nil {
		return nil, err
	}
	if imgURL != "" {
		if kind == model.SubmissionKindMini {
			sub.CoverURL = imgURL
		} else {
			sub.LogoURL = imgURL
		}
	}

	codeURL, err := s.uploadSvc.Save(code)
	if err != nil {
		return nil, err
	}
	sub.MiniCodeURL = codeURL

	if err := s.submissionRepo.Create(sub); err != nil {
		return nil, err
	}
	return sub, nil
}

// List 审核列表。
func (s *SubmissionService) List(req *request.PageReq) ([]model.Submission, int64, error) {
	req.Normalize()
	return s.submissionRepo.List(req)
}

// CountPending 待审核数量。
func (s *SubmissionService) CountPending() (int64, error) {
	return s.submissionRepo.CountPending()
}

// GetByID 查询单条投稿。
func (s *SubmissionService) GetByID(id uint) (*model.Submission, error) {
	return s.submissionRepo.GetByID(id)
}

// Delete 删除投稿。
func (s *SubmissionService) Delete(id uint) error {
	return s.submissionRepo.Delete(id)
}

// Review 审核：通过则生成正式条目，拒绝则仅标记状态。
func (s *SubmissionService) Review(id uint, approve bool, req *request.SubmissionReviewReq) (*model.Submission, error) {
	sub, err := s.submissionRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("投稿不存在")
	}
	if sub.Status != model.SubmissionPending {
		return nil, errors.New("该投稿已审核过")
	}

	now := time.Now()
	sub.ReviewedAt = &now
	sub.ReviewNote = strings.TrimSpace(req.Note)

	if !approve {
		sub.Status = model.SubmissionRejected
		if err := s.submissionRepo.Save(sub); err != nil {
			return nil, err
		}
		return sub, nil
	}

	slug := normalizeSlug(req.Slug)
	if slug == "" {
		slug = sub.Slug
	}
	if slug == "" {
		slug = fmt.Sprintf("u%d", sub.ID)
	}
	slug = s.uniqueSlug(slug, sub.Kind)

	categoryID := req.CategoryID
	if categoryID == 0 {
		categoryID = sub.CategoryID
	}

	if sub.Kind == model.SubmissionKindMini {
		m := &model.MiniProgram{
			Name:        sub.Name,
			Slug:        slug,
			AppID:       sub.AppID,
			Subject:     sub.Subject,
			Cover:       sub.CoverURL,
			MiniCodeImg: sub.MiniCodeURL,
			Description: sub.Description,
			BuilderTool: sub.BuilderTool,
			CategoryID:  categoryID,
			GuideURL:    sub.GuideURL,
			Status:      1,
		}
		if err := s.miniRepo.Create(m); err != nil {
			return nil, err
		}
		sub.PublishedID = m.ID
	} else {
		p := &model.Product{
			Name:        sub.Name,
			Slug:        slug,
			Logo:        sub.LogoURL,
			Tagline:     sub.Tagline,
			Description: sub.Description,
			WebsiteURL:  sub.WebsiteURL,
			CategoryID:  categoryID,
			Pricing:     sub.Pricing,
			Platforms:   sub.Platforms,
			Status:      1,
		}
		if err := s.productRepo.Create(p); err != nil {
			return nil, err
		}
		sub.PublishedID = p.ID
	}

	sub.Status = model.SubmissionApproved
	if err := s.submissionRepo.Save(sub); err != nil {
		return nil, err
	}
	return sub, nil
}

var slugRe = regexp.MustCompile(`[^a-z0-9]+`)

// normalizeSlug 只保留小写字母/数字/连字符；中文等会被丢弃（可能返回空串）。
func normalizeSlug(raw string) string {
	v := slugRe.ReplaceAllString(strings.ToLower(strings.TrimSpace(raw)), "-")
	return strings.Trim(v, "-")
}

// uniqueSlug 保证 slug 不与其他条目冲突。
func (s *SubmissionService) uniqueSlug(base string, kind int8) string {
	cand := base
	for i := 1; i <= 50; i++ {
		var err error
		if kind == model.SubmissionKindMini {
			_, err = s.miniRepo.GetBySlug(cand)
		} else {
			_, err = s.productRepo.GetBySlug(cand)
		}
		if err != nil {
			return cand
		}
		cand = fmt.Sprintf("%s-%d", base, i)
	}
	return fmt.Sprintf("%s-%d", base, time.Now().Unix())
}
