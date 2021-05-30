package xiaomi_factory

import "fmt"

type XiaomiSmartSoundBox struct {
}

func (box *XiaomiSmartSoundBox) Listen() {
	fmt.Println("xiaomi smart sound box listen")
}
