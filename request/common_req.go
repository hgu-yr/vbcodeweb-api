package request

// PageReq 通用分页与筛选请求。
type PageReq struct {
	Page       int    `form:"page" json:"page"`
	Size       int    `form:"size" json:"size"`
	Keyword    string `form:"keyword" json:"keyword"`
	CategoryID uint   `form:"category_id" json:"category_id"`
	Status     *int8  `form:"status" json:"status"`
	Featured   *int8  `form:"featured" json:"featured"`
}

// Normalize 为分页参数补默认值并做上限保护。
func (p *PageReq) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Size < 1 {
		p.Size = 10
	}
	if p.Size > 100 {
		p.Size = 100
	}
}
