package client_imp

import (
	"errors"
	"fmt"
	client_int "github.com/alicia-oss/jinx/client/int"
	"github.com/alicia-oss/jinx/pkg/log"
)

func NewRouter() client_int.IRouter {
	return &router{handlerMap: make(map[uint32]client_int.IMsgHandle)}
}

type router struct {
	handlerMap map[uint32]client_int.IMsgHandle
}

func (r *router) Route(request client_int.IRequest) (client_int.IMsgHandle, bool) {

	handle, ok := r.handlerMap[request.GetMsgId()]
	return handle, ok
}

func (r *router) AddRouter(msgId uint32, router client_int.IMsgHandle) error {
	_, ok := r.handlerMap[msgId]
	if ok {
		log.Error(fmt.Sprintf("conflict msg_id:%v", msgId), ModuleNameClient)
		return errors.New("conflict msg_id")
	}
	r.handlerMap[msgId] = router
	return nil
}
