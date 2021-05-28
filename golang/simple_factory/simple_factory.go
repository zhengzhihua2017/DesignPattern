package simple_factory

import "fmt"

//Animal is interface
type Animal interface {
	Say()
}

//NewAnimal
func NewAnimal(t int) Animal {
	if t == 1 {
		return &dog{}
	} else if t == 2 {
		return &cat{}
	}
	return nil
}

type dog struct{}

func (*dog) Say() {
	fmt.Println("汪 ~ 汪 ~")
}

type cat struct{}

func (*cat) Say() {
	fmt.Println("喵 ~ 喵 ~")
}
