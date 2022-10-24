package client_int

import (
	"net"
)

type IClient interface {
	Start()
	Close()
	AddRoute(msg uint32, handle IMsgHandle) error
	// GetTCPConnection 获取当前链接的套接字
	GetTCPConnection() net.Conn
	// GetRemoteAddr 获取服务器的TCP状态 IP Port
	GetRemoteAddr() net.Addr
	// Send 发送数据，将数据发送给client
	Send(data []byte, msgId uint32) error
	// SetAttr 给业务系统提供扩展属性
	SetAttr(key string, value interface{})
	GetAttr(key string) (interface{}, bool)
	DeleteAttr(key string)
}
