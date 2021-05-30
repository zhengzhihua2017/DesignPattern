package facade

import "fmt"

type step1 struct {
}

func newStep1() *step1 {
	return &step1{}
}

func (s1 *step1) run() {
	fmt.Println("step1 run")
}
