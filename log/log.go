package log

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func Infof(format string, v ...interface{}) {
	g.Log().Skip(2).Line().Infof(nil, format, v...)
}

func InfoContextf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Skip(1).Line().Infof(ctx, format, v...)
}

func Debugf(format string, v ...interface{}) {
	g.Log().Skip(1).Line().Debugf(nil, format, v...)
}

func DebugContextf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Skip(2).Line().Debugf(ctx, format, v...)
}

func Errorf(format string, v ...interface{}) {
	g.Log().Skip(1).Line().Errorf(nil, format, v...)
}

func ErrorContextf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Skip(2).Line().Errorf(ctx, format, v...)
}
