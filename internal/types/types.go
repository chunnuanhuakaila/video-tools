// Code generated by goctl. DO NOT EDIT.
package types

type CommonRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ParseReq struct {
	URL string `form:"url"`
}

type Data struct {
	Author string `json:"author"`
	Avatar string `json:"avatar"`
	Time   string `json:"time"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Url    string `json:"url"`
}

type ParseResult struct {
	*Data
}
