package utils

import (
	"math"

	opentracinglog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap/zapcore"
)

// ZapFieldsToOpentracing returns a table of standard opentracing field based on
// the inputed table of Zap field.
func ZapFieldsToOpentracing(zapFields ...zapcore.Field) []opentracinglog.Field {
	opentracingFields := make([]opentracinglog.Field, len(zapFields))

	for i, zapField := range zapFields {
		opentracingFields[i] = ZapFieldToOpentracing(zapField)
	}

	return opentracingFields
}

// ZapFieldToOpentracing returns a standard opentracing field based on the
// input Zap field.
func ZapFieldToOpentracing(zapField zapcore.Field) opentracinglog.Field {
	switch zapField.Type {

	case zapcore.BoolType:
		val := false
		if zapField.Integer >= 1 {
			val = true
		}
		return opentracinglog.Bool(zapField.Key, val)
	case zapcore.Float32Type:
		return opentracinglog.Float32(zapField.Key, math.Float32frombits(uint32(zapField.Integer)))
	case zapcore.Float64Type:
		return opentracinglog.Float64(zapField.Key, math.Float64frombits(uint64(zapField.Integer)))
	case zapcore.Int64Type:
		return opentracinglog.Int64(zapField.Key, int64(zapField.Integer))
	case zapcore.Int32Type:
		return opentracinglog.Int32(zapField.Key, int32(zapField.Integer))
	case zapcore.StringType:
		return opentracinglog.String(zapField.Key, zapField.String)
	case zapcore.Uint64Type:
		return opentracinglog.Uint64(zapField.Key, uint64(zapField.Integer))
	case zapcore.Uint32Type:
		return opentracinglog.Uint32(zapField.Key, uint32(zapField.Integer))
	case zapcore.ErrorType:
		return opentracinglog.Error(zapField.Interface.(error))
	default:
		return opentracinglog.Object(zapField.Key, zapField.Interface)
	}
}
