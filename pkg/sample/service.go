package sample

import (
	domain "golang-job-frame/domain/sample"
)
type service struct {
	repo domain.SampleRepository
}

func (s *service) GetSample (id int) (string) {
	if(id > 10) {
		return ""
	} else {
		return s.repo.GetSampleById(id)
	}
}