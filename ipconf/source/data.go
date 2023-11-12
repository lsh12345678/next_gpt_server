package source

import (
	"context"

	"gpt/common/config"
	"gpt/common/discovery"

	"github.com/bytedance/gopkg/util/logger"
)

func Init() {
	eventChan = make(chan *Event)
	ctx := context.Background()
	go DataHandler(&ctx)
	if config.IsDebug() {
		ctx := context.Background()
		testServiceRegister(&ctx, "7896", "node1")
		testServiceRegister(&ctx, "7897", "node2")
		testServiceRegister(&ctx, "7898", "node3")
	}
}

func DataHandler(ctx *context.Context) {
	dis := discovery.NewServiceDiscovery(ctx)
	defer dis.Close()
	setFunc := func(key, value string) {
		ed, err := discovery.UnMarshal([]byte(value))
		if err != nil {
			logger.CtxErrorf(*ctx, "DataHandler.setFunc.err :%s", err.Error())
		}
		if event := NewEvent(ed); ed != nil {
			event.Type = AddNodeEvent
			eventChan <- event
		}
	}
	delFunc := func(key, value string) {
		ed, err := discovery.UnMarshal([]byte(value))
		if err != nil {
			logger.CtxErrorf(*ctx, "DataHandler.delFunc.err :%s", err.Error())
		}
		if event := NewEvent(ed); ed != nil {
			event.Type = DelNodeEvent
			eventChan <- event
		}
	}
	if err := dis.WatchService(config.GetServicePathForIPConf(), setFunc, delFunc); err != nil {
		logger.CtxErrorf(*ctx, "DataHandler.dis.watchService.err :%s", err.Error())
	}
}
