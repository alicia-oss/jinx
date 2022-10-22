package jinx_int

type IRouter interface {
	Route(request IRequest) (IMsgHandle, bool)       //马上以非阻塞方式处理消息
	AddRouter(msgId uint32, router IMsgHandle) error //为消息添加具体的处理逻辑
}
