package seed

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"vbCodeWeb/model"
	"vbCodeWeb/utils"
)

// Run 灌入 mock 数据，按 slug 幂等：已存在的记录不会重复插入。
func Run(db *gorm.DB) {
	seedUsers(db)
	catIDs := seedCategories(db)
	seedProducts(db, catIDs)
	seedMiniPrograms(db, catIDs)
	seedArticles(db, catIDs)
	fmt.Println("mock 数据填充完成")
}

// ---------- 管理员 ----------

func seedUsers(db *gorm.DB) {
	var existing model.User
	if db.Where("username = ?", "admin").Limit(1).Find(&existing).RowsAffected == 0 {
		hash, err := utils.HashPassword("admin123")
		if err != nil {
			fmt.Printf("生成密码失败: %v\n", err)
			return
		}
		u := model.User{Username: "admin", Password: hash, Nickname: "超级管理员", Role: "admin", Status: 1}
		if err := db.Create(&u).Error; err != nil {
			fmt.Printf("创建管理员失败: %v\n", err)
			return
		}
		fmt.Println("默认管理员已创建：admin / admin123")
	}
}

// ---------- 分类 ----------

func seedCategories(db *gorm.DB) map[string]uint {
	categories := []model.Category{
		{Name: "AI 编程助手", Slug: "ai-assistant", Scope: 1, Sort: 100},
		{Name: "Web 应用生成", Slug: "web-builder", Scope: 1, Sort: 95},
		{Name: "全栈开发", Slug: "fullstack", Scope: 1, Sort: 90},
		{Name: "移动端开发", Slug: "mobile", Scope: 1, Sort: 85},
		{Name: "UI 设计生成", Slug: "ui-design", Scope: 1, Sort: 80},
		{Name: "代码测试", Slug: "testing", Scope: 1, Sort: 75},
		{Name: "工具类小程序", Slug: "mini-tool", Scope: 2, Sort: 100},
		{Name: "生活服务小程序", Slug: "mini-life", Scope: 2, Sort: 90},
		{Name: "内容娱乐小程序", Slug: "mini-fun", Scope: 2, Sort: 80},
		{Name: "行业观察", Slug: "industry", Scope: 3, Sort: 100},
		{Name: "实战教程", Slug: "tutorial", Scope: 3, Sort: 90},
	}

	ids := make(map[string]uint, len(categories))
	for i := range categories {
		c := categories[i]
		var existing model.Category
		if db.Where("slug = ?", c.Slug).Limit(1).Find(&existing).RowsAffected == 0 {
			if err := db.Create(&c).Error; err != nil {
				fmt.Printf("创建分类失败 %s: %v\n", c.Slug, err)
				continue
			}
			existing = c
		}
		ids[c.Slug] = existing.ID
	}
	fmt.Printf("分类就绪：%d 个\n", len(ids))
	return ids
}

// ---------- 工具/产品 ----------

