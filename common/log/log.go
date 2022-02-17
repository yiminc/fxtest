package log

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLog() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

var Module = fx.Options(
	fx.Provide(NewLog),
)