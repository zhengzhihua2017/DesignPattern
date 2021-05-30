package xiaomi_factory

import "fmt"

type XiaomiCellphone struct {
}

func (phone *XiaomiCellphone) Call() {
	fmt.Println("xiaomi cellphone call")
}
