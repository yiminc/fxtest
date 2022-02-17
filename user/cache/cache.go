package cache

import (
	"github.com/yiminc/fxtest/common/cache"
	"go.uber.org/fx"
)

type (
	cacheImpl struct {}
)

var Module = fx.Options(
	fx.Provide(NewCache),
	fx.Decorate(func(cc cache.Cache, uc *cacheImpl) cache.Cache {
		return uc
	}),
)

func NewCache() *cacheImpl {
	return &cacheImpl{}
}

func (s *cacheImpl) GetName() string {
	return "user cache"
}
