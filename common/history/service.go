package history

import (
	"context"
	"fmt"

	"github.com/yiminc/fxtest/common/cache"
	"github.com/yiminc/fxtest/common/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type historyService struct {
	logger          *zap.Logger
	throttledLogger *zap.Logger
	cache           cache.Cache
}

type initParams struct {
	fx.In

	Logger *zap.Logger
	ThrottledLogger *zap.Logger `name:"throttled"`

	Cache cache.Cache
}

type decoratedProvider struct {
	fx.Out

	Logger *zap.Logger
	ThrottledLogger *zap.Logger `name:"throttled"`
}

func DecorateHistory(params initParams) decoratedProvider {
	fmt.Println("decorate for history module")
	return decoratedProvider{
		Logger: params.Logger.With(zap.String("service-name", "history")),
		ThrottledLogger: params.ThrottledLogger.With(zap.String("service-name", "history")),
	}
}

var Module = fx.Module("history",
	fx.Decorate(DecorateHistory),
	fx.Provide(NewService),
	fx.Decorate(func(svc *historyService) service.Service {return svc}),
	fx.Invoke(lifecycleHooks),
)

func NewService(params initParams) *historyService {
	return &historyService{
		logger:          params.Logger,
		throttledLogger: params.ThrottledLogger,
		cache:           params.Cache,
	}
}

func (s *historyService) Start() error {
	s.logger.Info("history service start", zap.String("cache", s.cache.GetName()))
	s.throttledLogger.Info("history throttled logger")
	return nil
}

func lifecycleHooks(
	lc fx.Lifecycle,
	service *historyService,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return service.Start()
		},
	})
}
