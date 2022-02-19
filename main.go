package main

import (
	"context"
	"fmt"

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
		fmt.Println("decorate for test module")
		return logger.With(zap.String("test", "true"))
	}),
	fx.Decorate(func(engine log.MatchingEngine) log.MatchingEngine {
		return &matching.MatchingEngine{
			Name: "test engine",
		}
	}),
)

func main() {
	app := fx.New(
		TestModule,
		fx.NopLogger,
	)
	app.Start(context.Background())
}

//
//func NewLog() (*zap.Logger, error) {
//	return zap.NewDevelopment()
//}
//
//var LogModule = fx.Module("logger",
//	fx.Provide(NewLog),
//)
//
//type fooService struct {
//	logger *zap.Logger
//}
//
//func NewFooService(logger *zap.Logger) *fooService {
//	return &fooService{logger: logger}
//}
//
//func (s *fooService) Start() error {
//	s.logger.Info("start foo service")
//	return nil
//}
//
//var FooServiceModule = fx.Module(
//	"fooService",
//	fx.Provide(NewFooService),
//	fx.Decorate(func(logger *zap.Logger) *zap.Logger {return logger.With(zap.String("service", "foo"))}),
//	fx.Invoke(fooLifecycleHooks),
//)
//
//func fooLifecycleHooks(
//	lc fx.Lifecycle,
//	service *fooService,
//) {
//	lc.Append(fx.Hook{
//		OnStart: func(context.Context) error {
//			return service.Start()
//		},
//	})
//}
//
//var ServerModule = fx.Module(
//	"server",
//	LogModule,
//	FooServiceModule,
//)
//
//var TestModule = fx.Module(
//	"test",
//	ServerModule,
//	fx.Decorate(func(logger *zap.Logger) *zap.Logger {return logger.With(zap.String("test", "true"))}),
//)
//
//func main() {
//	app := fx.New(
//		TestModule,
//		fx.NopLogger,
//	)
//	app.Start(context.Background())
//}