package facade

import "fmt"

type step2 struct {
}

func newStep2() *step2 {
	return &step2{}
}

func (s2 *step2) run() {
	fmt.Println("step2 run")
}
