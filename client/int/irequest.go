package client_int

type ICtx interface {
	GetClient() IClient //获取请求连接信息
	GetDataLen() uint32 //获取消息数据段长度
	GetMsgId() uint32   //获取消息ID
	GetData() []byte    //获取消息内容
	SetAttr(key string, value interface{})
	GetAttr(key string) (interface{}, bool)
}
