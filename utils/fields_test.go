package utils

import (
	"testing"
	"time"

	opentracinglog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type stringer struct {
	string
}

func (s stringer) String() string {
	return s.string
}

func TestFieldsConversion(t *testing.T) {

	TestData := []struct {
		ZapField         zapcore.Field
		OpenTracingField opentracinglog.Field
	}{
		{
			zap.String("namespace", "123"),
			opentracinglog.String("namespace", "123"),
		},
		{
			zap.String("namespace", ""),
			opentracinglog.String("namespace", ""),
		},
		{
			zap.String("", "123"),
			opentracinglog.String("", "123"),
		},
		{
			zap.Stringer("namespace", stringer{""}),
			opentracinglog.String("namespace", ""),
		},
		{
			zap.Stringer("", stringer{"123"}),
			opentracinglog.String("", "123"),
		},
		{
			zap.Int("namespace", 1),
			opentracinglog.Int64("namespace", 1),
		},
		{
			zap.Int32("namespace", 1),
			opentracinglog.Int32("namespace", 1),
		},
		{
			zap.Int64("namespace", 1),
			opentracinglog.Int64("namespace", 1),
		},
		{
			zap.Uint32("namespace", 1),
			opentracinglog.Uint32("namespace", 1),
		},
		{
			zap.Uint64("namespace", 1),
			opentracinglog.Uint64("namespace", 1),
		},
		{
			zap.Duration("namespace", time.Second),
			opentracinglog.String("namespace", "1s"),
		},
		{
			zap.Float32("namespace", 1),
			opentracinglog.Float32("namespace", 1),
		},
		{
			zap.Float64("namespace", 1),
			opentracinglog.Float64("namespace", 1),
		},
		{
			zap.Bool("namespace", false),
			opentracinglog.Bool("namespace", false),
		},
		{
			zap.Bool("namespace", true),
			opentracinglog.Bool("namespace", true),
		},
	}

	for _, data := range TestData {
		result := ZapFieldToOpentracing(data.ZapField)
		if result.Key() != data.OpenTracingField.Key() {
			t.Errorf("Expected same key. Got %s but expected %s", result.Key(), data.OpenTracingField.Key())
		}
		if result.String() != data.OpenTracingField.String() {
			t.Errorf("Expected same string. Got %s but expected %s", result.String(), data.OpenTracingField.String())
		}
		if result.Value() != data.OpenTracingField.Value() {
			t.Errorf("Expected same value. Got %s but expected %s", result.Value(), data.OpenTracingField.Value())
		}
	}

}
