package jinx_int

type IMsgHandle interface {
	PreHandle(request IRequest)  //在处理conn业务之前的钩子方法
	Handle(request IRequest)     //处理conn业务的方法
	PostHandle(request IRequest) //处理conn业务之后的钩子方法
}

type IOnCloseHandle interface {
	Handle(conn IConnection)
}

type IPingHook interface {
	PreHandle(req IRequest)
	PostHandle(req IRequest)
}
