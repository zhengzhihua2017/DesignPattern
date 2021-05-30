package huawei_factory

import "fmt"

type HuaweiCellphone struct {
}

func (phone *HuaweiCellphone) Call() {
	fmt.Println("huawei cellphone call")
}
