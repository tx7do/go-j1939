package frames

import (
	"github.com/stretchr/testify/assert"
	"go-j1939/diagnosis"
	"testing"
)

func TestDM1FrameEncoder(t *testing.T) {
	dm1 := NewDM1()
	assert.NotNil(t, dm1)

	var dtc = diagnosis.NewDTC(5052, 12, 0, 3)
	dm1.SetSrcAddr(0x55)
	dm1.SetPriority(6)
	dm1.AddDTC(dtc)
	assert.Equal(t, dm1.GetDTCCount(), 1)
	assert.NotNil(t, dm1.GetDTCs())

	var dataBuffer = make([]byte, 6)
	var identifier uint32 = 0
	err := dm1.Encode(&identifier, dataBuffer)
	assert.Nil(t, err)
	assert.Equal(t, identifier, uint32(0x18feca55))

	var destBuffer = []byte{0xFF, 0xFF, 0xBC, 0x13, 0x0C, 0x03}
	assert.Equal(t, len(dataBuffer), len(destBuffer))
	confirm := true
	for i := 0; i < len(dataBuffer); i++ {
		if dataBuffer[i] != dataBuffer[i] {
			confirm = false
		}
	}
	assert.True(t, confirm)
}

func TestDM1FrameDecoder(t *testing.T) {
	var dataBuffer = []byte{0x55, 0xFF, 0xBC, 0x13, 0x0C, 0x03}
	var identifier uint32 = 0x18FECA41
	dm1, err := NewDM1AndDecode(identifier, dataBuffer)
	assert.NotNil(t, dm1)
	assert.Nil(t, err)
	assert.Equal(t, dm1.GetDTCCount(), 1)
	assert.Equal(t, dm1.GetPriority(), uint8(6))
	assert.Equal(t, dm1.GetSrcAddr(), uint32(0x41))
	assert.Equal(t, dm1.GetDstAddr(), uint32(0x00))
	assert.NotNil(t, dm1.GetDTCs())

	dtc := dm1.GetDTCs()[0]
	assert.NotNil(t, dtc)
	assert.Equal(t, dtc.GetSpn(), uint32(5052))
	assert.Equal(t, dtc.GetFmi(), uint8(12))
	assert.Equal(t, dtc.GetCm(), uint8(0))
	assert.Equal(t, dtc.GetOc(), uint8(3))
}
