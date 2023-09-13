package handlers

func NewHandler() Handler {
	return Handler{}
}

type Handler struct {
}

func (h Handler) Request(metric string, name string, val any) {

}
