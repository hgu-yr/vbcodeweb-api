package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 业务响应码。
const (
	CodeSuccess = 0    // 成功
	CodeFail    = 1    // 通用失败
	CodeInvalid = 4000 // 参数错误
	CodeUnauth  = 4001 // 未认证
	CodeServer  = 5000 // 服务端错误
)

// Response 是统一的接口返回结构。
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// PageData 是分页返回结构。
type PageData struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Size  int   `json:"size"`
}

// Success 返回成功响应，data 可为 nil。
func Success(data any) *Response {
	return &Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	}
}

// Fail 返回失败响应，使用默认失败码。
func Fail(message string) *Response {
	return &Response{
		Code:    CodeFail,
		Message: message,
	}
}

// FailWith 返回指定错误码的失败响应。
func FailWith(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

// HTTPStatus 返回响应码对应的 HTTP 状态码。
func (r *Response) HTTPStatus() int {
	switch r.Code {
	case CodeSuccess:
		return http.StatusOK
	case CodeInvalid:
		return http.StatusBadRequest
	case CodeUnauth:
		return http.StatusUnauthorized
	case CodeServer:
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}

// JSON 按对应 HTTP 状态码将响应写入 gin 上下文。
func (r *Response) JSON(c *gin.Context) {
	c.JSON(r.HTTPStatus(), r)
}

// OK 以成功响应写入 gin 上下文。
func OK(c *gin.Context, data any) {
	Success(data).JSON(c)
}

// OKPage 以分页成功响应写入 gin 上下文。
func OKPage(c *gin.Context, list any, total int64, page, size int) {
	Success(PageData{List: list, Total: total, Page: page, Size: size}).JSON(c)
}

// Err 以失败响应写入 gin 上下文。
func Err(c *gin.Context, message string) {
	Fail(message).JSON(c)
}
