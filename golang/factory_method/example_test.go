package factory_method

import (
	"fmt"
	"testing"
)

//TestType1 test get hiapi with factory
func TestCompute(t *testing.T) {
	add := FactoryCreateOperate(OPERATE_ADD)
	add.SetA(1)
	add.SetB(2)
	fmt.Printf("add %f\n", add.Compute())

	sub := FactoryCreateOperate(OPERATE_SUB)
	sub.SetA(1)
	sub.SetB(2)
	fmt.Printf("sub %f\n", sub.Compute())

	div := FactoryCreateOperate(OPERATE_DIV)
	div.SetA(1)
	div.SetB(2)
	fmt.Printf("div %f\n", div.Compute())

	mul := FactoryCreateOperate(OPERATE_MUL)
	mul.SetA(1)
	mul.SetB(2)
	fmt.Printf("mul %f\n", mul.Compute())
}
