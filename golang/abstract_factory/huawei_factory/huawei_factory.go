package huawei_factory

import (
	"abstract_factory/inc"
)

type HuaweiFactory struct {
}

func New() *HuaweiFactory {
	return &HuaweiFactory{}
}

func (*HuaweiFactory) CreateCellphone() inc.ICellphone {
	return &HuaweiCellphone{}
}

func (*HuaweiFactory) CreateIpad() inc.Ipad {
	return &HuaweiIpad{}
}

func (*HuaweiFactory) CreateSmartSoundBox() inc.ISmartSoundBox {
	return &HuaweiSmartSoundBox{}
}
