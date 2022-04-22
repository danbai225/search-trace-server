package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/server"
)

type cWordAssociateReq struct {
	Word string `json:"word" v:"required|min-length:2#必填|长度最小2"`
}

func cWordAssociate() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		req := &cWordAssociateReq{}
		if err := r.Parse(req); err != nil {
			_ = r.Response.WriteJsonExit(Msg{
				Code: errCode,
				Msg:  err.Error(),
			})
			return
		}
		list, err := server.WordAssociate(req.Word)
		if err == nil {
			_ = r.Response.WriteJson(Msg{}.ok(list))
		} else {
			_ = r.Response.WriteJson(Msg{}.err("搜索失败"))
		}
	}
}
