package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	huawei := CreateFactory(Huawei)
	huaweiPhone := huawei.CreateCellphone()
	huaweiPhone.Call()
	huaweiIpad := huawei.CreateIpad()
	huaweiIpad.Play()
	huaweiSmartSoundBox := huawei.CreateSmartSoundBox()
	huaweiSmartSoundBox.Listen()

	xiaomi := CreateFactory(Xiaomi)
	xiaomiPhone := xiaomi.CreateCellphone()
	xiaomiPhone.Call()
	xiaomiIpad := xiaomi.CreateIpad()
	xiaomiIpad.Play()
	xiaomiSmartSoundBox := xiaomi.CreateSmartSoundBox()
	xiaomiSmartSoundBox.Listen()

	samsung := CreateFactory(Samsung)
	samsungPhone := samsung.CreateCellphone()
	samsungPhone.Call()
	samsungIpad := samsung.CreateIpad()
	samsungIpad.Play()
	samsungSmartSoundBox := samsung.CreateSmartSoundBox()
	samsungSmartSoundBox.Listen()
}
