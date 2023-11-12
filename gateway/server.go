package gateway

//
//import (
//	"gpt/common/config"
//	"log"
//	"net"
//)
//
//// RunMain 启动网关服务
//func RunMain(path string) {
//	config.Init(path)
//	ln, err := net.ListenTCP("tcp", &net.TCPAddr{Port: config.GetGatewayTCPServerPort()})
//	if err != nil {
//		log.Fatalf("StartTCPEPollServer err:%s", err.Error())
//		panic(err)
//	}
//	initWorkPoll()
//
//}
