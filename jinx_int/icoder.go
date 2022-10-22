package jinx_int

import "io"

type ICoder interface {
	Decode(buf io.Reader) (IMessage, error)
	Encode(msg []byte, msgId uint32) ([]byte, error)
}

type IMessage interface {
	GetDataLen() uint32 //获取消息数据段长度
	GetMsgId() uint32   //获取消息ID
	GetData() []byte    //获取消息内容

	SetMsgId(uint32) //设计消息ID
	SetData([]byte)  //设计消息内容
}
