package client_int

type IRouter interface {
	Route(request ICtx) (IMsgHandle, bool)
	AddRouter(msgId uint32, router IMsgHandle) error //为消息添加具体的处理逻辑
}
