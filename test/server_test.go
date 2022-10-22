package test

import (
	"fmt"
	"jinx/jinx_imp"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("connection", func(t *testing.T) {
		server := jinx_imp.NewServer("Test", "tcp", "127.0.0.1", 9990)
		err := server.AddRouter(1, &jinx_imp.BaseMsgHandler{})
		if err != nil {
			fmt.Println("AddRouter err:", err.Error())
			return
		}
		server.Start()

		time.Sleep(5 * time.Minute)
	})

}