func seedProducts(db *gorm.DB, catIDs map[string]uint) {
	products := []model.Product{
		{Name: "Cursor", Slug: "cursor", Tagline: "AI 优先的代码编辑器", Description: "基于 VS Code 打造，内置多模型对话、Tab 补全与全仓库 Agent，是当前最主流的 AI 编程编辑器之一。", WebsiteURL: "https://cursor.com", CategoryID: catIDs["ai-assistant"], Pricing: "免费增值", Platforms: "Windows,macOS,Linux", Featured: 1},
		{Name: "Windsurf", Slug: "windsurf", Tagline: "首个 Agentic IDE", Description: "以 Cascade 智能体为核心的 IDE，强调上下文感知与多文件级联修改，适合中大型项目重构。", WebsiteURL: "https://windsurf.com", CategoryID: catIDs["ai-assistant"], Pricing: "免费增值", Platforms: "Windows,macOS,Linux", Featured: 1},
		{Name: "Trae", Slug: "trae", Tagline: "字节跳动出品的 AI IDE", Description: "面向中文开发者的 AI 原生 IDE，支持 Builder 模式从需求直接生成完整项目。", WebsiteURL: "https://www.trae.ai", CategoryID: catIDs["ai-assistant"], Pricing: "免费增值", Platforms: "Windows,macOS,Linux", Featured: 1},
		{Name: "GitHub Copilot", Slug: "github-copilot", Tagline: "久经考验的代码补全与对话助手", Description: "GitHub 官方 AI 编程助手，深度集成 VS Code、JetBrains 等主流 IDE，支持 Agent 模式。", WebsiteURL: "https://github.com/features/copilot", CategoryID: catIDs["ai-assistant"], Pricing: "付费", Platforms: "Windows,macOS,Linux,Web"},
		{Name: "Claude Code", Slug: "claude-code", Tagline: "终端里的编程 Agent", Description: "Anthropic 出品的命令行编程智能体，可读写文件、执行命令并完成跨文件重构。", WebsiteURL: "https://claude.com/product/claude-code", CategoryID: catIDs["ai-assistant"], Pricing: "付费", Platforms: "macOS,Linux,Windows"},
		{Name: "Cline", Slug: "cline", Tagline: "开源 VS Code Agent 插件", Description: "可在编辑器内执行多步骤任务的开源 Agent 插件，支持自带模型 API Key。", WebsiteURL: "https://cline.bot", CategoryID: catIDs["ai-assistant"], Pricing: "免费", Platforms: "VS Code"},
		{Name: "Continue", Slug: "continue", Tagline: "可定制的开源编程助手", Description: "支持本地与云端模型自由切换的开源 AI 编程助手，适合对数据合规敏感的场景。", WebsiteURL: "https://continue.dev", CategoryID: catIDs["ai-assistant"], Pricing: "免费", Platforms: "VS Code,JetBrains"},
		{Name: "Roo Code", Slug: "roo-code", Tagline: "开源编辑器内智能体", Description: "由 Cline 衍生出的开源 Agent，支持多模式与自定义工具调用。", WebsiteURL: "https://roocode.com", CategoryID: catIDs["ai-assistant"], Pricing: "免费", Platforms: "VS Code"},

		{Name: "v0", Slug: "v0", Tagline: "Vercel 出品的 UI 生成器", Description: "用自然语言描述界面，直接生成可用的 React + Tailwind 组件代码。", WebsiteURL: "https://v0.dev", CategoryID: catIDs["ui-design"], Pricing: "免费增值", Platforms: "Web", Featured: 1},
		{Name: "Onlook", Slug: "onlook", Tagline: "开源的视觉化 React 编辑器", Description: "在浏览器中直接对 React 应用进行可视化编辑，改动实时写回源码。", WebsiteURL: "https://onlook.com", CategoryID: catIDs["ui-design"], Pricing: "免费", Platforms: "Web,macOS"},

		{Name: "Bolt.new", Slug: "bolt-new", Tagline: "浏览器里的全栈生成器", Description: "StackBlitz 出品，在浏览器内直接生成、运行和部署全栈 Web 应用。", WebsiteURL: "https://bolt.new", CategoryID: catIDs["fullstack"], Pricing: "免费增值", Platforms: "Web", Featured: 1},
		{Name: "Lovable", Slug: "lovable", Tagline: "对话式全栈应用生成", Description: "通过自然语言对话从零构建带数据库与后端的完整应用，适合快速验证产品想法。", WebsiteURL: "https://lovable.dev", CategoryID: catIDs["fullstack"], Pricing: "免费增值", Platforms: "Web", Featured: 1},
		{Name: "Replit Agent", Slug: "replit-agent", Tagline: "云端一键搭建应用", Description: "Replit 的云端智能体，从描述到部署全程托管，无需本地环境。", WebsiteURL: "https://replit.com", CategoryID: catIDs["fullstack"], Pricing: "免费增值", Platforms: "Web"},
		{Name: "Devin", Slug: "devin", Tagline: "自主软件工程师", Description: "Cognition 推出的自主 AI 工程师，可独立完成端到端开发任务。", WebsiteURL: "https://devin.ai", CategoryID: catIDs["fullstack"], Pricing: "付费", Platforms: "Web"},
		{Name: "Firebase Studio", Slug: "firebase-studio", Tagline: "Google 的云端全栈工作台", Description: "基于云开发环境的全栈 AI 构建平台，深度集成 Firebase 生态。", WebsiteURL: "https://firebase.google.com/studio", CategoryID: catIDs["fullstack"], Pricing: "免费增值", Platforms: "Web"},

		{Name: "Qodo", Slug: "qodo", Tagline: "AI 测试与代码质量", Description: "聚焦测试生成、代码审查与覆盖率提升的 AI 工具，前身为 Codium。", WebsiteURL: "https://www.qodo.ai", CategoryID: catIDs["testing"], Pricing: "免费增值", Platforms: "VS Code,JetBrains"},

		{Name: "Augment Code", Slug: "augment-code", Tagline: "面向大型代码库的 AI 助手", Description: "主打超长上下文与团队级代码理解，适合大型遗留项目。", WebsiteURL: "https://www.augmentcode.com", CategoryID: catIDs["ai-assistant"], Pricing: "付费", Platforms: "VS Code,JetBrains"},
		{Name: "Tabnine", Slug: "tabnine", Tagline: "注重隐私的 AI 补全", Description: "强调代码隐私与私有化部署的 AI 代码补全工具，适合企业团队。", WebsiteURL: "https://www.tabnine.com", CategoryID: catIDs["ai-assistant"], Pricing: "免费增值", Platforms: "VS Code,JetBrains,Web"},
		{Name: "Supermaven", Slug: "supermaven", Tagline: "极速长上下文补全", Description: "以低延迟补全为卖点，支持超长上下文理解。", WebsiteURL: "https://supermaven.com", CategoryID: catIDs["ai-assistant"], Pricing: "免费增值", Platforms: "VS Code,JetBrains"},
	}

	for i := range products {
		p := products[i]
		p.Status = 1
		p.Logo = placeholder("120x120", p.Name)
		var existing model.Product
		if db.Where("slug = ?", p.Slug).Limit(1).Find(&existing).RowsAffected == 0 {
			if err := db.Create(&p).Error; err != nil {
				fmt.Printf("创建产品失败 %s: %v\n", p.Slug, err)
			}
		}
	}
	fmt.Println("工具/产品数据就绪")
}

