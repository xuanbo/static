package main

import (
	"flag"

	"github.com/gin-gonic/gin"

	"github.com/xuanbo/static/pkg/api"
	v1 "github.com/xuanbo/static/pkg/api/v1"
)

var addr string

func main() {
	flag.StringVar(&addr, "addr", ":12345", "监听地址, 默认 :12345")
	flag.Parse()

	router := gin.Default()

	// 静态资源
	{
		router.Static("/project", "./project")
		router.Static("/static", "./static")
	}

	// 首页
	{
		router.GET("/", api.Index)
		router.GET("/index", api.Index)
		router.GET("/index.htm", api.Index)
		router.GET("/index.html", api.Index)
	}

	// 接口
	{
		router.GET("/api/v1/tree", v1.Tree)
		router.POST("/api/v1/upload", v1.Upload)
		router.POST("/api/v1/remove", v1.Remove)
	}

	router.Run(addr)
}
