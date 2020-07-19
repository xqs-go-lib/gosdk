package mygin

//这个文件应该是单独写项目的，这里为了方便，直接写在这里，mysql，redis等，如果要写封装的话，也都要单独开项目
import (
	"github.com/go-mgo/mgo"
)

//这里主要是把获取连接，关闭连接等封装起来，避免手动操作有遗漏
type Client struct {
	Session *mgo.Session
	DBName  string
	ColName string
}

//就只写思想，不写实现了,可以做基本场景实现，最基本的单表的增删改查
func NewClient(session *mgo.Session, dbName string, colName string) (client *Client, err error) {
	return
}
func (client *Client) GetSessionAndCollection() (s *mgo.Session, c *mgo.Collection) {
	return
}
func (client *Client) Count(query interface{}) (n int, err error) {
	s, c := client.GetSessionAndCollection()
	defer s.Close()
	return c.Find(query).Count()
}
