package jinx_imp

import (
	"errors"
	"fmt"
	"jinx/jinx_int"
	"jinx/pkg/log"
)

func NewRouter() jinx_int.IRouter {
	return &router{handlerMap: make(map[uint32]jinx_int.IMsgHandle)}
}

type router struct {
	handlerMap map[uint32]jinx_int.IMsgHandle
}

func (r *router) Route(request jinx_int.IRequest) (jinx_int.IMsgHandle, bool) {

	handle, ok := r.handlerMap[request.GetMsgId()]
	return handle, ok
}

func (r *router) AddRouter(msgId uint32, router jinx_int.IMsgHandle) error {
	_, ok := r.handlerMap[msgId]
	if ok {
		log.Error(fmt.Sprintf("conflict msg_id:%v", msgId), ModuleNameRouter)
		return errors.New("conflict msg_id")
	}
	r.handlerMap[msgId] = router
	return nil
}
