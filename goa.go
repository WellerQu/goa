package goa

import (
	"net/http"
)

// App Web应用程序
type App interface {
	Use(middleware MiddlewareHandler)
	ListenAndServe(addr string) error
}

// MiddlewareHandler 中间件处理程序
type MiddlewareHandler func(c *Context, next func())

type bean struct {
	// 中间件函数链
	middlewares MiddlewareHandler
}

// NewApp 创建应用程序
func NewApp() App {
	defaultNext := func(ctx *Context, next func()) {
		next()
	}
	b := bean{defaultNext}
	return &b
}

// ListenAndServe ...
func (b *bean) ListenAndServe(addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req := Request{r, ""}
		res := Response{w, 0, 0}
		ctx := Context{req, res}

		b.middlewares(&ctx, func() {
			w.Write([]byte{})
		})
	})

	return http.ListenAndServe(addr, nil)
}

// Use 添加中间件
func (b *bean) Use(handler MiddlewareHandler) {
	m := b.middlewares
	b.middlewares = func(ctx *Context, next func()) {
		m(ctx, func() {
			handler(ctx, next)
		})
	}
}
