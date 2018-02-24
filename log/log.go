package log

import (
	"context"

	"github.com/opentracing-contrib/go-zap/utils"
	opentracing "github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DEBUG

// DebugWithContext logs on debug level and trace based on the context span if it exists
func DebugWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	DebugWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// DebugWithSpan logs on debug level and add the logs on the trace if span exists.
func DebugWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	Debug(log, fields...)
	logSpan(span, log, fields...)
}

// Debug logs on debug level
func Debug(log string, fields ...zapcore.Field) {
	zap.L().Debug(log, fields...)
}

// INFO

// InfoWithContext logs on info level and trace based on the context span if it exists
func InfoWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	InfoWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// InfoWithSpan logs on info level and add the logs on the trace if span exists.
func InfoWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	Info(log, fields...)
	logSpan(span, log, fields...)

}

// Info logs on info level
func Info(log string, fields ...zapcore.Field) {
	zap.L().Info(log, fields...)
}

// WARN

// WarnWithContext logs on warn level and trace based on the context span if it exists
func WarnWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	WarnWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// WarnWithSpan logs on warn level and add the logs on the trace if span exists.
func WarnWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	Warn(log, fields...)
	logSpan(span, log, fields...)

}

// Warn logs on warn level
func Warn(log string, fields ...zapcore.Field) {
	zap.L().Warn(log, fields...)
}

// ERROR

// ErrorWithContext logs on error level and trace based on the context span if it exists
func ErrorWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	ErrorWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// ErrorWithSpan logs on error level and add the logs on the trace if span exists.
func ErrorWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	Error(log, fields...)
	logSpan(span, log, fields...)
}

// Error logs on error level
func Error(log string, fields ...zapcore.Field) {
	zap.L().Error(log, fields...)
}

// DPANIC

// DPanicWithContext logs on dPanic level and trace based on the context span if it exists
func DPanicWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	DPanicWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// DPanicWithSpan logs on dPanic level and add the logs on the trace if span exists.
func DPanicWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	logSpan(span, log, fields...)
	DPanic(log, fields...)
}

// DPanic logs on dPanic level
func DPanic(log string, fields ...zapcore.Field) {
	zap.L().DPanic(log, fields...)
}

// PANIC

// PanicWithContext logs on panic level and trace based on the context span if it exists
func PanicWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	PanicWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// PanicWithSpan logs on panic level and add the logs on the trace if span exists.
func PanicWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	logSpan(span, log, fields...)
	Panic(log, fields...)
}

// Panic logs on panic level
func Panic(log string, fields ...zapcore.Field) {
	zap.L().Panic(log, fields...)
}

// FatalWithContext logs on fatal level and trace based on the context span if it exists
func FatalWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	FatalWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// FatalWithSpan logs on fatal level and add the logs on the trace if span exists.
func FatalWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	logSpan(span, log, fields...)
	Fatal(log, fields...)
}

// Fatal logs on fatal level
func Fatal(log string, fields ...zapcore.Field) {
	zap.L().Fatal(log, fields...)
}

func logSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	if span != nil {
		opentracingFields := make([]opentracinglog.Field, len(fields)+1)
		if log != "" {
			opentracingFields = append(opentracingFields, opentracinglog.String("event", log))
		}
		if len(fields) > 0 {
			opentracingFields = append(opentracingFields, utils.ZapFieldsToOpentracing(fields...)...)
		}
		span.LogFields(opentracingFields...)
	}
}
