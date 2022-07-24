package sample

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockRepo struct{
  mock.Mock
}

func (m *MockRepo) GetSampleById (id int) (string) {
	println("mock GetSampleById")
	return "sample"
}

func TestGetSample (t *testing.T) {
	// svc := service{ repo: repository{} }
	mock := MockRepo{}
	svc := Service{
		repo: &mock,
	}
	
	if svc.GetSample(1) != "sample" {
		t.Error("expect str is sample")
	}
}