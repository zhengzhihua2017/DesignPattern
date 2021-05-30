package samsung_factory

import "fmt"

type SamsungSmartSoundBox struct {
}

func (box *SamsungSmartSoundBox) Listen() {
	fmt.Println("samsung smart sound box listen")
}
