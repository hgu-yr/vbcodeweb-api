package request

import "time"

// ArticleReq 创建/更新文章请求。
type ArticleReq struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title" binding:"required"`
	Slug        string     `json:"slug" binding:"required"`
	Summary     string     `json:"summary"`
	Content     string     `json:"content"`
	Cover       string     `json:"cover"`
	CategoryID  uint       `json:"category_id"`
	Status      int8       `json:"status"`
	PublishedAt *time.Time `json:"published_at"`
}
