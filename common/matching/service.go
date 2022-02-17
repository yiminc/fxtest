package matching

import (
	"context"

	"github.com/yiminc/fxtest/common/cache"
	"github.com/yiminc/fxtest/common/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type matchingService struct {
	logger *zap.Logger
	cache cache.Cache
}

var Module = fx.Module("matching",
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.With(zap.String("service-name", "matching"))
	}),
	fx.Provide(NewService),
	fx.Decorate(func(dummy service.Service, svc *matchingService) service.Service {return svc}),
	fx.Invoke(lifecycleHooks),
)

func NewService(logger *zap.Logger, cache cache.Cache) *matchingService {
	return &matchingService{
		logger: logger,
		cache: cache,
	}
}

func (s *matchingService) Start() error {
	s.logger.Info("matching service start", zap.String("cache", s.cache.GetName()))
	return nil
}

func lifecycleHooks(
	lc fx.Lifecycle,
	service *matchingService,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return service.Start()
		},
	})
}
