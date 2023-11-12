package jsons

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
)

type statusCode int

const (
	successStatusCode statusCode = 0
	failStatusCode    statusCode = 1
	noDataStatusCode  statusCode = 2
)

type baseResp struct {
	Code   statusCode  `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func SetJSONRespErr(c *app.RequestContext, err error) {
	msg := err.Error()
	r := &baseResp{
		Code: failStatusCode,
		Msg:  msg,
	}
	logger.Errorf("request failed, path:%v, query:%v, body:%v, Err:%v", string(c.Request.Path()),
		string(c.Request.QueryString()),
		string(c.Request.Body()), err)
	c.PureJSON(500, r)
}

func SetJSONRespSuccess(c *app.RequestContext, result interface{}) {
	r := &baseResp{
		Code:   successStatusCode,
		Msg:    "success",
		Result: result,
	}
	c.PureJSON(200, r)
}
