package test

import (
	"fmt"
	client_imp "github.com/alicia-oss/jinx/client/imp"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	t.Run("client", func(t *testing.T) {
		client, err := client_imp.NewClient("tcp", "127.0.0.1", 9990)
		if err != nil {
			fmt.Println(err)
			return
		}
		client.AddRoute(1, &client_imp.BaseMsgHandler{})

		client.Start()
		time.Sleep(3 * time.Second)
		err = client.Send([]byte("hello , ni hao"), 1)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Minute)
	})
}
