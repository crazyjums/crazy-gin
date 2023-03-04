# CRAZY GIN

> 大部分时候，我们需要实现一个 Web 应用，第一反应是应该使用哪个框架。不同的框架设计理念和提供的功能有很大的差别。比如 Python 语言的 django和flask，前者大而全，后者小而美。Go语言/golang 也是如此，新框架层出不穷，比如Beego，Gin，Iris等。那为什么不直接使用标准库，而必须使用框架呢？在设计一个框架之前，我们需要回答框架核心为我们解决了什么问题。只有理解了这一点，才能想明白我们需要在框架中实现什么功能。

学习go web框架gin的核心原理。参考连接：[link](https://geektutu.com/post/gee.html)

# Quick Start

main.go:
```go
package main

import (
	gee "crazy-gin"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/user/me", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"name": "me",
			"uid":  123456,
		})
	})
	r.Get("/hello/:id", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"id": c.Param("id"),
		})
	})
	r.Get("/static/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"filepath": c.Param("filepath"),
		})
	})
	r.Get("/a/:lang/c", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"lang": c.Param("lang"),
		})
	})
	err := r.Run(":9001")
	if err != nil {
		return
	}
}
```

curl:
```BASH
$ curl "http://localhost:9001/hello/1234"
StatusCode        : 200
StatusDescription : OK
Content           : {"id":"1234"}
$ curl "http://localhost:9001/a/python/c"
StatusCode        : 200                                                                                            
StatusDescription : OK
Content           : {"lang":"python"}
$ curl "http://localhost:9001/assert/css/index.css"
StatusCode        : 200
StatusDescription : OK
Content           : {"filepath":"css/index.css"}
$ curl "http://localhost:9001/user/me"
StatusCode        : 200
StatusDescription : OK
Content           : {"name":"me","uid":123456}
```