package mediator

type Mediator struct {
	a *a
	b *b
}

func NewMediator() *Mediator {
	return &Mediator{
		a: &a{},
		b: &b{},
	}
}

func (m *Mediator) Run(i int) {
	switch i {
	case 1:
		m.a.Step1()
		m.b.Step2()
		m.a.Step3()
	case 2:
		m.b.Step1()
		m.a.Step2()
		m.b.Step3()
	}
}
