package util

import (
	gee "crazy-gin"
)

func Response(c int, m string, d interface{}) gee.H {
	return gee.H{
		"code": c,
		"msg":  m,
		"data": d,
	}
}

func Success(d gee.Any) gee.H {
	return gee.H{
		"code": 0,
		"msg":  "success",
		"data": d,
	}
}
