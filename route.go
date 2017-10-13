package router

import (
	"net/http"
)

const (
	HttpMethodGet     = "GET"
	HttpMethodPost    = "POST"
	HttpMethodHead    = "HEAD"
	HttpMethodPatch   = "PATCH"
	HttpMethodPut     = "PUT"
	HttpMethodDelete  = "DELETE"
	HttpMethodTrace   = "TRACE"
	HttpMethodConnect = "CONNECT"
	HttpMethodOptions = "OPTIONS"
)

var (
	AllHttpMethods = []string{
		HttpMethodGet,
		HttpMethodPost,
		HttpMethodHead,
		HttpMethodPatch,
		HttpMethodPut,
		HttpMethodDelete,
		HttpMethodTrace,
		HttpMethodConnect,
		HttpMethodOptions,
	}
)

type RouteHandler interface {
	Handle(s *Session)
}

type RouteFuncHandler func(s *Session)

func (h RouteFuncHandler) Handle(s *Session) {
	h(s)
}

func NewRoute(methods []string, path string, handler RouteHandler) *Route {
	route := new(Route)
	route.methods = methods
	route.path = RegularPath(path)
	route.handler = handler

	return route
}

type Route struct {
	methods []string

	path string

	handler RouteHandler
}

func (r *Route) HasMethod(m string) bool {
	for _, tmpMethod := range r.methods {
		if m == tmpMethod {
			return true
		}
	}

	return false
}

func (r *Route) match(hr *http.Request) bool {
	//匹配path
	if RegularPath(hr.RequestURI) != r.path {
		return false
	}

	//匹配方法
	if !r.HasMethod(hr.Method) {
		return false
	}

	return true
}
