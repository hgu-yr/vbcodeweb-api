package request

// MiniProgramReq 创建/更新小程序案例请求。
type MiniProgramReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	AppID       string `json:"app_id"`
	Subject     string `json:"subject"`
	Cover       string `json:"cover"`
	MiniCodeImg string `json:"mini_code_img"`
	Description string `json:"description"`
	BuilderTool string `json:"builder_tool"`
	CategoryID  uint   `json:"category_id"`
	GuideURL    string `json:"guide_url"`
	Status      int8   `json:"status"`
	Featured    int8   `json:"featured"`
}
