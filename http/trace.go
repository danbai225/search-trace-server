package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/model"
	"search-trace-server/server"
)

func cTraceAdd() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		trace := &model.Trace{}
		if err := r.Parse(trace); err != nil {
			return
		}
		server.AddTrace(trace)
	}
}
