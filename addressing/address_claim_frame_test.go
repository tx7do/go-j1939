package addressing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressClaimFrameDecoder(t *testing.T) {
	frame := NewAddressClaimFrame()

	id := uint32(0x14EE2535)
	raw := []byte{0x2A, 0x65, 0x00, 0x9D, 0x85, 0xF1, 0xB4, 0x2F}

	err := frame.Decode(id, raw)
	assert.Nil(t, err)
	assert.Equal(t, frame.GetDataLength(), uint32(8))
	assert.Equal(t, frame.GetSrcAddr(), uint32(0x35))
	assert.Equal(t, frame.GetDstAddr(), uint32(0x25))
	assert.Equal(t, frame.GetPriority(), uint8(5))

	name := frame.GetEcuName()
	assert.Equal(t, name.GetIdNumber(), uint32(25898))
	assert.Equal(t, name.GetManufacturerCode(), uint16(1256))
	assert.Equal(t, name.GetEcuInstance(), uint8(5))
	assert.Equal(t, name.GetFunctionInstance(), uint8(16))
	assert.Equal(t, name.GetFunction(), uint8(241))
	assert.Equal(t, name.GetVehicleSystem(), uint8(90))
	assert.Equal(t, name.GetVehicleSystemInstance(), uint8(15))
	assert.Equal(t, name.GetIndustryGroup(), uint8(2))
	assert.Equal(t, name.GetEcuInstance(), uint8(5))
	assert.False(t, name.IsArbitraryAddressCapable())

	fmt.Println(frame.ToString())
}

func TestAddressClaimFrameEncoder(t *testing.T) {
	name := NewEcuNameWithValue(25898, 1256, 5, 16, 241, 90, 15, 2, false)

	frame := NewAddressClaimFrameWithEcuName(name)

	frame.SetSrcAddr(0x35)
	frame.SetDstAddr(0x25)
	frame.SetPriority(5)

	length := frame.GetDataLength()
	assert.Equal(t, length, uint32(8))

	raw := []byte{0x2A, 0x65, 0x00, 0x9D, 0x85, 0xF1, 0xB4, 0x2F}
	buff := make([]byte, length)
	var identifier uint32 = 0

	err := frame.Encode(&identifier, buff)
	assert.Nil(t, err)
	assert.Equal(t, identifier, uint32(0x14EE2535))
	assert.Equal(t, raw, buff)
}
