package domain

import (
	"gpt/ipconf/source"
	"sort"
	"sync"
)

type Dispatcher struct {
	candidateTable map[string]*Endport
	sync.RWMutex
}

var dp *Dispatcher

func Init() {
	dp = &Dispatcher{}
	dp.candidateTable = make(map[string]*Endport)
	go func() {
		// 通过eventChan处理etcd中的节点变化事件
		for event := range source.EventChan() {
			switch event.Type {
			case source.AddNodeEvent:
				dp.addNode(event)
			case source.DelNodeEvent:
				dp.delNode(event)
			}
		}
	}()
}

func Dispatch(ctx *IpConfConext) []*Endport {
	// Step1: 获得候选endPort
	// 从所有有记录的候选服务器中选出来
	eds := dp.getCandidateEndport(ctx)
	// Step2: 逐一计算得分
	for _, ed := range eds {
		ed.CalculateScore(ctx)
	}
	// Step3: 全局排序，动静结合的排序策略
	sort.Slice(eds, func(i, j int) bool {
		// 优先基于活跃分数进行排序
		if eds[i].ActiveSource > eds[j].ActiveSource {
			return true
		}
		// 如果活跃分数相同，则使用静态分数排序
		if eds[i].ActiveSource == eds[j].ActiveSource {
			if eds[i].StaticSource > eds[j].StaticSource {
				return true
			}
			return false
		}
		return false
	})
	return eds
}

func (dp *Dispatcher) getCandidateEndport(ctx *IpConfConext) []*Endport {
	dp.RLock()
	defer dp.RUnlock()
	candidateList := make([]*Endport, 0, len(dp.candidateTable))
	for _, ed := range dp.candidateTable {
		candidateList = append(candidateList, ed)
	}
	return candidateList
}

func (dp *Dispatcher) delNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	delete(dp.candidateTable, event.Key())
}

// addNode 每个主机一个窗口，计算最近5次的一个静态动态分
func (dp *Dispatcher) addNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	var (
		ed *Endport
		ok bool
	)
	if ed, ok = dp.candidateTable[event.Key()]; !ok { // 不存在
		ed = NewEndport(event.IP, event.Port)
		dp.candidateTable[event.Key()] = ed
	}
	ed.UpdateStat(&Stat{
		ConnectNum:   event.ConnectNum,
		MessageBytes: event.MessageBytes,
	})

}
