package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"vbCodeWeb/model"
	"vbCodeWeb/repo"
)

// maxUploadBytes 单张图片大小上限。
const maxUploadBytes = 4 << 20 // 4MB

// allowedImageMime 允许的图片类型。
// 不接受 SVG：SVG 可携带脚本，直接访问时有 XSS 风险。
var allowedImageMime = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
	"image/gif":  true,
}

// UploadService 上传图片业务逻辑。
type UploadService struct {
	uploadRepo repo.UploadRepo
}

func NewUploadService(uploadRepo repo.UploadRepo) *UploadService {
	return &UploadService{uploadRepo: uploadRepo}
}

// Save 校验并保存一张图片，返回可访问的相对地址；文件为空时返回空串。
func (s *UploadService) Save(fh *multipart.FileHeader) (string, error) {
	if fh == nil || fh.Size == 0 {
		return "", nil
	}
	if fh.Size > maxUploadBytes {
		return "", fmt.Errorf("图片过大，单张上限 %dMB", maxUploadBytes>>20)
	}
	f, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf, err := io.ReadAll(io.LimitReader(f, maxUploadBytes+1))
	if err != nil {
		return "", err
	}
	if len(buf) > maxUploadBytes {
		return "", fmt.Errorf("图片过大，单张上限 %dMB", maxUploadBytes>>20)
	}
	mime := http.DetectContentType(buf)
	if !allowedImageMime[mime] {
		return "", errors.New("只支持 jpg / png / webp / gif 图片")
	}

	up := &model.Upload{Filename: fh.Filename, MimeType: mime, Size: int64(len(buf)), Data: buf}
	if err := s.uploadRepo.Create(up); err != nil {
		return "", err
	}
	return fmt.Sprintf("/api/uploads/%d", up.ID), nil
}

// GetByID 读取上传记录。
func (s *UploadService) GetByID(id uint) (*model.Upload, error) {
	return s.uploadRepo.GetByID(id)
}
