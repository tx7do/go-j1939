package spn

import (
	"errors"
	"fmt"
	j1939 "go-j1939"
	"go-j1939/spn/spec"
)

type NumericSPN struct {
	spnImpl
	NumSpec *spec.NumericSpec
	Value   uint32
}

func NewNumericSPN(number uint32, name string, offset uint32, formatGain, formatOffset float64, byteSize uint8, units string) *NumericSPN {
	c := &NumericSPN{}
	c.spnImpl.Spec = spec.NewSpec(number, name, offset)
	c.NumSpec = spec.NewNumericSpec(formatGain, formatOffset, byteSize, units)
	c.Value = 0xFFFFFFFF
	//c.Owner = nil
	return c
}

func (c *NumericSPN) GetType() EType {
	return NumericType
}

func (c *NumericSPN) GetFormattedValue() float64 {
	aux := float64(c.Value)

	//Apply gain and offset
	return aux*c.GetFormatGain() + c.GetFormatOffset()
}

func (c *NumericSPN) SetFormattedValue(value float64) {
	aux := (value - c.GetFormatOffset()) / c.GetFormatGain()

	threshold := float64(((uint64)(1)) << (c.GetByteSize() * 8))

	if aux >= 0 && (aux < threshold) {
		c.Value = uint32(aux)
		return
	}
}

func (c *NumericSPN) GetByteSize() uint8 {
	return c.NumSpec.GetByteSize()
}
func (c *NumericSPN) GetFormatGain() float64 {
	return c.NumSpec.GetFormatGain()
}
func (c *NumericSPN) GetFormatOffset() float64 {
	return c.NumSpec.GetFormatOffset()
}
func (c *NumericSPN) GetUnits() string {
	return c.NumSpec.GetUnits()
}
func (c *NumericSPN) GetNumericSpec() *spec.NumericSpec {
	return c.NumSpec
}

func (c *NumericSPN) ToString() string {
	retVal := c.spnImpl.ToString()
	return fmt.Sprintf("%s -> Value: %.2f %s\n",
		retVal, c.GetFormattedValue(), c.GetUnits())
}

// Decode 解码
func (c *NumericSPN) Decode(buffer []byte) error {
	length := len(buffer)
	if int(c.GetByteSize()) > length || c.GetByteSize() > j1939.SpnNumericMaxByteSize {
		return errors.New("[NumericSPN::Decode] Spn length is bigger than expected")
	}
	c.Value = 0
	for i := 0; i < int(c.GetByteSize()); i++ {
		c.Value |= uint32(buffer[i]) << (i * 8)
	}

	return nil
}

// Encode 编码
func (c *NumericSPN) Encode(buffer []byte) error {
	if c.GetByteSize() > j1939.SpnNumericMaxByteSize {
		return errors.New("[NumericSPN::Encode] Spn length is bigger than expected")
	}

	for i := 0; i < int(c.GetByteSize()); i++ {
		buffer[i] = byte(c.Value>>(i*8)) & 0xFF
	}

	return nil
}
