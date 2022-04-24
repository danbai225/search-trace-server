package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/ctx"
	"search-trace-server/model"
	"search-trace-server/server"
)

func cBlacklistList() func(c *ghttp.Request) {
	return func(c *ghttp.Request) {
		func(c *ctx.Ctx) {
			list, err := server.BlacklistList(c.GetUser().Name)
			if err == nil {
				_ = c.Response.WriteJson(Msg{}.ok(list))
			} else {
				_ = c.Response.WriteJson(Msg{}.err("查询失败"))
			}
		}(&ctx.Ctx{Request: c})
	}
}

type cBlacklistAddReq struct {
	Id           int64  `json:"id"`
	Enable       *bool  `json:"enable"`
	Name         string `json:"name"`
	Mode         int8   `json:"mode" v:"max:2|min:1"`
	MatchPattern int8   `json:"match_pattern" v:"max:2|min:1"`
	Rules        string `json:"rules"`
}

func cBlacklistAdd() func(c *ghttp.Request) {
	return func(c *ghttp.Request) {
		func(c *ctx.Ctx) {
			req := &cBlacklistAddReq{}
			if err := c.Parse(req); err != nil {
				_ = c.Response.WriteJsonExit(Msg{
					Code: errCode,
					Msg:  err.Error(),
				})
				return
			}
			list, err := server.BlacklistAdd(&model.Blacklist{
				ID:           req.Id,
				UserName:     c.GetUser().Name,
				Name:         req.Name,
				Enable:       req.Enable,
				Mode:         req.Mode,
				MatchPattern: req.MatchPattern,
				Rules:        req.Rules,
			})
			if err == nil {
				_ = c.Response.WriteJson(Msg{}.ok(list))
			} else {
				_ = c.Response.WriteJson(Msg{}.err("修改失败"))
			}
		}(&ctx.Ctx{Request: c})
	}
}

type cBlacklistDelReq struct {
	Id int64 `json:"id" v:"required#必填id"`
}

func cBlacklistDel() func(c *ghttp.Request) {
	return func(c *ghttp.Request) {
		func(c *ctx.Ctx) {
			req := &cBlacklistDelReq{}
			if err := c.Parse(req); err != nil {
				_ = c.Response.WriteJsonExit(Msg{
					Code: errCode,
					Msg:  err.Error(),
				})
				return
			}
			res, err := server.BlacklistDelete(req.Id)
			if err == nil {
				_ = c.Response.WriteJson(Msg{}.ok(res))
			} else {
				_ = c.Response.WriteJson(Msg{}.err("删除失败"))
			}
		}(&ctx.Ctx{Request: c})
	}
}
