package router

import (
	"strings"
)

func HttpMethodAllowed(m string) bool {
	for _, method := range AllHttpMethods {
		if m == method {
			return true
		}
	}

	return false
}

func RegularPath(path string) string {
	if strings.HasSuffix(path, "/") {
		path = string(path[:len(path)-1])
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return path
}
