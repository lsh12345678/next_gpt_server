package gateway

import "sync"

// epoller 对象 轮询器
type epoller struct {
	fd            int
	fdToConnTable sync.Map
}
