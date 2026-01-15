package enum

type TraceEnum string

const (
	TraceId    TraceEnum = "x-trace-id"
	SubTraceId TraceEnum = "x-sub-trace-id"
)

func (e TraceEnum) ToString() string {
	return string(e)
}
