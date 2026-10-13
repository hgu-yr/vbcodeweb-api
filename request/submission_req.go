package request

// SubmissionReq 用户投稿请求（multipart/form-data）。
type SubmissionReq struct {
	Kind        int8   `form:"kind"`
	Name        string `form:"name"`
	Slug        string `form:"slug"`
	Tagline     string `form:"tagline"`
	Description string `form:"description"`

	WebsiteURL string `form:"website_url"`
	Pricing    string `form:"pricing"`
	Platforms  string `form:"platforms"`

	AppID       string `form:"app_id"`
	Subject     string `form:"subject"`
	BuilderTool string `form:"builder_tool"`
	GuideURL    string `form:"guide_url"`

	CategoryID uint   `form:"category_id"`
	Contact    string `form:"contact"`
}

// SubmissionReviewReq 审核请求，仅通过时可指定 slug / 分类。
type SubmissionReviewReq struct {
	Slug       string `json:"slug"`
	CategoryID uint   `json:"category_id"`
	Note       string `json:"note"`
}