// ---------- 小程序案例 ----------

func seedMiniPrograms(db *gorm.DB, catIDs map[string]uint) {
	programs := []model.MiniProgram{
		{Name: "每日一句", Slug: "meiri-yiju", Subject: "独立开发者 @阿岳", Description: "每天推送一句治愈文案，支持收藏与分享。用 Cursor 从零到上线仅花一个下午。", BuilderTool: "Cursor", CategoryID: catIDs["mini-tool"], Featured: 1},
		{Name: "番茄作家", Slug: "fanqie-zuojia", Subject: "个人开发者", Description: "番茄工作法 + 写作统计，帮助专注写作的工具型小程序，由 Trae 生成主体代码。", BuilderTool: "Trae", CategoryID: catIDs["mini-tool"], Featured: 1},
		{Name: "极简记账", Slug: "jijian-jizhang", Subject: "学生团队", Description: "三秒记一笔，支持分类统计与月度报表，用 Windsurf 完成多文件重构。", BuilderTool: "Windsurf", CategoryID: catIDs["mini-tool"]},
		{Name: "待办卡片", Slug: "daiban-kapian", Subject: "独立开发者", Description: "卡片式待办清单，支持拖拽排序与提醒。全程由 Cursor Agent 生成。", BuilderTool: "Cursor", CategoryID: catIDs["mini-tool"]},
		{Name: "单位换算器", Slug: "danwei-huansuan", Subject: "工具控", Description: "覆盖长度、重量、面积、货币等常用换算，界面由 v0 生成后再接入逻辑。", BuilderTool: "v0", CategoryID: catIDs["mini-tool"]},

		{Name: "快递追踪", Slug: "kuaidi-zhuizong", Subject: "独立开发者", Description: "聚合多家快递单号查询，自动识别快递公司。用 Claude Code 编写后端接口联调。", BuilderTool: "Claude Code", CategoryID: catIDs["mini-life"], Featured: 1},
		{Name: "天气小报", Slug: "tianqi-xiaobao", Subject: "独立开发者", Description: "极简天气 + 穿衣建议 + 空气质量，界面清爽，由 Cursor 生成。", BuilderTool: "Cursor", CategoryID: catIDs["mini-life"]},
		{Name: "BMI 健康管家", Slug: "bmi-guanjia", Subject: "健康达人", Description: "BMI 计算、目标体重与健康建议，配合打卡记录。UI 由 v0 生成。", BuilderTool: "v0", CategoryID: catIDs["mini-life"]},
		{Name: "今天吃什么", Slug: "jintian-chi-shenme", Subject: "个人开发者", Description: "选择困难症救星，一键随机推荐午餐晚餐，由 Replit Agent 快速搭建。", BuilderTool: "Replit Agent", CategoryID: catIDs["mini-life"], Featured: 1},

		{Name: "单词打卡机", Slug: "danci-dakaji", Subject: "英语学习小组", Description: "每日单词打卡 + 连续天数统计 + 词库自建，由 Bolt.new 生成原型后迁移到小程序。", BuilderTool: "Bolt.new", CategoryID: catIDs["mini-fun"]},
		{Name: "心情晴雨表", Slug: "xinqing-qingyubiao", Subject: "独立开发者", Description: "记录每日心情并生成情绪曲线，界面由 Lovable 对话式生成。", BuilderTool: "Lovable", CategoryID: catIDs["mini-fun"]},
		{Name: "星座物语", Slug: "xingzuo-wuyu", Subject: "内容创作者", Description: "每日星座运势与配对测试，配图精美，主体由 Trae 生成。", BuilderTool: "Trae", CategoryID: catIDs["mini-fun"]},
	}

	for i := range programs {
		m := programs[i]
		m.Status = 1
		m.Cover = placeholder("200x120", m.Name)
		m.MiniCodeImg = placeholder("200x200", "MiniCode")
		var existing model.MiniProgram
		if db.Where("slug = ?", m.Slug).Limit(1).Find(&existing).RowsAffected == 0 {
			if err := db.Create(&m).Error; err != nil {
				fmt.Printf("创建小程序案例失败 %s: %v\n", m.Slug, err)
			}
		}
	}
	fmt.Println("小程序案例数据就绪")
}

