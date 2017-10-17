# router
Golang用的web router

## Usage

### 基本用法
```
package main

import (
	"fmt"
	"net/http"
	"router"
)

func main() {

	mux := router.NewRouter()
	mux.AnyFunc("/hello", hello)
	mux.GetFunc("/show/{name}/{age:[0-9]+?}", show)

	http.ListenAndServe(":8080", mux)
}

func hello(s *router.Session) {
	//获取response writer
	s.ResponseWriter.Write([]byte("hello"))
	
	//获取请求
	request := s.Request
	
	//获取当前匹配的路由
	currentRoute := s.GetCurrentRoute()
	
	fmt.Println(request, currentRoute)
}

func show(s *router.Session) {
	//获取路由参数
	name := s.GetRouteParam("name")
	age := s.GetRouteParam("age")

	fmt.Printf("My name is %s, and I'm %s years old.\r\n", name, age)
}

```
每次请求都会生成一个独立的session，不同的session之间互不影响

### 添加路由
```
mux.GetFunc("/hello", hello)   //添加GET方法路由
mux.PostFunc("/hello", hello)  //添加POST方法路由
mux.HeadFunc("/hello", hello)  //添加HEAD方法路由，其他方法以此类推

//多方法路由
mux.AddFunc([]string{"GET", "POST"}, "/hello", hello)

//所有方法路由
mux.AnyFunc("/hello", hello)
```

### 带参数路由
```
mux.GetFunc("/show/{name}/{age}", show)

//可选参数
mux.GetFunc("/show/{name}/{age:}", show)

//正则约束
mux.GetFunc("/show/{name}/{age:[0-9]+?}", show)
```

