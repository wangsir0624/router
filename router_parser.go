package router

import (
	"io"
)

type RouterParser interface {
	func Parse(r io.Reader) *Router
}

type CsvRouterParser struct {
	
}

func (p *CsvRouterParser) Parse(r io.Reader) *Router {
	
}
