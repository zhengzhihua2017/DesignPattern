package simple_factory

import "testing"

//TestType1 test get hiapi with factory
func TestType(t *testing.T) {
	animal1 := NewAnimal(Dog)
	animal1.Say()
	animal2 := NewAnimal(Cat)
	animal2.Say()
}
