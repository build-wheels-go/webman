package services

import (
	"context"
	"io"
	"webman/framework"
	"webman/framework/contract"
	"webman/framework/provider/log/formatter"
)

type WmLog struct {
	level     contract.LogLevel
	formatter contract.Formatter
	fielder   contract.CtxFielder
	output    io.Writer
	container framework.Container
}

// IsLevelEnable 判断该级别是否可以打印
func (log *WmLog) IsLevelEnable(level contract.LogLevel) bool {
	return level <= log.level
}

func (log *WmLog) logf(ctx context.Context, level contract.LogLevel, msg string, fields map[string]interface{}) error {
	if !log.IsLevelEnable(level) {
		return nil
	}

	//使用ctxFielder获取context中的信息
	fs := fields
	if log.fielder != nil {
		t := log.fielder(ctx)
		if t != nil {
			for k, v := range t {
				fs[k] = v
			}
		}
	}

	//判断是否绑定trace
	if log.container.IsBind(contract.TraceKey) {
		tracer := log.container.MustMake(contract.TraceKey).(contract.Trace)
		tc := tracer.GetTrace(ctx)
		m := tracer.ToMap(tc)
		if len(m) > 0 {
			for k, v := range m {
				fs[k] = v
			}
		}
	}

	// 将日志信息按照formatter序列化为字符串
	if log.formatter == nil {
		log.formatter = formatter.TextFormatter
	}

	return nil
}
