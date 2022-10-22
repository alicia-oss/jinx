package coder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"jinx/jinx_int"
)

const headLen = 4
const msgIdLen = 4

// TlvCoder TLV协议 4byte head_len  4 byte msg_id
type TlvCoder struct {
	MaxPacketSize uint32
}

func (c *TlvCoder) Decode(buf io.Reader) (jinx_int.IMessage, error) {
	msg := &message{}
	if err := binary.Read(buf, binary.LittleEndian, &msg.dataLen); err != nil {
		return nil, err
	}

	//读msgID
	if err := binary.Read(buf, binary.LittleEndian, &msg.msgId); err != nil {
		return nil, err
	}

	//判断dataLen的长度是否超出我们允许的最大包长度
	if msg.dataLen > c.MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}
	content := make([]byte, msg.dataLen)
	if err := binary.Read(buf, binary.LittleEndian, content); err != nil {
		return nil, err
	}
	msg.content = content
	return msg, nil

}

func (c *TlvCoder) Encode(data []byte, msgId uint32) ([]byte, error) {
	msg := message{
		dataLen: uint32(len(data)),
		msgId:   msgId,
		content: data,
	}
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}
