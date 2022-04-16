package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/model"
)

type cUserInfoRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int64  `json:"role"`
}

func cUserInfo() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		u := r.GetCtxVar("user").Interface().(*model.User)
		_ = r.Response.WriteJson(Msg{}.ok(cUserInfoRes{
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}))
	}
}
