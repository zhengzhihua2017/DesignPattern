package simple_factory

import "testing"

//TestType1 test get hiapi with factory
func TestType(t *testing.T) {
	animal1 := NewAnimal(1)
	animal1.Say()
	animal2 := NewAnimal(2)
	animal2.Say()
}
