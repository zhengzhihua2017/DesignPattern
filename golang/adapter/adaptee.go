package adapter

import "fmt"

//adaptee 是被适配的目标接口
type adaptee struct {
}

func newAdaptee() *adaptee {
	return &adaptee{}
}

func (p *adaptee) SpecificRequest() {
	fmt.Println("adaptee specific request")
}
