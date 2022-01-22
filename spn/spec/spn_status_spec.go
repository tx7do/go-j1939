package spec

type DescMap map[uint8]string

type StatusSpec struct {
	bitOffset   uint8
	bitSize     uint8
	valueToDesc DescMap
}

func NewStatusSpec(bitOffset, bitSize uint8, valueToDesc DescMap) *StatusSpec {
	c := &StatusSpec{
		bitOffset:   bitOffset,
		bitSize:     bitSize,
		valueToDesc: valueToDesc,
	}
	return c
}

func (c *StatusSpec) GetBitOffset() uint8 {
	return c.bitOffset
}
func (c *StatusSpec) SetBitOffset(offset uint8) {
	c.bitOffset = offset
}

func (c *StatusSpec) GetBitSize() uint8 {
	return c.bitSize
}
func (c *StatusSpec) SetBitSize(size uint8) {
	c.bitSize = size
}

func (c *StatusSpec) GetBitMask() uint8 {
	return (0xFF >> (8 - c.bitSize)) << c.bitOffset
}

func (c *StatusSpec) SetValueDescription(value uint8, desc string) {
	c.valueToDesc[value] = desc
}
func (c *StatusSpec) GetValueDescription(value uint8) string {
	ret, ok := c.valueToDesc[value]
	if ok {
		return ret
	}
	return ""
}
func (c *StatusSpec) ClearValueDescriptions() {
	c.valueToDesc = make(DescMap)
}
func (c *StatusSpec) GetValueDescriptionsMap() DescMap {
	return c.valueToDesc
}
