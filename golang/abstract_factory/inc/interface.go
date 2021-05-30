package inc

// 抽象工厂接口,需要能够生产手机和Ipad
type IAbstractFactory interface {
	CreateCellphone() ICellphone
	CreateIpad() Ipad
	CreateSmartSoundBox() ISmartSoundBox
}

// 手机接口
type ICellphone interface {
	Call()
}

// Ipad接口
type Ipad interface {
	Play()
}

// 智能音箱接口
type ISmartSoundBox interface {
	Listen()
}
