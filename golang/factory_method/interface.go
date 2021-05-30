package factory_method

type IOperation interface {
	SetA(numberA float64)
	SetB(numberB float64)
	Compute() float64
}
