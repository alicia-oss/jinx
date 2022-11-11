package client_int

type IMsgHandle interface {
	PreHandle(request ICtx)  //在处理conn业务之前的钩子方法
	Handle(request ICtx)     //处理conn业务的方法
	PostHandle(request ICtx) //处理conn业务之后的钩子方法
}
