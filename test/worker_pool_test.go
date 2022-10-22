package test

import (
	"jinx/jinx_imp"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	t.Run("worker pool", func(t *testing.T) {
		pool := jinx_imp.NewWorkerPool(4, 5)
		pool.Start()
		time.Sleep(3 * time.Second)
		pool.Submit(&jinx_imp.BaseMsgHandler{}, nil)
		pool.Submit(&jinx_imp.BaseMsgHandler{}, nil)
		pool.Submit(&jinx_imp.BaseMsgHandler{}, nil)
		pool.Submit(&jinx_imp.BaseMsgHandler{}, nil)
		pool.Submit(&jinx_imp.BaseMsgHandler{}, nil)
		time.Sleep(4 * time.Second)
		pool.Stop()
		time.Sleep(1 * time.Second)

	})
}
