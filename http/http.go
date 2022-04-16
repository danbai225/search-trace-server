package http

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Start() {
	s := g.Server()
	s.SetPort(49492)
	s.BindMiddleware("/*", MiddlewareCORS)
	s.BindHandler("/trace/add", cTraceAdd())
	api := s.Group("/api")
	api.POST("/get_token", cGetToken())
	v1 := api.Group("/v1")
	v1.Middleware(checkV1)
	v1.GET("/user/info", cUserInfo())
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
