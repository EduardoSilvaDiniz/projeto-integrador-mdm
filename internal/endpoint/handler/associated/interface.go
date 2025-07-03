package associated

type Handlers struct {
	AssociatedController AssociatedController
}

func NewHandlers(assoc AssociatedController) *Handlers {
	return &Handlers{AssociatedController: assoc}
}
