package bulider

type SamsungBuilder struct {
	*Computer
}

func NewSamsungBuilder() *SamsungBuilder {
	sb := &SamsungBuilder{
		Computer: NewComputer("Samsung"),
	}

	return sb
}

func (builder *SamsungBuilder) BuildCPU() {
	builder.Computer.BuildCPU("Samsung-BuildCPU")
}

func (builder *SamsungBuilder) BuildMainboard() {
	builder.Computer.BuildMainboard("Samsung-BuildMainboard")
}

func (builder *SamsungBuilder) BuildHD() {
	builder.Computer.BuildHD("Samsung-BuildHD")
}

func (builder *SamsungBuilder) GetComputer() *Computer {
	return builder.Computer
}

type DellBuilder struct {
	*Computer
}

func NewDellBuilder() *DellBuilder {
	db := &DellBuilder{
		Computer: NewComputer("Dell"),
	}

	return db
}

func (builder *DellBuilder) BuildCPU() {
	builder.Computer.BuildCPU("Dell-BuildCPU")
}

func (builder *DellBuilder) BuildMainboard() {
	builder.Computer.BuildMainboard("Dell-BuildMainboard")
}

func (builder *DellBuilder) BuildHD() {
	builder.Computer.BuildHD("Dell-BuildHD")
}

func (builder *DellBuilder) GetComputer() *Computer {
	return builder.Computer
}
