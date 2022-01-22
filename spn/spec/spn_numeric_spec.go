package spec

type NumericSpec struct {
	formatGain   float64
	formatOffset float64
	byteSize     uint8
	units        string
}

func NewNumericSpec(formatGain, formatOffset float64, byteSize uint8, units string) *NumericSpec {
	c := &NumericSpec{
		formatGain:   formatGain,
		formatOffset: formatOffset,
		byteSize:     byteSize,
		units:        units,
	}
	return c
}

func (c *NumericSpec) GetByteSize() uint8 {
	return c.byteSize
}
func (c *NumericSpec) SetByteSize(size uint8) {
	c.byteSize = size
}

func (c *NumericSpec) GetFormatGain() float64 {
	return c.formatGain
}
func (c *NumericSpec) SetFormatGain(gain float64) {
	c.formatGain = gain
}

func (c *NumericSpec) GetFormatOffset() float64 {
	return c.formatOffset
}
func (c *NumericSpec) SetFormatOffset(offset float64) {
	c.formatOffset = offset
}

func (c *NumericSpec) GetUnits() string {
	return c.units
}
func (c *NumericSpec) SetUnits(units string) {
	c.units = units
}

func (c *NumericSpec) GetMaxValue() uint32 {
	return 0xFAFFFFFF >> (4 - c.byteSize) * 8
}

func (c *NumericSpec) FormatValue(value uint32) float64 {
	aux := float64(value)
	return aux*c.formatGain + c.formatOffset
}

func (c *NumericSpec) GetMaxFormattedValue() float64 {
	return c.FormatValue(c.GetMaxValue())
}
func (c *NumericSpec) GetMinFormattedValue() float64 {
	return c.FormatValue(0)
}
