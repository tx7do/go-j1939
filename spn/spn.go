package spn

import (
	"fmt"
	"go-j1939/spn/spec"
)

type EType int32

const (
	// NumericType 数字类型
	NumericType EType = 0

	// StatusType 状态类型
	StatusType EType = 1

	// StringType 字符串类型
	StringType EType = 2
)

type SPN interface {
	// GetType 获取spn的类型
	GetType() EType

	// GetSpnNumber 获取spn的编号
	GetSpnNumber() uint32

	// GetName 获取spn的名字
	GetName() string

	GetSpec() *spec.Spec

	GetByteSize() uint8

	GetOffset() uint32
	SetOffset(offset uint32)

	// Decode 解码
	Decode(buffer []byte) error
	// Encode 编码
	Encode(buffer []byte) error

	// ToString 转换为字符串
	ToString() string

	//SetOwner(owner *generic_frame.GenericFrame)

	GetFormattedValue() float64
	SetFormattedValue(value float64)

	SetStatusValue(value uint8)
	GetStatusValue() uint8

	SetStringValue(value string)
	GetStringValue() string
}

type spnImpl struct {
	SPN
	Spec *spec.Spec
	//Owner *generic_frame.GenericFrame
}

func (c *spnImpl) GetOffset() uint32 {
	return c.Spec.GetOffset()
}
func (c *spnImpl) SetOffset(_ uint32) {
}

func (c *spnImpl) GetSpnNumber() uint32 {
	return c.Spec.GetSpnNumber()
}
func (c *spnImpl) GetName() string {
	return c.Spec.GetName()
}

func (c *spnImpl) GetSpec() *spec.Spec {
	return c.Spec
}

func (c *spnImpl) ToString() string {
	return fmt.Sprintf("SPN %d : %s",
		c.GetSpnNumber(), c.GetName())
}

func (c *spnImpl) GetByteSize() uint8 {
	return 0
}

//func (c *spnImpl) SetOwner(owner *generic_frame.GenericFrame) {
//	c.Owner = owner
//}

func (c *spnImpl) GetFormattedValue() float64 {
	return 0
}
func (c *spnImpl) SetFormattedValue(_ float64) {

}

func (c *spnImpl) SetStatusValue(_ uint8) {

}
func (c *spnImpl) GetStatusValue() uint8 {
	return 0
}

func (c *spnImpl) SetStringValue(_ string) {

}
func (c *spnImpl) GetStringValue() string {
	return ""
}
