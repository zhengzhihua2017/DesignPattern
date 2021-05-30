package facade

import "fmt"

type step3 struct {
}

func newStep3() *step3 {
	return &step3{}
}

func (s3 *step3) run() {
	fmt.Println("step3 run")
}
