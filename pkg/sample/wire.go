// +build wireinject

package sample

// import (
// 	domain "golang-job-frame/domain/sample"

// 	"github.com/google/wire"
// )

// var ProviderSet = wire.NewSet(
// 	ProviderHandler,
// 	ProviderService,
// 	ProviderRepository,
// 	wire.Bind(new(domain.SampleHandler), new(*handler)),
// 	wire.Bind(new(domain.SampleService), new(*service)),
// 	wire.Bind(new(domain.SampleRepository), new(*repository)),
// )

// // var ProviderRepositorySet = wire.NewSet(
// // 	ProviderRepository,
// // 	wire.Bind(new(domain.SampleRepository), new(*repository)),
// // )

// func Wire() (*handler, error) {
// 	wire.Build(ProviderSet)
// 	return &handler{}, nil
// }

// func WireService() (*service, error) {
// 	wire.Build(ProviderSet)
// 	return &service{}, nil
// }