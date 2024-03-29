package jinx_imp

import (
	"fmt"
	"github.com/alicia-oss/jinx/jinx_int"
	"github.com/alicia-oss/jinx/pkg/log"
)

type BaseMsgHandler struct{}

func (b *BaseMsgHandler) PreHandle(req jinx_int.IRequest) {
	log.Info(fmt.Sprintf("PreHandle req: %v...", req), "handler")
}

func (b *BaseMsgHandler) Handle(req jinx_int.IRequest) {
	req.GetConnection().Send(req.GetData(), req.GetMsgId())
}

func (b *BaseMsgHandler) PostHandle(req jinx_int.IRequest) {
}

type BaseOnCloseHandler struct{}

func (o *BaseOnCloseHandler) Handle(conn jinx_int.IConnection) {
	log.Info("BaseOnCloseHandler", "base_handler")
}

type PingHandler struct {
	jinx_int.IPingHook
}

func (o *PingHandler) Handle(req jinx_int.IRequest) {
	log.Info("PingHandler", "ping_handler")
	req.GetConnection().Ping()
}

type DefaultPingHook struct{}

func (d DefaultPingHook) PreHandle(req jinx_int.IRequest) {
}

func (d DefaultPingHook) PostHandle(req jinx_int.IRequest) {
}
