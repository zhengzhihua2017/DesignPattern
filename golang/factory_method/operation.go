package factory_method

type operation struct {
	numberA float64
	numberB float64
}

func newOperation() *operation {
	return &operation{}
}

func (op *operation) SetA(numberA float64) {
	op.numberA = numberA
}

func (op *operation) SetB(numberB float64) {
	op.numberB = numberB
}

func (op *operation) GetA() float64 {
	return op.numberA
}

func (op *operation) GetB() float64 {
	return op.numberB
}

func (op *operation) Compute() float64 {
	return 0
}
