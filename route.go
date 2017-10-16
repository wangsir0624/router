package router

import (
	"net/http"
	"regexp"
	"strings"
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

var pathRegExp *regexp.Regexp

func init() {
	pathRegExp = regexp.MustCompile(`/\{([0-9A-Za-z_]+?)(?:(\:.*?))?\}`)
}

type RouteHandler interface {
	Handle(s *Session)
}

type RouteFuncHandler func(s *Session)

func (h RouteFuncHandler) Handle(s *Session) {
	h(s)
}

type Route struct {
	methods []string

	path string

	handler RouteHandler

	isRegExp bool           //是否是带参数路由
	regExp   *regexp.Regexp //匹配正则表达式
	params   []string       //参数列表
}

func NewRoute(methods []string, path string, handler RouteHandler) *Route {
	route := new(Route)
	route.methods = methods
	route.path = RegularPath(path)
	route.handler = handler
	route.isRegExp = false
	route.regExp = nil
	route.params = make([]string, 0)

	//是否是正则表达式路由
	params := pathRegExp.FindAllStringSubmatch(route.path, -1)
	if len(params) > 0 {
		route.isRegExp = true
		tmpRegExp := route.path
		for _, param := range params {
			route.params = append(route.params, param[1])
			if param[2] == "" {
				tmpRegExp = strings.Replace(tmpRegExp, param[0], "/([0-9A-Za-z_]+)", 1)
			} else if param[2] == ":" {
				tmpRegExp = strings.Replace(tmpRegExp, param[0], "(?:/([0-9A-Za-z_]+))?", 1)
			} else {
				tmpRegExp = strings.Replace(tmpRegExp, param[0], "/("+string(param[2][1:len(param[2])])+")", 1)
			}
		}

		route.regExp = regexp.MustCompile("^" + tmpRegExp + "$")
	}

	return route
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
	//匹配方法
	if !r.HasMethod(hr.Method) {
		return false
	}

	//匹配path
	if RegularPath(hr.URL.Path) != r.path {
		return false
	}

	return true
}

func (r *Route) matchRegExp(hr *http.Request) (bool, map[string]string) {
	params := make(map[string]string)

	//匹配方法
	if !r.HasMethod(hr.Method) {
		return false, params
	}

	path := RegularPath(hr.URL.Path)
	matched := r.regExp.MatchString(path)
	if !matched {
		return false, params
	} else {
		tmpParams := r.regExp.FindStringSubmatch(path)
		for i, p := range r.params {
			params[p] = tmpParams[i+1]
		}

		return true, params
	}
}
