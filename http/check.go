package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"search-trace-server/server"
)

const v1TokenKey string = "token"

func checkV1(r *ghttp.Request) {
	token := r.GetHeader(v1TokenKey)
	if token == "" {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}
	user, err := server.UserGetUserByToken(token)
	if err != nil {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}
	r.SetCtxVar("user", user)
	r.Middleware.Next()
}
