package bulider

import "fmt"

type Computer struct {
	brand    string
	describe string
}

func NewComputer(brand string) *Computer {
	cp := &Computer{
		brand:    brand,
		describe: brand,
	}
	return cp
}

func (computer *Computer) BuildCPU(cpu string) {
	computer.describe = fmt.Sprintf("%s %s", computer.describe, cpu)
}

func (computer *Computer) BuildMainboard(mainboard string) {
	computer.describe = fmt.Sprintf("%s %s", computer.describe, mainboard)
}

func (computer *Computer) BuildHD(hd string) {
	computer.describe = fmt.Sprintf("%s %s", computer.describe, hd)
}

func (computer *Computer) Show() {
	fmt.Println(computer.describe)
}
