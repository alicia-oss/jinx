package client_imp

import (
	"fmt"
	client_int "jinx/client/int"
	"jinx/pkg/log"
)

type BaseMsgHandler struct{}

func (b *BaseMsgHandler) PreHandle(req client_int.IRequest) {
	log.Info(fmt.Sprintf("client PreHandle req: %v...", string(req.GetData())), "handler")
}

func (b *BaseMsgHandler) Handle(req client_int.IRequest) {
	log.Info("client Handle...", "handler")
}

func (b *BaseMsgHandler) PostHandle(req client_int.IRequest) {
	log.Info("client PostHandle...", "handler")
}
