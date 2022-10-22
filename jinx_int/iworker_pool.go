package jinx_int

type IWorkerPool interface {
	Submit(handle IMsgHandle, request IRequest)
	Start()
	Stop()
	GetWorkerSize() int
	GetJobChannelSize() int
}
