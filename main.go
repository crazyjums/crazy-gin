package main

import (
	gee "crazy-gin"
	"fmt"
	"net/http"
)

func v2() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<b>Hello crazyGin</b>")
	})
	r.Post("/getUser", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"name": "jums",
			"uid":  6666,
			"age":  18,
		})
	})
	r.AddRoute(http.MethodGet, "/h", func(c *gee.Context) {
		c.String(http.StatusOK, "I am string")
	})
	err := r.Run(":9001")
	if err != nil {
		fmt.Println(err)
	}
}

func test() {
	r := gee.NewRouter()
	r.AddRoute("GET", "/a", nil)
	r.AddRoute("GET", "/a/b", nil)
	r.AddRoute("GET", "/a/:id", nil)
	r.AddRoute("GET", "/a/:lang/c", nil)
	r.AddRoute("GET", "/a/*filepath", nil)

	fmt.Println(r)

	n, ps := r.GetRoute("GET", "/a/123456")
	fmt.Println("n=", n, ",pa=", ps)
}

func v3() {
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
	r.Get("/assert/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"filepath": c.Param("filepath"),
		})
	})
	r.Post("/a/:lang/c", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"lang": c.Param("lang"),
		})
	})
	err := r.Run(":9001")
	if err != nil {
		return
	}
}

func main() {
	//v2()
	v3()
	//test()
}
