package bulider

type DirectorA struct {
}

func NewDirectorA() *DirectorA {
	return &DirectorA{}
}

//Construct 先安装主板,再安装CPU,再安装硬盘
func (*DirectorA) Construct(builder IBuilder) {
	builder.BuildMainboard()
	builder.BuildCPU()
	builder.BuildHD()
}

type DirectorB struct {
}

func NewDirectorB() *DirectorB {
	return &DirectorB{}
}

//Construct 先安装主板,再安装硬盘,再安装CPU
func (*DirectorB) Construct(builder IBuilder) {
	builder.BuildMainboard()
	builder.BuildHD()
	builder.BuildCPU()
}
