package factory_method

type OperationSub struct {
	*operation
}

func newOperationSub() *OperationSub {
	return &OperationSub{
		operation: newOperation(),
	}
}

func (op *OperationSub) Compute() float64 {
	return op.GetA() - op.GetB()
}
