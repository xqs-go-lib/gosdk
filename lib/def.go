package lib

//用来定义一些确定的值，比如各个环境的代名词，以及可能需要的请求头的集合，方便定制context可以把所有的请求头一起取出来
const (
	// PROD 正式环境
	PROD = "prod"
	// DEV 开发环境，一般用于本机测试
	DEV = "dev"
	// TEST 测试环境
	TEST = "test"
	// PRE 预发布环境
	PRE = "pre"
)
const (
	TraceID = "traceID"
)
