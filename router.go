package router

import (
	"fmt"
	"net/http"
	"sync"
)

type Router struct {
	routes     []*Route
	routeMutex *sync.Mutex
}

func NewRouter() *Router {
	router := new(Router)
	router.routes = make([]*Route, 0)
	router.routeMutex = new(sync.Mutex)

	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, hr *http.Request) {
	//匹配路由
	var session *Session
	session = r.match(w, hr)
	if session == nil {
		http.NotFoundHandler().ServeHTTP(w, hr)
	} else {
		session.route.handler.Handle(session)
	}
}

func (r *Router) Get(path string, handler RouteHandler) *Route {
	return r.Add([]string{"GET"}, path, handler)
}

func (r *Router) GetFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"GET"}, path, f)
}

func (r *Router) Post(path string, handler RouteHandler) *Route {
	return r.Add([]string{"POST"}, path, handler)
}

func (r *Router) PostFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"POST"}, path, f)
}

func (r *Router) Head(path string, handler RouteHandler) *Route {
	return r.Add([]string{"HEAD"}, path, handler)
}

func (r *Router) HeadFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"HEAD"}, path, f)
}

func (r *Router) Patch(path string, handler RouteHandler) *Route {
	return r.Add([]string{"PATCH"}, path, handler)
}

func (r *Router) PatchFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"PATCH"}, path, f)
}

func (r *Router) Put(path string, handler RouteHandler) *Route {
	return r.Add([]string{"PUT"}, path, handler)
}

func (r *Router) PutFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"PUT"}, path, f)
}

func (r *Router) Delete(path string, handler RouteHandler) *Route {
	return r.Add([]string{"DELETE"}, path, handler)
}

func (r *Router) DeleteFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"DELETE"}, path, f)
}

func (r *Router) Trace(path string, handler RouteHandler) *Route {
	return r.Add([]string{"TRACE"}, path, handler)
}

func (r *Router) TraceFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"TRACE"}, path, f)
}

func (r *Router) Connect(path string, handler RouteHandler) *Route {
	return r.Add([]string{"CONNECT"}, path, handler)
}

func (r *Router) ConnectFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"CONNECT"}, path, f)
}

func (r *Router) Options(path string, handler RouteHandler) *Route {
	return r.Add([]string{"OPTIONS"}, path, handler)
}

func (r *Router) OptionsFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc([]string{"OPTIONS"}, path, f)
}

func (r *Router) Any(path string, handler RouteHandler) *Route {
	return r.Add(AllHttpMethods, path, handler)
}

func (r *Router) AnyFunc(path string, f func(s *Session)) *Route {
	return r.AddFunc(AllHttpMethods, path, f)
}

func (r *Router) Add(methods []string, path string, handler RouteHandler) *Route {
	//检查方法是否合法
	if !HttpMethodsAllAllowed(methods) {
		panic(fmt.Errorf("unsupported http method"))
	}

	route := NewRoute(methods, path, handler)

	r.routeMutex.Lock()
	r.routes = append(r.routes, route)
	r.routeMutex.Unlock()

	return route
}

func (r *Router) AddFunc(methods []string, path string, f func(s *Session)) *Route {
	handler := RouteFuncHandler(f)

	return r.Add(methods, path, handler)
}

func (r *Router) match(w http.ResponseWriter, hr *http.Request) *Session {
	for _, route := range r.routes {
		if !route.isRegExp {
			if route.match(hr) {
				session := NewSession(w, hr)
				session.route = route
				return session
			}
		} else {
			matched, params := route.matchRegExp(hr)
			if matched {
				session := NewSession(w, hr)
				session.route = route
				session.routeParams = params
				return session
			}
		}
	}

	return nil
}
