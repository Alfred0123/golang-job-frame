package sample

import (
	domain "golang-job-frame/domain/sample"
)

type handler struct {
	svc domain.SampleService
}

func (h *handler) Run () {
	println("sample job start")
	result := h.svc.GetSample(1)
	println(result)
}