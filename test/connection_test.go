package test

import (
	"fmt"
	"jinx/jinx_imp"
	"jinx/pkg/coder"
	"net"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	t.Run("connection", func(t *testing.T) {
		server := jinx_imp.NewServer("Test", "tcp", "127.0.0.1", 9990)
		err := server.AddRouter(1, &jinx_imp.BaseMsgHandler{})
		if err != nil {
			fmt.Println("AddRouter err:", err.Error())
			return
		}
		server.Start()
		d, err := net.Dial("tcp", "127.0.0.1:9990")
		if err != nil {
			fmt.Println("dial err:", err.Error())
		}
		time.Sleep(5 * time.Second)
		tlvCoder := coder.TlvCoder{MaxPacketSize: 512}
		data := "hello world yes  you are right"
		bytesss := []byte(data)
		encode, err := tlvCoder.Encode(bytesss, 1)
		_, err = d.Write(encode)
		if err != nil {
			fmt.Println("Write err:", err.Error())
		}
		time.Sleep(2 * time.Second)
		d.Close()
		time.Sleep(2 * time.Second)
		server.Stop()
	})
}
