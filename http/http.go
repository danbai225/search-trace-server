package http

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Start() {
	s := g.Server()
	s.SetPort(49492)
	s.BindMiddleware("/*", MiddlewareCORS)
	baseDir := "./dist/"
	s.BindHandler("GET://{name}.html", func(r *ghttp.Request) {
		r.Response.ServeFile(fmt.Sprint(baseDir, r.Get("name").String(), ".html"))
	})
	s.BindHandler("GET://{name}.js", func(r *ghttp.Request) {
		r.Response.ServeFile(fmt.Sprint(baseDir, r.Get("name").String(), ".js"))
	})
	s.BindHandler("GET://{name}.css", func(r *ghttp.Request) {
		r.Response.ServeFile(fmt.Sprint(baseDir, r.Get("name").String(), ".css"))
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
