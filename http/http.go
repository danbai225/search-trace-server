package http

import (
	"fmt"
	logs "github.com/danbai225/go-logs"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"io"
	"os"
	"search-trace-server/config"
)

func Start() {
	s := g.Server()
	arr := []io.Writer{
		logs.GetHttpWriter(),
	}
	if !config.C.Production {
		arr = append(arr, os.Stdout)
	}
	s.SetLogger(glog.NewWithWriter(io.MultiWriter(arr...)))
	s.SetPort(49492)
	s.BindMiddleware("/*", MiddlewareCORS)
	//前端路由
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
	//api路由
	api := s.Group("/api")
	api.POST("/get_token", cGetToken())
	v1 := api.Group("/v1")
	v1.Middleware(checkV1)

	//用户
	userGroup := v1.Group("/user")
	userGroup.GET("/info", cUserInfo())
	userGroup.POST("/add", cUserAdd())
	userGroup.GET("/list", cUserList())
	userGroup.POST("/del", cUserDel())

	//记录
	traceGroup := v1.Group("/trace")
	traceGroup.POST("/add", cTraceAdd())
	traceGroup.GET("/search_keyword", cTraceSearchKeyword())

	//词
	wordGroup := v1.Group("/word")
	wordGroup.GET("/associate", cWordAssociate())

	//黑名单
	blacklistGroup := v1.Group("/blacklist")
	blacklistGroup.GET("/list", cBlacklistList())
	blacklistGroup.POST("/add", cBlacklistAdd())
	blacklistGroup.POST("/del", cBlacklistDel())
	blacklistGroup.POST("/add_domain", cBlacklistAddDomain())
	blacklistGroup.POST("/del_domain", cBlacklistDelDomain())

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
