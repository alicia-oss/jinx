package jinx_imp

import (
	"errors"
	"github.com/alicia-oss/jinx/jinx_int"
	"github.com/alicia-oss/jinx/pkg/log"
	"sync"
)

func NewConnManager() jinx_int.IConnManager {
	return &connManager{
		cMap:      &sync.Map{},
		closeChan: make(chan struct{}),
	}
}

type connManager struct {
	cMap      *sync.Map
	closeChan chan struct{}
}

func (cm *connManager) Add(conn jinx_int.IConnection) error {
	select {
	case <-cm.closeChan:
		return errors.New("connManager has closed")
	default:
		cm.cMap.Store(conn.GetConnID(), conn)
		return nil
	}
}

func (cm *connManager) Remove(conn jinx_int.IConnection) {
	cm.cMap.Delete(conn.GetConnID())
}

func (cm *connManager) Get(connID uint32) (jinx_int.IConnection, bool) {
	load, ok := cm.cMap.Load(connID)
	if ok {
		return load.(jinx_int.IConnection), ok
	}
	return nil, ok
}

func (cm *connManager) Len() int {
	length := 0
	cm.cMap.Range(func(k, v interface{}) bool {
		length++
		return true
	})
	return length
}

func (cm *connManager) ClearConn() {
	log.Info("Connection Manager are closing ......", ModuleNameConnManager)
	close(cm.closeChan)
	cm.cMap.Range(func(_, value any) bool {
		conn := value.(jinx_int.IConnection)
		conn.Stop()
		return true
	})
}
