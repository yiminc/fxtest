package main

import (
	"context"

	"github.com/yiminc/fxtest/common/cache"
	"github.com/yiminc/fxtest/common/history"
	"github.com/yiminc/fxtest/common/log"
	"github.com/yiminc/fxtest/common/matching"
	usercache "github.com/yiminc/fxtest/user/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var ServiceModule = fx.Module("server",
	log.Module,
	cache.Module,
	usercache.Module,
	history.Module,
	matching.Module,
)

var TestModule = fx.Options(
	ServiceModule,
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.With(zap.String("test", "true"))
	}),
)

func main() {
	app := fx.New(
		TestModule,
		fx.NopLogger,
	)
	app.Start(context.Background())
}
