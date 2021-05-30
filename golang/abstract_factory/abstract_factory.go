package abstract_factory

import (
	"abstract_factory/huawei_factory"
	"abstract_factory/inc"
	"abstract_factory/samsung_factory"
	"abstract_factory/xiaomi_factory"
)

type FactoryType int

const (
	Huawei FactoryType = iota
	Xiaomi
	Samsung
	Unsupported
)

// 根据给定参数创建工厂
func CreateFactory(ft FactoryType) inc.IAbstractFactory {
	switch ft {
	case Huawei:
		return huawei_factory.New()
	case Xiaomi:
		return xiaomi_factory.New()
	case Samsung:
		return samsung_factory.New()
	}
	return nil
}
