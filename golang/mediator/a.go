package mediator

import "fmt"

type a struct {
}

func (*a) Step1() {
	fmt.Println("a step 1")
}

func (*a) Step2() {
	fmt.Println("a step 2")
}

func (*a) Step3() {
	fmt.Println("a step 3")
}
