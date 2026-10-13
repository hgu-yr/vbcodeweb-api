package handler

import (
	"github.com/gin-gonic/gin"

	"vbCodeWeb/service"
)

// UploadController 上传图片读取接口。
type UploadController struct {
	uploadService service.UploadService
}

func NewUploadController(uploadService service.UploadService) *UploadController {
	return &UploadController{uploadService: uploadService}
}

// Get 按 id 输出图片二进制。
func (uc *UploadController) Get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		c.Status(404)
		return
	}
	u, err := uc.uploadService.GetByID(id)
	if err != nil {
		c.Status(404)
		return
	}
	c.Header("Cache-Control", "public, max-age=31536000, immutable")
	c.Data(200, u.MimeType, u.Data)
}
