package cmd

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// log the zap logger
	log *zap.Logger
)

// Rfc3339NanoEncoder to encode time field to RFC3339Nano format.
func Rfc3339NanoEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339Nano))
}

// InitLogger must be first to be called.
func InitLogger(debug bool) {
	var cfg zap.Config

	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.LevelKey = "lvl"
		cfg.EncoderConfig.MessageKey = "msg"
		cfg.EncoderConfig.TimeKey = "timestamp"
		cfg.EncoderConfig.EncodeTime = Rfc3339NanoEncoder
	}

	var err error
	if log, err = cfg.Build(); err != nil {
		panic("Error create logger.")
	}
	zap.ReplaceGlobals(log)
}
