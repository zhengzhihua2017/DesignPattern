package factory_method

type OperateType int

const (
	OPERATE_ADD OperateType = iota
	OPERATE_SUB
	OPERATE_DIV
	OPERATE_MUL
)

func FactoryCreateOperate(op OperateType) IOperation {
	switch op {
	case OPERATE_ADD:
		return newOperationAdd()
	case OPERATE_SUB:
		return newOperationSub()
	case OPERATE_DIV:
		return newOperationDiv()
	case OPERATE_MUL:
		return newOperationMul()
	}
	return nil
}
