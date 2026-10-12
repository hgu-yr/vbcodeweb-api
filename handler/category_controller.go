package handler

import (
	"github.com/gin-gonic/gin"

	"vbCodeWeb/common"
	"vbCodeWeb/request"
	"vbCodeWeb/service"
)

// CategoryController 分类接口。
type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

// List 分类列表
func (cc *CategoryController) List(c *gin.Context) {
	var req request.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	list, total, err := cc.categoryService.List(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OKPage(c, list, total, req.Page, req.Size)
}

// Detail 分类详情
func (cc *CategoryController) Detail(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	data, err := cc.categoryService.GetByID(id)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Create 新建分类
func (cc *CategoryController) Create(c *gin.Context) {
	var req request.CategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := cc.categoryService.Create(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Update 更新分类
func (cc *CategoryController) Update(c *gin.Context) {
	var req request.CategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := cc.categoryService.Update(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Delete 删除分类
func (cc *CategoryController) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := cc.categoryService.Delete(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "删除成功")
}
