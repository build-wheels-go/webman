package contract

import (
	"context"
	"io"
	"time"
)

const LogKey = "wm:log"

type LogLevel uint32

const (
	// 未知的日志级别
	UnknownLevel LogLevel = iota
	// panic级别，导致程序崩溃的日志信息
	PanicLevel
	// fatal级别，导致本次请求异常终止的日志信息
	FatalLevel
	// error级别，出现错误，但是不一定影响后续程序执行
	ErrorLevel
	// warn级别，出现错误，一定不影响后续程序执行
	WarnLevel
	// info级别，正常的日志信息
	InfoLevel
	// debug级别，程序调试时的日志信息
	DebugLevel
	// tarce级别，表示最详细的日志信息，信息量大，包含调用堆栈等信息
	TraceLevel
)

// 定义从context获取信息的方法
type CtxFielder func(ctx context.Context) map[string]interface{}

// 定义输出格式的通用方法
type Formatter func(level LogLevel, time time.Time, msg string, fields map[string]interface{}) ([]byte, error)

type Log interface {
	Panic(ctx context.Context, msg string, fields map[string]interface{})

	Fatal(ctx context.Context, msg string, fields map[string]interface{})

	Error(ctx context.Context, msg string, fields map[string]interface{})

	Warn(ctx context.Context, msg string, fields map[string]interface{})

	Info(ctx context.Context, msg string, fields map[string]interface{})

	Debug(ctx context.Context, msg string, fields map[string]interface{})

	Trace(ctx context.Context, msg string, fields map[string]interface{})
	// 设置日志级别
	SetLogLevel(level LogLevel)
	// 从content获取上下文字段
	SetCtxFielder(handler CtxFielder)
	// 设置输出格式
	SetFormatter(formatter Formatter)
	// 设置输出管道
	SetOutput(output io.Writer)
}
