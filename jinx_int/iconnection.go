package jinx_int

import "net"

type IConnection interface {
	// Start 启动链接 ，让当前链接开始工作
	Start()
	// Stop 停止链接， 结束当前链接
	Stop()
	// GetTCPConnection 获取当前链接的套接字
	GetTCPConnection() net.Conn
	// GetConnID 获取当前链接的链接ID
	GetConnID() uint32
	// GetRemoteAddr 获取远程客户端的TCP状态 IP Port
	GetRemoteAddr() net.Addr
	// Send 发送数据，将数据发送给client
	Send(data []byte, msgId uint32) error
}
