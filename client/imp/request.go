package client_imp

import (
	client_int "github.com/alicia-oss/jinx/client/int"
	"github.com/alicia-oss/jinx/jinx_int"
)

func NewRequest(msg jinx_int.IMessage, conn client_int.IClient) client_int.IRequest {
	return &request{
		IMessage: msg,
		conn:     conn,
	}
}

type request struct {
	jinx_int.IMessage
	conn client_int.IClient
}

func (r *request) GetClient() client_int.IClient {
	return r.conn
}
