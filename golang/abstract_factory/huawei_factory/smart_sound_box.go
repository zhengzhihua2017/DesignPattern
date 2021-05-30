package huawei_factory

import "fmt"

type HuaweiSmartSoundBox struct {
}

func (box *HuaweiSmartSoundBox) Listen() {
	fmt.Println("huawei smart sound box listen")
}
