package router

import (
	"vbCodeWeb/handler"
	"vbCodeWeb/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 注册所有 HTTP 路由。
func InitRouter(
	category *handler.CategoryController,
	product *handler.ProductController,
	miniProgram *handler.MiniProgramController,
	article *handler.ArticleController,
	auth *handler.AuthController,
) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		// 鉴权
		api.POST("/auth/login", auth.Login)
		api.GET("/auth/profile", middleware.JWTAuth(), auth.Profile)

		// 分类（读）
		api.GET("/categories", category.List)
		api.GET("/categories/:id", category.Detail)

		// 工具/产品（读 + 点击统计）
		api.GET("/products", product.List)
		api.GET("/products/slug/:slug", product.DetailBySlug)
		api.GET("/products/:id", product.Detail)
		api.POST("/products/:id/click", product.Click)

		// 小程序案例（读 + 查看统计）
		api.GET("/mini-programs", miniProgram.List)
		api.GET("/mini-programs/slug/:slug", miniProgram.DetailBySlug)
		api.GET("/mini-programs/:id", miniProgram.Detail)
		api.POST("/mini-programs/:id/click", miniProgram.Click)

		// 文章（读 + 阅读统计）
		api.GET("/articles", article.List)
		api.GET("/articles/slug/:slug", article.DetailBySlug)
		api.GET("/articles/:id", article.Detail)
		api.POST("/articles/:id/view", article.View)

		// 以下写操作需要登录
		admin := api.Group("", middleware.JWTAuth())
		{
			admin.POST("/categories", category.Create)
			admin.PUT("/categories", category.Update)
			admin.DELETE("/categories/:id", category.Delete)

			admin.POST("/products", product.Create)
			admin.PUT("/products", product.Update)
			admin.DELETE("/products/:id", product.Delete)

			admin.POST("/mini-programs", miniProgram.Create)
			admin.PUT("/mini-programs", miniProgram.Update)
			admin.DELETE("/mini-programs/:id", miniProgram.Delete)

			admin.POST("/articles", article.Create)
			admin.PUT("/articles", article.Update)
			admin.DELETE("/articles/:id", article.Delete)
		}
	}
	return r
}
