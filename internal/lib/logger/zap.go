package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func NewLogger() (*zap.SugaredLogger, error) {
	var err error
	cfgLogger := zap.NewProductionConfig()
	cfgLogger.DisableStacktrace = true
	cfgLogger.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	prodLogger, err := cfgLogger.Build()
	if err != nil {
		return nil, err
	}

	return prodLogger.Sugar(), nil
}
