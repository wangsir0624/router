package router

import (
	"net/http"
)

type Session struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	route *Route //当前请求匹配到的路由
    
    routeParams map[string]string  //路由参数
}

func NewSession(w http.ResponseWriter, r *http.Request) *Session {
	session := new(Session)
	session.ResponseWriter = w
	session.Request = r

	return session
}

func (s *Session) GetCurrentRoute() *Route {
	return s.route
}

func (s *Session) GetRouteParam(key string) string {
    return ""
}
