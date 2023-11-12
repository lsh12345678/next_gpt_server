package domain

import (
	"sync"
)

type Endport struct {
	IP           string       `json:"ip"`
	Port         string       `json:"port"`
	ActiveSource float64      `json:"-"`
	StaticSource float64      `json:"-"`
	Stats        *Stat        `json:"-"`
	window       *stateWindow `json:"-"`
	lock         sync.Mutex   `json:"-"`
}

func NewEndport(ip, port string) *Endport {
	ed := &Endport{
		IP:   ip,
		Port: port,
	}
	ed.window = newStateWindow()
	ed.Stats = ed.window.getStat()
	go func() {
		for stat := range ed.window.statChan {
			ed.window.appendStat(stat)
			newStat := ed.window.getStat()
			ed.Stats = newStat
			//atomic.SwapPointer((*unsafe.Pointer)((unsafe.Pointer)(ed.Stats)), unsafe.Pointer(newStat))
		}
	}()
	return ed
}

func (ed *Endport) UpdateStat(s *Stat) {
	ed.window.statChan <- s
}

func (ed *Endport) CalculateScore(ctx *IpConfConext) {
	// 如果 stats 字段是空的，则直接使用上一次计算的结果，此次不更新
	if ed.Stats != nil {
		ed.ActiveSource = ed.Stats.CalculateActiveSource()
		ed.StaticSource = ed.Stats.CalculateStaticSource()
	}
}
