package samsung_factory

import "fmt"

type SamsungCellphone struct {
}

func (phone *SamsungCellphone) Call() {
	fmt.Println("samsung cellphone call")
}
