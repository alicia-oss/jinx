package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestName(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", 12)
	ctx = context.WithValue(ctx, "ip", "192.168.0.1")

	t.Run("1", func(t *testing.T) {
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&MyFormatter{})
		logrus.Info("AAA")

	})

}
