package lib

//定义一个基本的返回值的结构体
type Resp struct {
	Code   ErrCode     `json:"code"`
	Msg    string      `json:"msg,omitempty"`
	Detail string      `json:"detail,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
