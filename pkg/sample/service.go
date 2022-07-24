package sample

import (
	domain "golang-job-frame/domain/sample"
)
type Service struct {
	repo domain.SampleRepository
}

func (s *Service) GetSample (id int) (string) {
	if(id > 10) {
		return ""
	} else {
		return s.repo.GetSampleById(id)
	}
}