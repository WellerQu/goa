package bean

import "net/http"

// Response 响应对象
type Response struct {
	http.ResponseWriter
	status int
	length int
}

// Header 返回响应头对象, 实现http.ResponseWriter
func (res *Response) Header() http.Header {
	return res.ResponseWriter.Header()
}

// WriteHeader 实现http.ResponseWriter
func (res *Response) WriteHeader(status int) {
	if res.status != 0 {
		return
	}

	res.ResponseWriter.WriteHeader(status)
	res.status = status
}

// Write 将数据写入响应流中, 实现http.ResponseWriter
func (res *Response) Write(b []byte) (int, error) {
	n, err := res.ResponseWriter.Write(b)
	res.length += n

	return n, err
}

// Status 获取设置的状态码
func (res *Response) Status() int {
	return res.status
}

// ContentLength 获取响应流内容长度
func (res *Response) ContentLength() int {
	return res.length
}
