package handler

import (
	"github.com/gin-gonic/gin"

	"vbCodeWeb/common"
	"vbCodeWeb/request"
	"vbCodeWeb/service"
)

// ArticleController 文章接口。
type ArticleController struct {
	articleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) *ArticleController {
	return &ArticleController{articleService: articleService}
}

// List 文章列表
func (ac *ArticleController) List(c *gin.Context) {
	var req request.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	list, total, err := ac.articleService.List(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OKPage(c, list, total, req.Page, req.Size)
}

// Detail 文章详情(按 id)
func (ac *ArticleController) Detail(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	data, err := ac.articleService.GetByID(id)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// DetailBySlug 文章详情(按 slug)
func (ac *ArticleController) DetailBySlug(c *gin.Context) {
	data, err := ac.articleService.GetBySlug(c.Param("slug"))
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Create 新建文章
func (ac *ArticleController) Create(c *gin.Context) {
	var req request.ArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := ac.articleService.Create(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Update 更新文章
func (ac *ArticleController) Update(c *gin.Context) {
	var req request.ArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	data, err := ac.articleService.Update(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, data)
}

// Delete 删除文章
func (ac *ArticleController) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := ac.articleService.Delete(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "删除成功")
}

// View 记录一次阅读
func (ac *ArticleController) View(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		common.Err(c, "id 参数错误")
		return
	}
	if err := ac.articleService.IncrView(id); err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, "ok")
}
