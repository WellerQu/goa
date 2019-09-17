package goa

// Method Http谓词
type Method string

const (
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
	POST    Method = "POST"
	GET     Method = "GET"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	ALL     Method = "ALL"
)

// RouteHandler 路由处理程序
type RouteHandler func(ctx *Context)

// namedHandler 具名路由处理程序生成器
func namedHandler(method Method, routePath string, handler RouteHandler) MiddlewareHandler {
	return func(ctx *Context, next func()) {
		ctx.Request.RoutePath = routePath
		reqPath := ctx.Request.URL.Path
		reqMethod := ctx.Request.Method

		if string(method) != reqMethod && method != ALL {
			next()
			return
		}

		// 尝试访问Params参数, 如果能取到参数则认为请求路径与路由匹配, 否则不匹配
		_, err := ctx.Request.Params()
		if reqPath != routePath && err != nil {
			next()
			return
		}

		handler(ctx)
	}
}

type router struct {
	app App
}

// AppRouter Web应用程序路由
type AppRouter interface {
	Use(middleware MiddlewareHandler)
	Get(path string, handler RouteHandler)
	Post(path string, handler RouteHandler)
}

// NewRouter 创建一个路由实例
func NewRouter(app App) AppRouter {
	r := &router{app}
	return r
}

func (r *router) Use(handler MiddlewareHandler) {
	r.app.Use(handler)
}

func (r *router) Post(path string, handler RouteHandler) {
	r.Use(namedHandler(POST, path, handler))
}

func (r *router) Get(path string, handler RouteHandler) {
	r.Use(namedHandler(GET, path, handler))
}

func (r *router) Put(path string, handler RouteHandler) {
	r.Use(namedHandler(PUT, path, handler))
}

func (r *router) Patch(path string, handler RouteHandler) {
	r.Use(namedHandler(PATCH, path, handler))
}

func (r *router) Delete(path string, handler RouteHandler) {
	r.Use(namedHandler(DELETE, path, handler))
}

func (r *router) All(path string, handler RouteHandler) {
	r.Use(namedHandler(ALL, path, handler))
}
