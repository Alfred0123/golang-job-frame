// +build wireinject

package sample

import (
	// domain "golang-job-frame/domain/sample"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// ProviderHandler,
	// ProviderService,
	// ProviderRepository,
	wire.Struct(new(handler), "*"),
	wire.Struct(new(service), "*"),
	wire.Struct(new(repository)),
	// wire.Bind(new(), new(*handler)),
	// wire.Bind(new(*service)),
	// wire.Bind(new(*repository)),
)

// var ProviderRepositorySet = wire.NewSet(
// 	ProviderRepository,
// 	wire.Bind(new(domain.SampleRepository), new(*repository)),
// )

func Wire() (*handler, error) {
	wire.Build(ProviderSet)
	return &handler{}, nil
}

func WireService() (*service, error) {
	wire.Build(ProviderSet)
	return &service{}, nil
}

func WireRepository() (*repository, error) {
	wire.Build(ProviderSet)
	return &repository{}, nil
}