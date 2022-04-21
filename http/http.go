package http

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"os"
)

func Start() {
	s := g.Server()
	s.SetPort(49492)
	s.BindMiddleware("/*", MiddlewareCORS)
	baseDir := "./dist"
	index := fmt.Sprint(baseDir, "/index.html")
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.ServeFile(index)
	})
	s.BindHandler("GET:/*", func(r *ghttp.Request) {
		path := fmt.Sprint(baseDir, r.Request.URL.Path)
		_, err := os.Stat(path)
		if err != nil {
			r.Response.ServeFile(index)
			return
		}
		r.Response.ServeFile(path)
	})
	s.BindHandler("/trace/add", cTraceAdd())
	api := s.Group("/api")
	api.POST("/get_token", cGetToken())
	v1 := api.Group("/v1")
	v1.Middleware(checkV1)

	userGroup := v1.Group("/user")
	userGroup.GET("/info", cUserInfo())

	traceGroup := v1.Group("/trace")
	traceGroup.POST("/add", cTraceAdd())
	traceGroup.GET("/search_keyword", cTraceSearchKeyword())

	wordGroup := v1.Group("/word")
	wordGroup.GET("/associate", cWordAssociate())

	blacklistGroup := v1.Group("/blacklist")
	blacklistGroup.GET("/list", cBlacklistList())
	blacklistGroup.POST("/add", cBlacklistAdd())
	s.Run()
}
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

type Page struct {
	PageNum   int
	PageSize  int
	PageCount int
}
