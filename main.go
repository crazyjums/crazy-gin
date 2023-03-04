package main

import (
	gee "crazy-gin"
	"my-gin/middleware"
	"my-gin/util"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(middleware.Logger4All())

	r.GET("/test", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"code": 0,
			"msg":  "success",
		})
	})

	v1 := r.Group("/v1")
	v1.Use(middleware.Only4V1())
	v1.GET("/u", func(c *gee.Context) {
		c.JSON(http.StatusOK, util.Success(map[string]string{
			"uid": "12312",
		}))
	})

	_ = r.Run(":9001")
}
