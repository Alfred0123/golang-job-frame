package domain

type (
	SampleHandler interface {
		Run()
	}
	SampleService interface {
		GetSample(id int) (string)
	}
	SampleRepository interface {
		GetSampleById(id int) (string)
	}
)