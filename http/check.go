package http

import (
	logs "github.com/danbai225/go-logs"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"search-trace-server/server"
)

const v1TokenKey string = "_token"

func checkV1(r *ghttp.Request) {
	token := r.GetHeader(v1TokenKey)
	logs.Info(token, 14)
	if token == "" {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}
	user, err := server.UserGetUserByToken(token)
	logs.Info(user, err, 20)
	if err != nil {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}
	r.SetCtxVar("user", user)
	r.Middleware.Next()
}
