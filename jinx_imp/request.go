package jinx_imp

import "github.com/alicia-oss/jinx/jinx_int"

func NewRequest(msg jinx_int.IMessage, conn jinx_int.IConnection) jinx_int.IRequest {
	return &request{
		IMessage: msg,
		conn:     conn,
	}
}

type request struct {
	jinx_int.IMessage
	conn jinx_int.IConnection
}

func (r *request) GetConnection() jinx_int.IConnection {
	return r.conn
}
