package bean

import (
	"encoding/json"
	"errors"
	"strings"
)

// Context 中间件的上下文对象, 包含了请求和响应对象
type Context struct {
	Request  Request
	Response Response
}

// SID 获取请求头的中的SID
func (ctx *Context) SID() (string, error) {
	authorizationHeaders, ok := ctx.Request.Header["Authorization"]
	if !ok || len(authorizationHeaders) == 0 {
		return "", errors.New("请求头中缺少Authorization字段")
	}

	if strings.Compare("", authorizationHeaders[0]) == 0 {
		return "", errors.New("请求头中有Authorization字段, 但为空")
	}

	return authorizationHeaders[0], nil
}

// JSON 返回JSON格式的响应数据
func (ctx *Context) JSON(m interface{}) error {
	w := ctx.Response

	headers := w.Header()
	headers.Set("Content-Type", "application/json")

	encoder := json.NewEncoder(ctx.Response.ResponseWriter)
	err := encoder.Encode(m)

	if err != nil {
		ctx.Error(err)
		return err
	}

	return nil
}

type format struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// FormatJSON 格式化输出
func (ctx *Context) FormatJSON(code int, message string, data interface{}) error {
	ret := format{code, message, data}

	return ctx.JSON(ret)
}

// Success 返回成功
func (ctx *Context) Success(data interface{}) error {
	return ctx.FormatJSON(0, "SUCCESS", data)
}

// Fail 返回失败
func (ctx *Context) Fail(message string) error {
	return ctx.FormatJSON(1, message, nil)
}

// String 返回String格式的响应数据
func (ctx *Context) String(content string) {
	w := ctx.Response

	headers := w.Header()
	headers.Set("Content-Type", "text/plain")

	ctx.Response.Write([]byte(content))
}

// Error 返回一个错误
func (ctx *Context) Error(err error) {
	w := ctx.Response

	w.WriteHeader(500)
	ctx.Response.Write([]byte(err.Error()))
}
