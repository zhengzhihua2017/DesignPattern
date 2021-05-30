package adapter

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	*adaptee
}

//NewAdapter 是Adapter的工厂函数
func NewAdapter() ITarget {
	return &adapter{
		adaptee: newAdaptee(),
	}
}

func (p *adapter) Request() {
	p.adaptee.SpecificRequest()
}
