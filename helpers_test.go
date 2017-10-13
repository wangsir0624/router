package router

import (
    "testing"
)

func TestRegularPath(t *testing.T) {
    data := []struct{
        input string
        expect string
    }{
        struct{
            input string
            expect string
        }{"/", "/"},
        struct{
            input string
            expect string
        }{"test", "/test"},
        struct{
            input string
            expect string
        }{"/test", "/test"},
        struct{
            input string
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