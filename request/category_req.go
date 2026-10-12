package request

// CategoryReq 创建/更新分类请求。
type CategoryReq struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required"`
	Slug  string `json:"slug" binding:"required"`
	Scope int8   `json:"scope"`
	Sort  int    `json:"sort"`
}
