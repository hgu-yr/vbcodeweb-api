package handler

import (
	"github.com/gin-gonic/gin"

	"vbCodeWeb/common"
	"vbCodeWeb/request"
	"vbCodeWeb/service"
)

// ProductController vibe coding 工具接口。
type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

// List 产品列表
func (pc *ProductController) List(c *gin.Context) {
	var req request.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	list, total, err := pc.productService.List(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OKPage(c, list, total, req.Page, req.Size)
}

// Detail 产品详情(按 id)
func (pc *ProductController) Detail(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	data, err := pc.productService.GetByID(id)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// DetailBySlug 产品详情(按 slug)
func (pc *ProductController) DetailBySlug(c *gin.Context) {
	data, err := pc.productService.GetBySlug(c.Param("slug"))
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Create 新建产品
func (pc *ProductController) Create(c *gin.Context) {
	var req request.ProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := pc.productService.Create(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Update 更新产品
func (pc *ProductController) Update(c *gin.Context) {
	var req request.ProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := pc.productService.Update(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Delete 删除产品
func (pc *ProductController) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := pc.productService.Delete(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "删除成功")
}

// Click 记录一次点击(用于跳转统计)
func (pc *ProductController) Click(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := pc.productService.IncrClick(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "ok")
}
