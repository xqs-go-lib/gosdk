package lib

//定义返回码
type ErrCode int

//这里是
const (
	CodeOk     ErrCode = 1
	CodePara           = 2
	CodeLogOut         = 3
	CodeBan            = 4
)

var CodeMap = map[ErrCode]string{
	CodePara:   "参数错误",
	CodeLogOut: "用户未登陆",
	CodeBan:    "用户已被封禁",
}
