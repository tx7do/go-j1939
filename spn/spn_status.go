package spn

import (
	"errors"
	"fmt"
	"go-j1939/spn/spec"
)

type StatusSPN struct {
	spnImpl
	StatSpec *spec.StatusSpec
	Value    uint8
}

func NewStatusSPN(number uint32, name string, offset uint32, bitOffset, bitSize uint8, valueToDesc spec.DescMap) *StatusSPN {
	c := StatusSPN{}
	c.spnImpl.Spec = spec.NewSpec(number, name, offset)
	c.StatSpec = spec.NewStatusSpec(bitOffset, bitSize, valueToDesc)
	c.Value = 0xFF >> (8 - c.GetBitSize())
	//c.Owner = nil
	return &c
}

func (c *StatusSPN) GetType() EType {
	return StatusType
}

func (c *StatusSPN) SetStatusValue(value uint8) {

	if value < (1 << c.GetBitSize()) {
		c.Value = value
		return
	}
}

func (c *StatusSPN) GetStatusValue() uint8 {
	return c.Value
}

func (c *StatusSPN) GetBitOffset() uint8 {
	return c.StatSpec.GetBitOffset()
}

func (c *StatusSPN) GetBitSize() uint8 {
	return c.StatSpec.GetBitSize()
}

func (c *StatusSPN) GetByteSize() uint8 {
	return 1
}

func (c *StatusSPN) ToString() string {
	retVal := c.spnImpl.ToString()
	return fmt.Sprintf("%s -> Status: %s(%d)\n",
		retVal, c.GetValueDescription(c.Value), c.Value)
}

func (c *StatusSPN) GetValueDescription(value uint8) string {
	return c.StatSpec.GetValueDescription(value)
}

func (c *StatusSPN) GetValueDescriptionsMap() spec.DescMap {
	return c.StatSpec.GetValueDescriptionsMap()
}

func (c *StatusSPN) GetStatusSpec() *spec.StatusSpec {
	return c.StatSpec
}

// Decode 解码
func (c *StatusSPN) Decode(buffer []byte) error {
	if c.GetBitOffset() > 7 || c.GetBitSize() > 8 || c.GetBitOffset()+c.GetBitSize() > 8 {
		return errors.New("[StatusSPN::Decode] Format incorrect to decode properly this spnImpl")
	}

	mask := byte(0xFF >> (8 - c.GetBitSize()))
	c.Value = (buffer[0] >> c.GetBitOffset()) & mask

	return nil
}

// Encode 编码
func (c *StatusSPN) Encode(buffer []byte) error {
	if c.GetBitOffset() > 7 || c.GetBitSize() > 8 || c.GetBitOffset()+c.GetBitSize() > 8 {
		return errors.New("[StatusSPN::Encode] Format incorrect to encode properly this spnImpl")
	}

	mask := uint8((0xFF >> (8 - c.GetBitSize())) << c.GetBitOffset())
	value := c.Value << c.GetBitOffset()

	if (value & mask) != value {
		return errors.New("[StatusSPN::Encode] Value to encode is bigger than expected")
	}

	//Clear the bits from the buffer
	buffer[0] = (buffer[0]) & (^mask)

	//Set the new value
	buffer[0] = (buffer[0]) | (value)

	return nil
}
