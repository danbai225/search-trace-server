package http

type Msg struct {
	Msg  string      `json:"msg"`
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
}

const (
	okCode  = 0 //成功
	errCode = 1 //失败
)

func (Msg) ok(data interface{}) Msg {
	return Msg{Msg: "ok", Code: okCode, Data: data}
}
func (Msg) err(msg string) Msg {
	return Msg{Msg: msg, Code: errCode}
}
