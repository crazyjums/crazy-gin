package main

import (
	gee "crazy-gin"
	"fmt"
	"html/template"
	"my-gin/middleware"
	"my-gin/util"
	"net/http"
	"time"
)

func formatAdData(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func main() {
	r := gee.New()
	r.Use(middleware.Logger4All())
	r.SetFuncMap(template.FuncMap{
		"formatAsData": formatAdData,
	})
	r.LoadHtmlGlob("static/tmpl/*")
	r.Static("/assert", "./static")

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
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
