// +build wireinject

package wire

import (
	domain "golang-job-frame/domain/sample"

	sample "golang-job-frame/pkg/sample"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	sample.ProviderHandler,
	sample.ProviderService,
	sample.ProviderRepository,
	wire.Bind(new(domain.SampleHandler), new(*sample.Handler)),
	wire.Bind(new(domain.SampleService), new(*sample.Service)),
	wire.Bind(new(domain.SampleRepository), new(*sample.Repository)),
)

// var ProviderRepositorySet = wire.NewSet(
// 	ProviderRepository,
// 	wire.Bind(new(domain.SampleRepository), new(*repository)),
// )

func Wire() (*sample.Handler, error) {
	wire.Build(ProviderSet)
	return &sample.Handler{}, nil
}

func WireService() (*sample.Service, error) {
	wire.Build(ProviderSet)
	return &sample.Service{}, nil
}