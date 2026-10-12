package request

// ProductReq 创建/更新 vibe coding 工具请求。
type ProductReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Logo        string `json:"logo"`
	Tagline     string `json:"tagline"`
	Description string `json:"description"`
	WebsiteURL  string `json:"website_url"`
	CategoryID  uint   `json:"category_id"`
	Pricing     string `json:"pricing"`
	Platforms   string `json:"platforms"`
	Status      int8   `json:"status"`
	Featured    int8   `json:"featured"`
}
