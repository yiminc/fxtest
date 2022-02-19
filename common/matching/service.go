package matching

import (
	"context"
	"fmt"

	"github.com/yiminc/fxtest/common/cache"
	"github.com/yiminc/fxtest/common/log"
	"github.com/yiminc/fxtest/common/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type matchingService struct {
	logger *zap.Logger
	cache cache.Cache
	engine log.MatchingEngine
}

var Module = fx.Module("matching",
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		fmt.Println("decorate for matching module")
		return logger.With(zap.String("service-name", "matching"))
	}),
	fx.Provide(NewService),
	fx.Decorate(func(svc *matchingService) service.Service {return svc}),
	fx.Provide(NewMatchingEngine),
	fx.Invoke(lifecycleHooks),
)

func NewMatchingEngine() log.MatchingEngine {
	return &MatchingEngine{
		Name: "default matching engine",
	}
}

type MatchingEngine struct {
	Name string
}

func (m *MatchingEngine) Start() {
	fmt.Printf("start %v\n", m.Name)
}

func NewService(logger *zap.Logger, cache cache.Cache, engine log.MatchingEngine) *matchingService {
	return &matchingService{
		logger: logger,
		cache: cache,
		engine: engine,
	}
}

func (s *matchingService) Start() error {
	s.logger.Info("matching service start", zap.String("cache", s.cache.GetName()))
	s.engine.Start()
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
