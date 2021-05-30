package simple_factory

import "fmt"

//Animal is interface
type Animal interface {
	Say()
}

const (
	Dog int = 1
	Cat int = 2
)

//NewAnimal
func NewAnimal(t int) Animal {
	if t == Dog {
		return &dog{}
	} else if t == Cat {
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
