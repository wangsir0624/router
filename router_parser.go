package router

import (
	"io"
)

type RouterParser interface {
	func Parse(r io.Reader) *Router
}

type CsvRouterParser struct {
	delimiter string
}

func NewCsvRouterParser() *CsvRouterParser {
	parser := new(CscRouterParser)
	parser.delimiter = "\t"
	
	return parser
}

func (p *CsvRouterParser) Parse(r io.Reader) *Router, err {
	
}

func (p *CsvRouterParser) SetDelimiter(d string) {
	
}
