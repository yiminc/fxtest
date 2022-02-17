package cache

import "go.uber.org/fx"

type (
	Cache interface {
		GetName() string
	}

	cacheImpl struct {}
)

var Module = fx.Options(
	fx.Provide(NewCache),
	fx.Provide(func(impl *cacheImpl) Cache {return impl}),
)

func NewCache() *cacheImpl {
	return &cacheImpl{}
}

func (s *cacheImpl) GetName() string {
	return "common cache"
}
