package sample

type service struct {
	repo repository
}

func (s *service) GetSample (id int) (string) {
	if(id > 10) {
		return ""
	} else {
		return repo.GetSampleById(id)
	}
}