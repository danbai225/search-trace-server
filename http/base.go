package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/server"
)

type cGetTokenReq struct {
	Name string `json:"name" v:"required#请输入账号"`
	Pass string `json:"pass" v:"required#请输入密码"`
}

func cGetToken() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		req := &cGetTokenReq{}
		if err := r.Parse(req); err != nil {
			r.Response.WriteJsonExit(Msg{
				Code: errCode,
				Msg:  err.Error(),
			})
			return
		}
		msg := Msg{}
		token, _, err1 := server.LoginByPass(req.Name, req.Pass)
		if err1 == nil {
			msg = msg.ok(token)
		} else {
			msg = msg.err("验证不通过")
		}
		r.Response.WriteJson(msg)
	}
}
