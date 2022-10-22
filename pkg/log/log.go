package log

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
)

func init() {
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&MyFormatter{})
	logrus.AddHook(NewContextHook(logrus.ErrorLevel, logrus.WarnLevel, logrus.InfoLevel))
}

var Log = logger{}

type logger struct {
}

func (l logger) Info(ctx context.Context, msg string) {
	logrus.WithFields(logrus.Fields{
		"user_id": ctx.Value("user_id"),
		"ip":      ctx.Value("ip"),
	}).Info(msg)
}

func (l logger) Error(ctx context.Context, msg string) {
	logrus.WithFields(logrus.Fields{
		"user_id": ctx.Value("user_id"),
		"ip":      ctx.Value("ip"),
	}).Error(msg)
}

type MyFormatter struct{}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fields := entry.Data
		fStr := make([]string, len(fields))
		i := 0
		for key, value := range fields {
			fStr[i] = fmt.Sprintf("[%s : %s]", key, value)
			i++
		}
		s := strings.Join(fStr, " ")
		newLog = fmt.Sprintf("[%s] [%s] %s %s\n",
			timestamp, entry.Level, s, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func Info(msg, modName string) {
	logrus.WithFields(logrus.Fields{
		Module: modName,
	}).Info(msg)
}

func Error(msg, modName string) {
	logrus.WithFields(logrus.Fields{
		Module: modName,
	}).Info(msg)
}

func InfoWithCtx(ctx context.Context, msg string) {
	logrus.WithFields(logrus.Fields{
		"user_id": ctx.Value("user_id"),
		"ip":      ctx.Value("ip"),
	}).Error(msg)
}
