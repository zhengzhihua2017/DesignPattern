package factory_method

type OperationDiv struct {
	*operation
}

func newOperationDiv() *OperationDiv {
	return &OperationDiv{
		operation: newOperation(),
	}
}

func (op *OperationDiv) Compute() float64 {
	return op.GetA() / op.GetB()
}
