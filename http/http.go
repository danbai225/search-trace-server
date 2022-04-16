package http

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Start() {
	s := g.Server()
	s.SetPort(49492)
	s.BindMiddleware("/*", MiddlewareCORS)
	api := s.Group("/api")
	api.POST("/get_token", cGetToken())

	v1 := api.Group("/v1")
	v1.Middleware(checkV1)

	userGroup := v1.Group("/user")
	userGroup.GET("/info", cUserInfo())

	traceGroup := v1.Group("/trace")
	traceGroup.POST("/add", cTraceAdd())

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
