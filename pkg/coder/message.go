package coder

type message struct {
	dataLen uint32
	msgId   uint32
	content []byte
}

func (m *message) GetDataLen() uint32 {
	return m.dataLen
}

func (m *message) GetMsgId() uint32 {
	return m.msgId
}

func (m *message) GetData() []byte {
	return m.content
}

func (m *message) SetMsgId(u uint32) {
	m.msgId = u
}

func (m *message) SetData(bytes []byte) {
	m.content = bytes
	m.dataLen = uint32(len(bytes))
}
