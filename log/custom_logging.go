package log

import (
	"context"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func getRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if rid, ok := ctx.Value("rid").(string); ok {
		return rid
	}
	return ""
}

func getFields(ctx context.Context, msg string, v ...interface{}) *log.Entry {
	rid := getRequestID(ctx)
	pc, _, _, _ := runtime.Caller(1)
	callingFunc := runtime.FuncForPC(pc).Name()
	return log.WithFields(logrus.Fields{
		"rid": rid,
		"fc":  callingFunc,
	})
}

func Info(ctx context.Context, msg string, v ...interface{}) {
	fields := getFields(ctx, msg, v...)
	if len(v) > 0 {
		fields.Infof(msg, v...)
	} else {
		fields.Info(msg)
	}
}

func Debug(ctx context.Context, msg string, v ...interface{}) {
	fields := getFields(ctx, msg, v...)
	if len(v) > 0 {
		fields.Infof(msg, v...)
	} else {
		fields.Info(msg)
	}
}

func Error(ctx context.Context, msg string, v ...interface{}) {
	fields := getFields(ctx, msg, v...)
	if len(v) > 0 {
		fields.Infof(msg, v...)
	} else {
		fields.Info(msg)
	}
}
