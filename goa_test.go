package goa

import (
	"testing"
)

func TestCore(t *testing.T) {
	// Example 1: 使用中间件
	// curl -i "http://localhost:8082/anything"
	/*
		beanApp := NewApp()
		beanApp.Use(func(ctx *Context, next func()) {
			ctx.Response.WriteHeader(204)
			log.Println("1 Level")
			next()
		})
		beanApp.Use(func(ctx *Context, next func()) {
			log.Println("2 Level")
			next()
		})
		beanApp.ListenAndServe(":8082")
		//*/

	// Example 2: 使用路由
	// curl -i "http://localhost:8082/hello"
	// curl -i -p "http://localhost:8082/world"
	/*
		beanApp := NewApp()
		router := beanApp.GetRouter()
		router.Get("/hello", func(ctx *Context, next func()) {
			ctx.Response.Write([]byte("Hello"))
		})
		router.Post("/world", func(ctx *Context, next func()) {
			ctx.Response.Write([]byte("World"))
		})
		beanApp.ListenAndServe(":8082")
		//*/

	// Example 3: 在路由中使用中间件开启跨域
	// curl -i -X OPTIONS -H "origin: http://example.com" "http://localhost:8082/anything"
	/*
		beanApp := NewApp()
		router := beanApp.GetRouter()
		router.Use(func(ctx *Context, next func()) {
			r := ctx.Request
			w := ctx.Response

			if origin := r.Header.Get("origin"); origin != "" {
				headers := w.Header()
				headers.Set("Access-Control-Allow-Origin", origin)
				headers.Set("Access-Control-Allow-Credentials", "true")
				headers.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				headers.Set("Access-Control-Allow-Headers",
					"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

				log.Println("设置跨域规则")
			}

			if r.Method == "OPTIONS" {
				w.WriteHeader(204)

				log.Println("无条件响应OPTIONS请求")
			}

			next()
		})
		router.Get("/hello", func(ctx *Context, next func()) {
			ctx.Response.Write([]byte("Hello"))
		})
		beanApp.ListenAndServe(":8082")
		//*/

	// Example 4: 返回JSON数据
	// curl -i "http://localhost:8082/json"
	/*
		beanApp := NewApp()
		router := beanApp.GetRouter()
		router.Get("/json", func(ctx *Context, next func()) {
			type Ret struct {
				Code    int         `json:"code"`
				Message string      `json:"message"`
				Data    interface{} `json:"data"`
			}

			ctx.JSON(Ret{Code: 200, Message: "SUCCESS", Data: "Hello World"})
		})
		beanApp.ListenAndServe(":8082")
		//*/

	// Example 5: 处理404等异常请求
	// curl -i "http://localhost:8082/404NotFound"
	/*
		beanApp := NewApp()
		router := beanApp.GetRouter()
		router.Get("/normal", func(ctx *Context, next func()) {
			ctx.Response.WriteHeader(204)
			// 不要调用next, 否则就会走到404中间件
		})
		router.Use(func(ctx *Context, next func()) {
			ctx.Response.WriteHeader(404)
		})
		beanApp.ListenAndServe(":8082")
		//*/

	// Example 5: 处理404等异常请求
	// curl -i "http://localhost:8082/"
	/*
		beanApp := NewApp()
		router := beanApp.GetRouter()
		router.Get("/", func(ctx *Context, next func()) {
			ctx.String("Hello World")
		})
		router.Use(func(ctx *Context, next func()) {
			ctx.Response.WriteHeader(404)
		})
		beanApp.ListenAndServe(":8082")
		//*/

	// Example 6: 鉴权
	/*
		beanApp := NewApp()
		router := beanApp.GetRouter()
		router.Use(func(ctx *Context, next func()) {
			r := ctx.Request
			authr := r.Header.Get("Authorization")
			if authr != "123" {
				ctx.Response.WriteHeader(401)
				ctx.Response.Write([]byte("Failed"))
			} else {
				next()
			}
		})
		router.Get("/", func(ctx *Context, next func()) {
			ctx.String("Hello World")
		})
		beanApp.ListenAndServe(":8082")
		//*/
}
