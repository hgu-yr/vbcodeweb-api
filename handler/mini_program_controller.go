package handler

import (
	"github.com/gin-gonic/gin"

	"vbCodeWeb/common"
	"vbCodeWeb/request"
	"vbCodeWeb/service"
)

// MiniProgramController 小程序案例接口。
type MiniProgramController struct {
	miniProgramService service.MiniProgramService
}

func NewMiniProgramController(miniProgramService service.MiniProgramService) *MiniProgramController {
	return &MiniProgramController{miniProgramService: miniProgramService}
}

// List 小程序案例列表
func (mc *MiniProgramController) List(c *gin.Context) {
	var req request.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	list, total, err := mc.miniProgramService.List(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OKPage(c, list, total, req.Page, req.Size)
}

// Detail 小程序案例详情(按 id)
func (mc *MiniProgramController) Detail(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	data, err := mc.miniProgramService.GetByID(id)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// DetailBySlug 小程序案例详情(按 slug)
func (mc *MiniProgramController) DetailBySlug(c *gin.Context) {
	data, err := mc.miniProgramService.GetBySlug(c.Param("slug"))
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Create 新建小程序案例
func (mc *MiniProgramController) Create(c *gin.Context) {
	var req request.MiniProgramReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := mc.miniProgramService.Create(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Update 更新小程序案例
func (mc *MiniProgramController) Update(c *gin.Context) {
	var req request.MiniProgramReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := mc.miniProgramService.Update(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Delete 删除小程序案例
func (mc *MiniProgramController) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := mc.miniProgramService.Delete(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "删除成功")
}

// Click 记录一次查看
func (mc *MiniProgramController) Click(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := mc.miniProgramService.IncrClick(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "ok")
}
