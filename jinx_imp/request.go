package jinx_imp

import "github.com/alicia-oss/jinx/jinx_int"

func NewRequest(msg jinx_int.IMessage, conn jinx_int.IConnection) jinx_int.IRequest {
	return &request{
		IMessage: msg,
		conn:     conn,
		attrs:    make(map[string]interface{}),
	}
}

type request struct {
	jinx_int.IMessage
	conn  jinx_int.IConnection
	attrs map[string]interface{}
}

func (r *request) GetConnection() jinx_int.IConnection {
	return r.conn
}

func (r *request) SetAttr(key string, value interface{}) {
	r.attrs[key] = value
}

func (r *request) GetAttr(key string) (interface{}, bool) {
	val, ok := r.attrs[key]
	return val, ok
}
