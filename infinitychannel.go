package infinitychannel

import (
  "container/list"
  "sync"
)

type Infinitychannel struct {
  send chan interface{}
  recv chan interface{}
  storage *list.List
  cond *sync.Cond
}

func New() (send, recv chan interface{}) {
  var mutex sync.Mutex
  cond := sync.NewCond(&mutex)
  ic := &Infinitychannel{make(chan interface{}), make(chan interface{}), list.New(), cond}
  go ic.receiveItems()
  go ic.sendItems()
  return ic.send, ic.recv
}

func (ic *Infinitychannel) receiveItems() {
  for {
    item := <- ic.recv
    ic.storage.PushBack(item)
    ic.cond.Signal()
  }
}

// I think there's still a race condition here, but it's close
func (ic *Infinitychannel) sendItems() {
  for {
    ic.cond.L.Lock()
    for ic.storage.Front() == nil {
      ic.cond.Wait()
    }
    ic.cond.L.Unlock()

    item := ic.storage.Remove(ic.storage.Front())
    ic.send <- item
  }
}

