package dragon

import "github.com/julienschmidt/httprouter"

type DRouter struct {
	Router *httprouter.Router
}

// new drouter
func NewDRouter(router *httprouter.Router) *DRouter {
	return &DRouter{
		Router: router,
	}
}

func (dRouter *DRouter) GET(path string, handle func(ctx *HttpContext)) {
	dRouter.Router.Handle("GET", path, WrapController(handle))
}

func (dRouter *DRouter) HEAD(path string, handle func(ctx *HttpContext)) {
	dRouter.Router.Handle("HEAD", path, WrapController(handle))
}

func (dRouter *DRouter) OPTIONS(path string, handle func(ctx *HttpContext)) {
	dRouter.Router.Handle("OPTIONS", path, WrapController(handle))
}

func (dRouter *DRouter) POST(path string, handle func(ctx *HttpContext)) {
	dRouter.Router.Handle("POST", path, WrapController(handle))
}

func (dRouter *DRouter) PUT(path string, handle func(ctx *HttpContext)) {
	dRouter.Router.Handle("PUT", path, WrapController(handle))
}

func (dRouter *DRouter) PATCH(path string, handle func(ctx *HttpContext)) {
	dRouter.Router.Handle("PATCH", path, WrapController(handle))
}

func (dRouter *DRouter) DELETE(path string, handle func(ctx *HttpContext)) {
	dRouter.Router.Handle("DELETE", path, WrapController(handle))
}
