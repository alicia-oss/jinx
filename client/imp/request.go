package client_imp

import (
	client_int "github.com/alicia-oss/jinx/client/int"
	"github.com/alicia-oss/jinx/jinx_int"
)

func NewRequest(msg jinx_int.IMessage, conn client_int.IClient) client_int.ICtx {
	return &request{
		IMessage: msg,
		conn:     conn,
		attrs:    make(map[string]interface{}),
	}
}

type request struct {
	jinx_int.IMessage
	conn  client_int.IClient
	attrs map[string]interface{}
}

func (r *request) GetClient() client_int.IClient {
	return r.conn
}

func (r *request) SetAttr(key string, value interface{}) {
	r.attrs[key] = value
}
func (r *request) GetAttr(key string) (interface{}, bool) {
	val, ok := r.attrs[key]
	return val, ok
}
