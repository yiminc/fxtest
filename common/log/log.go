package log

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLog() (providedLogger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return providedLogger{}, nil
	}
	return providedLogger{
		Logger: logger,
		ThrottledLogger: logger.With(zap.String("throttled", "true")),
	}, nil
}

type providedLogger struct {
	fx.Out

	Logger *zap.Logger
	ThrottledLogger *zap.Logger `name:"throttled"`
}

var Module = fx.Options(
	fx.Provide(NewLog),
)

type MatchingEngine interface {
	Start()
}
