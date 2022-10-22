package jinx_int

type IServer interface {
	// Start 启动服务器方法
	Start()
	// Stop 停止服务器方法
	Stop()
	// AddRouter 为消息添加具体的处理逻辑
	AddRouter(msgId uint32, router IMsgHandle) error
}
