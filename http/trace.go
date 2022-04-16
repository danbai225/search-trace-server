package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/model"
	"search-trace-server/server"
)

type cTraceAddReq struct {
	Title   string `json:"title"  v:"required#必填Title"`
	Url     string `json:"url" v:"url"`
	Content string `json:"content" v:"required#Content"`
}

func cTraceAdd() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		trace := &cTraceAddReq{}
		if err := r.Parse(trace); err != nil {
			_ = r.Response.WriteJsonExit(Msg{
				Code: errCode,
				Msg:  err.Error(),
			})
			return
		}
		u := r.Context().Value("user").(*model.User)
		trace2 := model.Trace{
			UserName: u.Name,
			Title:    trace.Title,
			Url:      trace.Url,
			Content:  trace.Url,
		}
		err := server.TraceCreate(&trace2)
		if err == nil {
			_ = r.Response.WriteJson(Msg{}.ok(nil))
		} else {
			_ = r.Response.WriteJson(Msg{}.err("添加失败"))
		}
	}
}
