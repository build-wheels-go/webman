package trace

import (
	"context"
	"net/http"
	"time"
	"webman/framework"
	"webman/framework/contract"
	"webman/framework/gin"
)

var _ contract.Trace = (*WmTraceService)(nil)

type TraceKey string

var ContextKey = TraceKey("trace-key")

type WmTraceService struct {
	idService contract.ID

	traceIDGenerator contract.ID
	spanIDGenerator  contract.ID
}

func NewWmTraceService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	idService := container.MustMake(contract.IDKey).(contract.ID)
	return &WmTraceService{idService: idService}, nil
}

func (service *WmTraceService) WithTrace(ctx context.Context, trace *contract.TraceContext) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(string(ContextKey), trace)
		return ginCtx
	} else {
		newCtx := context.WithValue(ctx, ContextKey, trace)
		return newCtx
	}
}

func (service *WmTraceService) GetTrace(ctx context.Context) *contract.TraceContext {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		if val, ok := ginCtx.Get(string(ContextKey)); ok {
			return val.(*contract.TraceContext)
		}
	}
	if tc, ok := ctx.Value(ContextKey).(*contract.TraceContext); ok {
		return tc
	}
	return nil
}

func (service *WmTraceService) NewTrace() *contract.TraceContext {
	var traceId, spanId string
	if service.traceIDGenerator != nil {
		traceId = service.traceIDGenerator.NewID()
	} else {
		traceId = service.idService.NewID()
	}

	if service.spanIDGenerator != nil {
		spanId = service.spanIDGenerator.NewID()
	} else {
		spanId = service.idService.NewID()
	}
	return &contract.TraceContext{
		TraceId:    traceId,
		SpanId:     spanId,
		PspanId:    "",
		CspanId:    "",
		Annotation: map[string]string{},
	}
}

// ChildSpan instance a sub trace with new span id
func (service *WmTraceService) StartSpan(trace *contract.TraceContext) *contract.TraceContext {
	var cspnId string
	if service.spanIDGenerator != nil {
		cspnId = service.spanIDGenerator.NewID()
	} else {
		cspnId = service.idService.NewID()
	}
	return &contract.TraceContext{
		TraceId:    trace.TraceId,
		PspanId:    "",
		SpanId:     trace.SpanId,
		CspanId:    cspnId,
		Annotation: map[string]string{contract.TraceKeyTime: time.Now().String()},
	}
}

func (service *WmTraceService) ToMap(trace *contract.TraceContext) map[string]string {
	m := make(map[string]string)
	if trace == nil {
		return m
	}
	m[contract.TraceKeyTraceId] = trace.TraceId
	m[contract.TraceKeySpanId] = trace.SpanId
	m[contract.TraceKeyPspanId] = trace.PspanId
	m[contract.TraceKeyCspanId] = trace.CspanId

	if len(trace.Annotation) > 0 {
		for k, v := range trace.Annotation {
			m[k] = v
		}
	}

	return m
}

// get trace by http
func (service *WmTraceService) ExtractHTTP(req *http.Request) *contract.TraceContext {
	tc := &contract.TraceContext{}
	tc.TraceId = req.Header.Get(contract.TraceKeyTraceId)
	tc.PspanId = req.Header.Get(contract.TraceKeySpanId)
	tc.SpanId = req.Header.Get(contract.TraceKeyCspanId)
	tc.CspanId = ""

	if tc.TraceId == "" {
		if service.traceIDGenerator != nil {
			tc.TraceId = service.traceIDGenerator.NewID()
		} else {
			tc.TraceId = service.idService.NewID()
		}
	}

	if tc.SpanId == "" {
		if service.spanIDGenerator != nil {
			tc.SpanId = service.spanIDGenerator.NewID()
		} else {
			tc.SpanId = service.idService.NewID()
		}
	}
	return tc
}

// set trace to http
func (service *WmTraceService) InjectHTTP(req *http.Request, trace *contract.TraceContext) *http.Request {
	req.Header.Add(contract.TraceKeyTraceId, trace.TraceId)
	req.Header.Add(contract.TraceKeySpanId, trace.SpanId)
	req.Header.Add(contract.TraceKeyPspanId, trace.PspanId)
	req.Header.Add(contract.TraceKeyCspanId, trace.CspanId)

	return req
}
