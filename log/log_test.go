package log

import (
	"context"
	"github.com/gordan0410/common/enum"
	"testing"
)

func TestNewLog(t *testing.T) {
	logger := NewLogger(LogField{ProjectName: "testProject", Version: "testVersion"})

	logger.Info().Interface("test", "test").Msg("測試")

	ctx := context.Background()
	ctx = context.WithValue(ctx, enum.TraceId.ToString(), "testTid")
	ctx = context.WithValue(ctx, enum.SubTraceId.ToString(), "testStid")

	logger.WithCtx(ctx).Info().Interface("test", "test").Msg("測試ctx")

	nilCtx := context.Background()
	logger.WithCtx(nilCtx).Info().Interface("test2", "test2").Msg("nil 測試ctx")

	WithCtx(ctx).Info().Msg("WithCtx test")
	WithCtx(ctx).Info().Interface("interface", "123").Msg("normalTest")
}