// ---------- 文章 ----------

func seedArticles(db *gorm.DB, catIDs map[string]uint) {
	now := time.Now()
	articles := []model.Article{
		{Title: "Vibe Coding 入门指南：用自然语言写代码", Slug: "vibe-coding-guide", Summary: "什么是 Vibe Coding？它和传统编程有何不同？本文带你从零理解这一新范式。", CategoryID: catIDs["industry"], Status: 1, PublishedAt: &now,
			Content: "# Vibe Coding 入门指南\n\nVibe Coding 指用自然语言描述需求、由 AI 生成并迭代代码的开发方式。\n\n## 核心要点\n\n1. 描述意图，而非实现\n2. 快速迭代，小步验证\n3. 人负责判断，AI 负责产出\n\n> 建议从一个小工具开始练习。"},
		{Title: "2026 年最值得尝试的 10 款 AI 编程工具", Slug: "top-10-ai-coding-tools", Summary: "从 Cursor 到 Windsurf、Trae、Bolt.new，盘点当下最值得上手的 AI 编程产品。", CategoryID: catIDs["industry"], Status: 1, PublishedAt: &now,
			Content: "# 2026 年 AI 编程工具盘点\n\n- Cursor：AI 优先编辑器\n- Windsurf：Agentic IDE\n- Trae：中文友好\n- Bolt.new：浏览器全栈\n- Lovable：对话建应用\n\n每款工具的定位与适用场景各不相同。"},
		{Title: "我用 30 分钟做了一个微信小程序", Slug: "build-miniprogram-in-30min", Summary: "记录一次完整的 Vibe Coding 小程序开发过程，从想法到可运行 Demo。", CategoryID: catIDs["tutorial"], Status: 1, PublishedAt: &now,
			Content: "# 30 分钟做一个小程序\n\n## 00:00 需求\n做一个每日一句小程序。\n\n## 00:05 生成\n用 Cursor 描述需求，生成页面结构。\n\n## 00:20 调试\n修正样式与接口。\n\n## 00:30 完成\n提交审核。"},
		{Title: "Cursor vs Windsurf vs Trae 深度横评", Slug: "cursor-windsurf-trae-compare", Summary: "三款主流 AI IDE 在补全、Agent 能力、中文支持与价格上的全面对比。", CategoryID: catIDs["industry"], Status: 1, PublishedAt: &now,
			Content: "# 三款 AI IDE 横评\n\n| 工具 | 补全 | Agent | 中文 |\n| --- | --- | --- | --- |\n| Cursor | 强 | 强 | 中 |\n| Windsurf | 强 | 很强 | 中 |\n| Trae | 中 | 强 | 强 |\n\n选择取决于团队习惯与合规要求。"},
		{Title: "Vibe Coding 的边界与陷阱", Slug: "vibe-coding-pitfalls", Summary: "AI 生成代码并非万能，聊聊哪些场景不适合，以及如何控制质量风险。", CategoryID: catIDs["industry"], Status: 0, PublishedAt: nil,
			Content: "# Vibe Coding 的边界\n\n- 安全敏感场景需人工审计\n- 复杂架构仍需人类设计\n- 生成代码要纳入测试流程\n\n草稿占位。"},
	}

	for i := range articles {
		a := articles[i]
		var existing model.Article
		if db.Where("slug = ?", a.Slug).Limit(1).Find(&existing).RowsAffected == 0 {
			if err := db.Create(&a).Error; err != nil {
				fmt.Printf("创建文章失败 %s: %v\n", a.Slug, err)
			}
		}
	}
	fmt.Println("文章数据就绪")
}

// placeholder 生成一个占位图地址，方便前端直接展示。
func placeholder(size, text string) string {
	return fmt.Sprintf("https://placehold.co/%s/1e293b/ffffff?text=%s", size, text)
}
