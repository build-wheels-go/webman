package contract

import (
	"context"
	"net/http"
)

const TraceKey = "wm:trace"

const (
	TraceKeyTraceId = "trace_id"
	TraceKeySpanId  = "span_id"
	TraceKeyCspanId = "cspan_id"
	TraceKeyPspanId = "pspan_id"
	TraceKeyCaller  = "caller"
	TraceKeyMethod  = "method"
	TraceKeyTime    = "time"
)

type TraceContext struct {
	TraceId string //traceId,全局唯一
	SpanId  string //当前系统spanId
	PspanId string //父级spanId
	CspanId string //子节点spanId

	Annotation map[string]string //标记各种信息
}

type Trace interface {
	WithTrace(ctx context.Context, trace *TraceContext) context.Context
	GetTrace(ctx context.Context) *TraceContext
	NewTrace() *TraceContext
	StartSpan(trace *TraceContext) *TraceContext
	ToMap(trace *TraceContext) map[string]string
	ExtractHTTP(req *http.Request) *TraceContext
	InjectHTTP(req *http.Request, trace *TraceContext) *http.Request
}
