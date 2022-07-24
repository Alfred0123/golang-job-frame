package sample

import (
	// domain "golang-job-frame/domain/sample"
	"sync"
)

var (
	hdl *handler
	hdlOnce sync.Once
	
	svc *service
	svcOnce sync.Once

	repo *repository
	repoOnce sync.Once
)

func ProviderHandler(svc service) (*handler) {
	hdlOnce.Do(func() {
		hdl = &handler {
			svc: svc,
		}
	})
	return hdl
}

func ProviderService(repo repository) (*service) {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})
	return svc
}

func ProviderRepository() (*repository) {
	repoOnce.Do(func() {
		repo = &repository{}
	})
	return repo
}