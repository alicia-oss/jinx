package jinx_imp

import (
	"fmt"
	"jinx/jinx_int"
	"jinx/jinx_pkg/log"
)

type BaseMsgHandler struct{}

func (b *BaseMsgHandler) PreHandle(req jinx_int.IRequest) {
	log.Info(fmt.Sprintf("PreHandle req: %v...", req), "handler")
}

func (b *BaseMsgHandler) Handle(req jinx_int.IRequest) {
	log.Info("Handle...", "handler")
}

func (b *BaseMsgHandler) PostHandle(req jinx_int.IRequest) {
	log.Info("PostHandle...", "handler")
}
