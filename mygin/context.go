package mygin

//可以封装gin.context,这个就是根据实际需求，和业务场景来了比如一些header头，可以批量获取然后放到map里，也可以定制reply
import (
	"github.com/felix-xqs/gosdk/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

//没有特殊需求的话，可以不做
type MContext struct {
	*gin.Context
}

func NewMContext(c *gin.Context) *MContext {
	return &MContext{c}
}

//这些回复语句，都可以根据实际场景进行编写，比如回复detail与msg，Reply400，Reply403等等
func (c *MContext) Reply(httpCode int, obj interface{}) {
	c.JSON(httpCode, obj)
}
func (c *MContext) ReplyOK(data interface{}) {
	resp := lib.Resp{
		Code: lib.CodeOk,
		Data: data,
	}
	c.Reply(http.StatusOK, resp)
}

func (c *MContext) ReplyFail(code lib.ErrCode) {
	resp := &lib.Resp{
		Code: code,
		Msg:  lib.CodeMap[code],
	}
	c.Reply(http.StatusOK, resp)
}

//
