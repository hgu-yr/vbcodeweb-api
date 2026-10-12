package main

import (
	"os"

	"vbCodeWeb/handler"
	"vbCodeWeb/repo"
	"vbCodeWeb/router"
	"vbCodeWeb/seed"
	"vbCodeWeb/service"
	"vbCodeWeb/utils"
)

func main() {
	db := utils.InitDB()

	// 填充 mock 数据：go run . seed
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seed.Run(db)
		return
	}

	// 依赖注入
	categoryRepo := repo.NewCategoryRepo(db)
	categorySvc := service.NewCategoryService(*categoryRepo)
	categoryCtrl := handler.NewCategoryController(*categorySvc)

	productRepo := repo.NewProductRepo(db)
	productSvc := service.NewProductService(*productRepo)
	productCtrl := handler.NewProductController(*productSvc)

	miniProgramRepo := repo.NewMiniProgramRepo(db)
	miniProgramSvc := service.NewMiniProgramService(*miniProgramRepo)
	miniProgramCtrl := handler.NewMiniProgramController(*miniProgramSvc)

	articleRepo := repo.NewArticleRepo(db)
	articleSvc := service.NewArticleService(*articleRepo)
	articleCtrl := handler.NewArticleController(*articleSvc)

	userRepo := repo.NewUserRepo(db)
	authSvc := service.NewAuthService(*userRepo)
	authCtrl := handler.NewAuthController(*authSvc)

	// 云平台（如 Railway）通过 PORT 环境变量指定监听端口，默认 8899
	port := os.Getenv("PORT")
	if port == "" {
		port = "8899"
	}

	r := router.InitRouter(categoryCtrl, productCtrl, miniProgramCtrl, articleCtrl, authCtrl)
	r.Run(":" + port)
}
