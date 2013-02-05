package infinitychannel

import (
	"container/list"
	"sync"
)

//This is a million times less good than
// https://github.com/kylelemons/iq/blob/master/iq_slice.go
// use that instead.

type infinitychannel struct {
	send    chan interface{}
	recv    chan interface{}
	hasData bool
	storage *list.List
	mutex   sync.Mutex
	cond    *sync.Cond
}

func New() (client_send, client_recv chan interface{}) {
	ic := &infinitychannel{}
	ic.send = make(chan interface{})
	ic.recv = make(chan interface{})
	ic.storage = list.New()
	var mutex sync.Mutex
	ic.cond = sync.NewCond(&mutex)
	go ic.receiveItems()
	go ic.sendItems()
	return ic.recv, ic.send
}

func (ic *infinitychannel) receiveItems() {
	for {
		item := <-ic.recv
		ic.mutex.Lock()
		ic.storage.PushBack(item)
		ic.hasData = true
		ic.mutex.Unlock()
		ic.cond.Signal()
	}
}

func (ic *infinitychannel) sendItems() {
	for {
		ic.cond.L.Lock()
		for ic.hasData == false {
			ic.cond.Wait()
		}
		ic.cond.L.Unlock()

		ic.mutex.Lock()
		item := ic.storage.Remove(ic.storage.Front())
		if ic.storage.Len() == 0 {
			ic.hasData = false
		}
		ic.mutex.Unlock()
		ic.send <- item
	}
}
