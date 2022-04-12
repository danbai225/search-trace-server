package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/modle"
	"search-trace-server/server"
	"time"
)

func cTraceAdd() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		trace := &modle.Trace{}
		if err := r.Parse(trace); err != nil {
			return
		}
		trace.BrowseTime = time.Now()
		server.AddTrace(trace)
	}
}
