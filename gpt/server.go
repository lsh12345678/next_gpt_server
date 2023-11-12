package gpt

import "github.com/cloudwego/hertz/pkg/app/server"

// RunMain 启动web容器
func RunMain(path string) {

	s := server.Default(server.WithHostPorts(":6790"))
	s.POST("/call_model", CallModel)
	s.Spin()
}
