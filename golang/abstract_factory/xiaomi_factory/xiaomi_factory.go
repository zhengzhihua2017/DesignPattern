package xiaomi_factory

import (
	"abstract_factory/inc"
)

type XiaomiFactory struct {
}

func New() *XiaomiFactory {
	return &XiaomiFactory{}
}

func (*XiaomiFactory) CreateCellphone() inc.ICellphone {
	return &XiaomiCellphone{}
}

func (*XiaomiFactory) CreateIpad() inc.Ipad {
	return &XiaomiIpad{}
}

func (*XiaomiFactory) CreateSmartSoundBox() inc.ISmartSoundBox {
	return &XiaomiSmartSoundBox{}
}
