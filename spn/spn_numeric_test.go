package spn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var wheelSpeed *NumericSPN = nil

func init() {
	wheelSpeed = NewNumericSPN(84, "Wheel Speed", 1, 0.00390625, 0, 2, "km/h")
}

func TestNumericSpnBasic(t *testing.T) {
	numeric := NewNumericSPN(100, "test_numeric", 3, 2.5, -2, 2, "%")

	assert.Equal(t, numeric.GetSpnNumber(), uint32(100))
	assert.Equal(t, numeric.GetName(), "test_numeric")
	assert.Equal(t, numeric.GetOffset(), uint32(3))
	assert.Equal(t, numeric.GetFormatGain(), 2.5)
	assert.Equal(t, numeric.GetFormatOffset(), float64(-2))
	assert.Equal(t, numeric.GetByteSize(), uint8(2))
	assert.Equal(t, numeric.GetUnits(), "%")
	assert.Equal(t, numeric.GetType(), NumericType)
}

func TestNumericSpnSetValue(t *testing.T) {
	c := wheelSpeed

	c.SetFormattedValue(50)

	assert.Equal(t, c.GetFormattedValue(), float64(50))
	assert.Equal(t, c.GetByteSize(), uint8(2))
	assert.Equal(t, c.GetFormatGain(), 0.00390625)
	assert.Equal(t, c.GetFormatOffset(), float64(0))
	assert.Equal(t, c.GetUnits(), "km/h")
}

func TestNumericSpnEncoder(t *testing.T) {
	{
		numeric := NewNumericSPN(10, "test_byte_size_1", 2, 2, -10, 1, "L")

		assert.Equal(t, numeric.GetByteSize(), uint8(1))

		numeric.SetFormattedValue(56)

		length := numeric.GetByteSize()

		buff := make([]byte, length)
		raw := []byte{0x21}

		err := numeric.Encode(buff)
		assert.Nil(t, err)
		assert.Equal(t, raw, buff)
		fmt.Println(numeric.ToString())
		//fmt.Println(fmt.Sprintf("%x\n", buff))
	}

	////////////////////////////////////////////////////////////

	{
		numeric1 := NewNumericSPN(100, "test_byte_size_2", 3, 2.5, -2, 2, "%")
		assert.Equal(t, numeric1.GetByteSize(), uint8(2))

		numeric1.SetFormattedValue(38)

		length := numeric1.GetByteSize()

		buff1 := make([]byte, length)
		raw1 := []byte{0x10, 0x00}

		err := numeric1.Encode(buff1)
		assert.Nil(t, err)
		assert.Equal(t, raw1, buff1)
		fmt.Println(numeric1.ToString())
	}

	////////////////////////////////////////////////////////////

	{
		numeric2 := NewNumericSPN(200, "test_byte_size_3", 3, 2.1, -225, 3, "m")
		assert.Equal(t, numeric2.GetByteSize(), uint8(3))

		numeric2.SetFormattedValue(5211937.2)

		length := numeric2.GetByteSize()

		buff2 := make([]byte, length)
		raw2 := []byte{0x3E, 0xDF, 0x25}

		err := numeric2.Encode(buff2)
		assert.Nil(t, err)
		assert.Equal(t, raw2, buff2)
		fmt.Println(numeric2.ToString())
	}

	////////////////////////////////////////////////////////////

	{
		numeric3 := NewNumericSPN(250, "test_byte_size_4", 4, 0.05, 3000, 4, "Pa")
		assert.Equal(t, numeric3.GetByteSize(), uint8(4))

		numeric3.SetFormattedValue(65495982.25)

		length := numeric3.GetByteSize()

		buff3 := make([]byte, length)
		raw3 := []byte{0x3D, 0xDF, 0x12, 0x4E}

		err := numeric3.Encode(buff3)
		assert.Nil(t, err)
		assert.Equal(t, raw3, buff3)
		fmt.Println(numeric3.ToString())
	}
}

func TestNumericSpnDecoder(t *testing.T) {
	{
		numeric := NewNumericSPN(10, "test_byte_size_1", 2, 2, -10, 1, "L")
		assert.Equal(t, numeric.GetByteSize(), uint8(1))

		raw := []byte{0x21}
		err := numeric.Decode(raw)
		assert.Nil(t, err)
		assert.Equal(t, numeric.GetFormattedValue(), float64(56))
	}

	////////////////////////////////////////////////////////////

	{
		numeric1 := NewNumericSPN(100, "test_byte_size_2", 3, 2.5, -2, 2, "%")
		assert.Equal(t, numeric1.GetByteSize(), uint8(2))

		raw1 := []byte{0x10, 0x00}
		err := numeric1.Decode(raw1)
		assert.Nil(t, err)
		assert.Equal(t, numeric1.GetFormattedValue(), float64(38))
	}

	////////////////////////////////////////////////////////////

	{
		numeric2 := NewNumericSPN(200, "test_byte_size_3", 3, 2.1, -225, 3, "m")
		assert.Equal(t, numeric2.GetByteSize(), uint8(3))
		raw := []byte{0x3E, 0xDF, 0x25}
		err := numeric2.Decode(raw)
		assert.Nil(t, err)
		assert.Equal(t, numeric2.GetFormattedValue(), 5211937.2)
	}

	////////////////////////////////////////////////////////////

	{
		numeric3 := NewNumericSPN(250, "test_byte_size_4", 4, 0.05, 3000, 4, "Pa")
		assert.Equal(t, numeric3.GetByteSize(), uint8(4))
		raw := []byte{0x3D, 0xDF, 0x12, 0x4E}
		err := numeric3.Decode(raw)
		assert.Nil(t, err)
		assert.Equal(t, numeric3.GetFormattedValue(), 65495982.25)
	}

}
