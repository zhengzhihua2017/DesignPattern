package bulider

import "testing"

func TestDirector(t *testing.T) {
	da1 := NewDirectorA()
	dasb1 := NewSamsungBuilder()
	da1.Construct(dasb1)
	dasbComputer1 := dasb1.GetComputer()
	dasbComputer1.Show()

	da2 := NewDirectorA()
	dasdb2 := NewDellBuilder()
	da2.Construct(dasdb2)
	dasbComputer2 := dasdb2.GetComputer()
	dasbComputer2.Show()

	da3 := NewDirectorB()
	dasb3 := NewSamsungBuilder()
	da3.Construct(dasb3)
	dasbComputer3 := dasb3.GetComputer()
	dasbComputer3.Show()

	da4 := NewDirectorB()
	dasdb4 := NewDellBuilder()
	da4.Construct(dasdb4)
	dasbComputer4 := dasdb4.GetComputer()
	dasbComputer4.Show()
}
