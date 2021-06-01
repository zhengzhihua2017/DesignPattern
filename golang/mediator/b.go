package mediator

import "fmt"

type b struct {
}

func (*b) Step1() {
	fmt.Println("b step 1")
}

func (*b) Step2() {
	fmt.Println("b step 2")
}

func (*b) Step3() {
	fmt.Println("b step 3")
}
