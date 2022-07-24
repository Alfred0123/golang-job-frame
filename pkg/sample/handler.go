package sample

import (
	domain "golang-job-frame/domain/sample"
)

type Handler struct {
	svc domain.SampleService
}

func (h *Handler) Run () {
	println("sample job start")
	result := h.svc.GetSample(1)
	println(result)
}