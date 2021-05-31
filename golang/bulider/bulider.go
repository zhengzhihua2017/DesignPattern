package bulider

type IBuilder interface {
	BuildCPU()       // 组装cpu
	BuildMainboard() // 组装主板
	BuildHD()        // 组装硬盘
}
