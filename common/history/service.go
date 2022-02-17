package history

import (
	"context"

	"github.com/yiminc/fxtest/common/cache"
	"github.com/yiminc/fxtest/common/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type historyService struct {
	logger *zap.Logger
	cache cache.Cache
}

var Module = fx.Module("history",
	//fx.Decorate(func(logger *zap.Logger) *zap.Logger {
	//	return logger.With(zap.String("service-name", "history"))
	//}),
	fx.Provide(NewService),
	fx.Decorate(func(dummy service.Service, svc *historyService) service.Service {return svc}),
	fx.Invoke(lifecycleHooks),
)

func NewService(logger *zap.Logger, cache cache.Cache) *historyService {
	return &historyService{
		logger: logger,
		cache: cache,
	}
}

func (s *historyService) Start() error {
	s.logger.Info("history service start", zap.String("cache", s.cache.GetName()))
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
