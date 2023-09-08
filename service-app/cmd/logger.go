package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func setupLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123Z)
	config.EncoderConfig.EncodeDuration = customDurationEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil

}

func customDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	if d >= time.Second {
		sec := d.Seconds()
		enc.AppendFloat64(sec)
		enc.AppendString("s")
	} else if d >= time.Millisecond {
		milliseconds := d.Milliseconds()
		enc.AppendInt64(milliseconds)
		enc.AppendString("ms")
	} else {
		microseconds := d.Microseconds()
		enc.AppendInt64(microseconds)
		enc.AppendString("μs")
	}
}
