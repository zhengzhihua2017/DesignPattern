package factory_method

type OperationAdd struct {
	*operation
}

func newOperationAdd() *OperationAdd {
	return &OperationAdd{
		operation: newOperation(),
	}
}

func (op *OperationAdd) Compute() float64 {
	return op.GetA() + op.GetB()
}
