package factory_method

type OperationMul struct {
	*operation
}

func newOperationMul() *OperationMul {
	return &OperationMul{
		operation: newOperation(),
	}
}

func (op *OperationMul) Compute() float64 {
	return op.GetA() * op.GetB()
}
