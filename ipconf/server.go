package ipconf

import (
	"gpt/common/config"
	"gpt/ipconf/domain"
	"gpt/ipconf/source"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// RunMain 启动web容器
func RunMain(path string) {
	config.Init(path)
	source.Init()
	domain.Init()
	s := server.Default(server.WithHostPorts(":6789"))
	s.GET("/ip/list", GetIpInfoList)
	s.Spin()
}
