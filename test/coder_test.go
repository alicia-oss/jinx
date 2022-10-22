package test

import (
	"bytes"
	"fmt"
	"jinx/pkg/coder"
	"testing"
)

func TestCoder(t *testing.T) {
	t.Run("coder", func(t *testing.T) {
		tlvCoder := coder.TlvCoder{MaxPacketSize: 512}
		data := "hello world yes  you are right"
		bytesss := []byte(data)
		encode, err := tlvCoder.Encode(bytesss, 1)
		data2 := "no no no, my wrong"
		bytesss = []byte(data2)
		encode2, _ := tlvCoder.Encode(bytesss, 2)
		encode = append(encode, encode2...)
		if err != nil {
			fmt.Println("encoder err:", err)
		}
		reader := bytes.NewReader(encode)
		message, err := tlvCoder.Decode(reader)
		if err != nil {
			fmt.Println("Decode err:", err)
		}
		fmt.Printf("%v \n", message)

		message, err = tlvCoder.Decode(reader)
		if err != nil {
			fmt.Println("Decode err:", err)
		}
		fmt.Printf("%v", message)
	})
}
