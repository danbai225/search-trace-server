package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/ctx"
)

type cUserInfoRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int64  `json:"role"`
}

func cUserInfo() func(c *ghttp.Request) {
	return func(c *ghttp.Request) {
		func(c *ctx.Ctx) {
			u := c.GetUser()
			_ = c.Response.WriteJson(Msg{}.ok(cUserInfoRes{
				Name:  u.Name,
				Email: u.Email,
				Role:  u.Role,
			}))
		}(&ctx.Ctx{Request: c})
	}
}
