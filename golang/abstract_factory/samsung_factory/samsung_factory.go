package samsung_factory

import "abstract_factory/inc"

type SamsungFactory struct {
}

func New() *SamsungFactory {
	return &SamsungFactory{}
}

func (*SamsungFactory) CreateCellphone() inc.ICellphone {
	return &SamsungCellphone{}
}

func (*SamsungFactory) CreateIpad() inc.Ipad {
	return &SamsungIpad{}
}

func (*SamsungFactory) CreateSmartSoundBox() inc.ISmartSoundBox {
	return &SamsungSmartSoundBox{}
}
