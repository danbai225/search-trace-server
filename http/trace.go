package http

import (
	"github.com/gogf/gf/v2/frame/g"
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
		name := ""
		u := r.Context().Value("user")
		if u == nil {
			name = "admin"
		} else {
			name = u.(*model.User).Name
		}
		trace2 := model.Trace{
			UserName: name,
			Title:    trace.Title,
			Url:      trace.Url,
			Content:  trace.Content,
		}
		err := server.TraceCreate(&trace2)
		if err == nil {
			_ = r.Response.WriteJson(Msg{}.ok(nil))
		} else {
			_ = r.Response.WriteJson(Msg{}.err("添加失败"))
		}
	}
}

type cTraceSearchKeywordReq struct {
	Key      string `json:"key" v:"required|min-length:2#必填关键字|长度最小2"`
	PageSize int    `json:"page_size" v:"required|min:1|max:100#必填页大小|页面大小最小为1|页面大小最大为100"`
	PageNum  int    `json:"page_num" v:"required|min:1#必填页数|页号最小为1"`
}

func cTraceSearchKeyword() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		req := &cTraceSearchKeywordReq{}
		if err := r.Parse(req); err != nil {
			_ = r.Response.WriteJsonExit(Msg{
				Code: errCode,
				Msg:  err.Error(),
			})
			return
		}
		u := r.Context().Value("user").(*model.User)
		list, PageTotal, err := server.TraceSearchForKeyword(u.Name, req.Key, req.PageSize, req.PageNum)
		if err == nil {
			_ = r.Response.WriteJson(Msg{}.ok(g.Map{
				"list":       list,
				"page_total": PageTotal,
				"page_num":   req.PageNum,
			}))
		} else {
			_ = r.Response.WriteJson(Msg{}.err("搜索失败"))
		}
	}
}
