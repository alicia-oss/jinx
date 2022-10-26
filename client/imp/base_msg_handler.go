package client_imp

import (
	client_int "github.com/alicia-oss/jinx/client/int"
)

type BaseMsgHandler struct{}

func (b *BaseMsgHandler) PreHandle(req client_int.IRequest) {
}

func (b *BaseMsgHandler) Handle(req client_int.IRequest) {
}

func (b *BaseMsgHandler) PostHandle(req client_int.IRequest) {
}
