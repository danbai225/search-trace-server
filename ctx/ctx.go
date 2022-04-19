package ctx

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"search-trace-server/model"
)

type Ctx struct {
	*ghttp.Request
}

func (c *Ctx) GetUser() *model.User {
	value := c.Context().Value("user")
	if value != nil {
		return value.(*model.User)
	}
	return nil
}
