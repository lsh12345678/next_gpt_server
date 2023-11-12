package gpt

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"gpt/common/jsons"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
)

var client *http.Client
var once sync.Once

const openaiAPIURL = "http://nb.nextweb.fun/api/proxy/v1/chat/completions"
const openaiAPIKey = "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"

func init() {
	once.Do(func() {
		client = http.DefaultClient
	})
}

type ModelContent struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CallModelReq struct {
	Model       string         `json:"model"`
	Messages    []ModelContent `json:"messages"`
	Stream      bool           `json:"stream"`
	Temperature float64        `json:"temperature"`
}

type CallModelResp struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

// CallModel API 适配应用层
func CallModel(c context.Context, ctx *app.RequestContext) {
	data := ctx.Request.Body()
	logger.CtxInfof(c, "receive call model request, req:%v", data)

	//var req *CallModelReq
	//if err := json.Unmarshal(data, &req); err != nil {
	//	jsons.SetJSONRespErr(ctx, err)
	//	return
	//}
	//
	req, err := http.NewRequest("POST", openaiAPIURL, bytes.NewBuffer(data))
	if err != nil {
		jsons.SetJSONRespErr(ctx, err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		jsons.SetJSONRespErr(ctx, err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	jsons.SetJSONRespSuccess(ctx, body)

}
