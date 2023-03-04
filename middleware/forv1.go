package middleware

import (
	gee "crazy-gin"
	"log"
	"time"
)

func Only4V1() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Println("this middleware is only for v1")
	}
}

func Logger4All() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()

		c.Next()

		since := time.Since(t)
		log.Printf("code[%d], url[%s], consume[%d]", c.StatusCode, c.Resp.URL.Path, since)
	}
}
