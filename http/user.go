package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/ctx"
	"search-trace-server/model"
	"search-trace-server/server"
)

const admin = 1

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

type cUserAddReq struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" v:"length:1,128|required"`
	Email string `json:"email" v:"email|required"`
	Pass  string `json:"pass" v:"required|length:6,128"`
}

func cUserAdd() func(c *ghttp.Request) {
	return func(c *ghttp.Request) {
		func(c *ctx.Ctx) {
			u := c.GetUser()
			if u.Role != admin {
				_ = c.Response.WriteJsonExit(Msg{
					Code: errCode,
					Msg:  "无权限",
				})
			}
			req := &cUserAddReq{}
			if err := c.Parse(req); err != nil {
				_ = c.Response.WriteJsonExit(Msg{
					Code: errCode,
					Msg:  err.Error(),
				})
				return
			}
			res, err := server.UserCreate(&model.User{
				ID:    req.ID,
				Name:  req.Name,
				Email: req.Email,
				Pass:  req.Pass,
				Role:  0,
			})
			if err == nil {
				_ = c.Response.WriteJson(Msg{}.ok(res))
			} else {
				_ = c.Response.WriteJson(Msg{}.err("创建失败"))
			}
		}(&ctx.Ctx{Request: c})
	}
}

func cUserList() func(c *ghttp.Request) {
	return func(c *ghttp.Request) {
		func(c *ctx.Ctx) {
			u := c.GetUser()
			if u.Role != admin {
				_ = c.Response.WriteJsonExit(Msg{
					Code: errCode,
					Msg:  "无权限",
				})
			}
			res, err := server.UserList()
			if err == nil {
				_ = c.Response.WriteJson(Msg{}.ok(res))
			} else {
				_ = c.Response.WriteJson(Msg{}.err("获取失败"))
			}
		}(&ctx.Ctx{Request: c})
	}
}

type cUserDelReq struct {
	Id int64 `json:"id" v:"required#必填id"`
}

func cUserDel() func(c *ghttp.Request) {
	return func(c *ghttp.Request) {
		func(c *ctx.Ctx) {
			req := &cUserDelReq{}
			if err := c.Parse(req); err != nil {
				_ = c.Response.WriteJsonExit(Msg{
					Code: errCode,
					Msg:  err.Error(),
				})
				return
			}
			res, err := server.UserDelete(req.Id)
			if err == nil {
				_ = c.Response.WriteJson(Msg{}.ok(res))
			} else {
				_ = c.Response.WriteJson(Msg{}.err("删除失败"))
			}
		}(&ctx.Ctx{Request: c})
	}
}
