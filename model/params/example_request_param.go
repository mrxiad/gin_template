package params

/*
这个定义请求参数的结构体,用于json请求的参数解析
*/

type ExampleRequestParam struct {
	// 这里定义请求参数的字段
	Id int64 `json:"id"`
}

//其他一些关于example的请求参数结构体
