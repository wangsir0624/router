package router

import (
	"net/http"
)

type Router struct {
	routes []*Route
}

func (r *Router) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}
