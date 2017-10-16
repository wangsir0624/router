package router

import (
	"testing"
)

func TestHttpMethodAllowed(t *testing.T) {
	data := []struct {
		input  string
		expect bool
	}{
		struct {
			input  string
			expect bool
		}{"GET", true},
		struct {
			input  string
			expect bool
		}{"POST", true},
		struct {
			input  string
			expect bool
		}{"HEAD", true},
		struct {
			input  string
			expect bool
		}{"PUT", true},
		struct {
			input  string
			expect bool
		}{"PATCH", true},
		struct {
			input  string
			expect bool
		}{"TEST", false},
	}

	for _, tmpData := range data {
		output := HttpMethodAllowed(tmpData.input)
		expect := tmpData.expect
		if output != expect {
			t.Errorf("want %s, but got %s\r\n", expect, output)
		}
	}
}

func TestHttpMethodsAllAllowed(t *testing.T) {
	data := []struct {
		input  []string
		expect bool
	}{
		struct {
			input  []string
			expect bool
		}{[]string{"GET", "POST"}, true},
		struct {
			input  []string
			expect bool
		}{[]string{"PATCH", "POST", "HEAD", "CONNECT"}, true},
		struct {
			input  []string
			expect bool
		}{[]string{"GET", "POST", "TEST"}, false},
	}

	for _, tmpData := range data {
		output := HttpMethodsAllAllowed(tmpData.input)
		expect := tmpData.expect
		if output != expect {
			t.Errorf("want %s, but got %s\r\n", expect, output)
		}
	}
}

func TestRegularPath(t *testing.T) {
	data := []struct {
		input  string
		expect string
	}{
		struct {
			input  string
			expect string
		}{"/", "/"},
		struct {
			input  string
			expect string
		}{"test", "/test"},
		struct {
			input  string
			expect string
		}{"/test", "/test"},
		struct {
			input  string
			expect string
		}{"test/", "/test"},
	}

	for _, tmpData := range data {
		output := RegularPath(tmpData.input)
		expect := tmpData.expect
		if output != expect {
			t.Errorf("want %s, but got %s\r\n", expect, output)
		}
	}
}
