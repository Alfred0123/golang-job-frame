package sample

type handler struct {
	svc service
}

func (h *handler) Run () {
	println("sample job start")
	result := h.svc.GetSample(1)
	println(result)
}