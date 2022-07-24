package sample

import (
	domain "golang-job-frame/domain/sample"
	"sync"
)

var (
	hdl *Handler
	hdlOnce sync.Once
	
	svc *Service
	svcOnce sync.Once

	repo *Repository
	repoOnce sync.Once
)

func ProviderHandler(svc domain.SampleService) (*Handler) {
	hdlOnce.Do(func() {
		hdl = &Handler {
			svc: svc,
		}
	})
	return hdl
}

func ProviderService(repo domain.SampleRepository) (*Service) {
	svcOnce.Do(func() {
		svc = &Service{
			repo: repo,
		}
	})
	return svc
}

func ProviderRepository() (*Repository) {
	repoOnce.Do(func() {
		repo = &Repository{}
	})
	return repo
}