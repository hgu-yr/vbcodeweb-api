package handler

import (
	"github.com/gin-gonic/gin"

	"vbCodeWeb/common"
	"vbCodeWeb/request"
	"vbCodeWeb/service"
)

// SubmissionController 用户投稿接口。
type SubmissionController struct {
	submissionService service.SubmissionService
}

func NewSubmissionController(submissionService service.SubmissionService) *SubmissionController {
	return &SubmissionController{submissionService: submissionService}
}

// Create 公开投稿（multipart/form-data，图片字段为 image / code）
func (sc *SubmissionController) Create(c *gin.Context) {
	var req request.SubmissionReq
	if err := c.ShouldBind(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	image, _ := c.FormFile("image")
	code, _ := c.FormFile("code")

	data, err := sc.submissionService.Submit(&req, image, code, c.ClientIP())
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, gin.H{"id": data.ID, "status": data.Status})
}

// List 投稿列表（需登录）
func (sc *SubmissionController) List(c *gin.Context) {
	var req request.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	list, total, err := sc.submissionService.List(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OKPage(c, list, total, req.Page, req.Size)
}

// Pending 待审核数量（需登录）
func (sc *SubmissionController) Pending(c *gin.Context) {
	n, err := sc.submissionService.CountPending()
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, gin.H{"count": n})
}

// Approve 审核通过
func (sc *SubmissionController) Approve(c *gin.Context) {
	sc.review(c, true)
}

// Reject 审核拒绝
func (sc *SubmissionController) Reject(c *gin.Context) {
	sc.review(c, false)
}

func (sc *SubmissionController) review(c *gin.Context, approve bool) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	var req request.SubmissionReviewReq
	// body 可选：仅传 note 或什么都不传也允许
	_ = c.ShouldBindJSON(&req)

	data, err := sc.submissionService.Review(id, approve, &req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Delete 删除投稿（需登录）
func (sc *SubmissionController) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := sc.submissionService.Delete(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "删除成功")
}
