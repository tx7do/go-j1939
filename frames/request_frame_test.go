package frames

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestFrameDecoder(t *testing.T) {
	frame := NewRequestFrame()

	identifier := uint32(0x10EA5821)

	{
		raw := []byte{0x05, 0xFE, 0x00}
		err := frame.Decode(identifier, raw)
		assert.Nil(t, err)

		assert.Equal(t, frame.GetSrcAddr(), uint32(0x21))
		assert.Equal(t, frame.GetDstAddr(), uint32(0x58))
		assert.Equal(t, frame.GetPriority(), uint8(4))
		assert.Equal(t, frame.GetRequestPGN(), uint32(0xFE05))
	}

	{
		raw := []byte{0x00, 0xBA, 0x00}
		err := frame.Decode(identifier, raw)
		assert.Nil(t, err)

		assert.Equal(t, frame.GetSrcAddr(), uint32(0x21))
		assert.Equal(t, frame.GetDstAddr(), uint32(0x58))
		assert.Equal(t, frame.GetPriority(), uint8(4))
		assert.Equal(t, frame.GetRequestPGN(), uint32(0xBA00))
	}
}

func TestRequestFrameEncoder(t *testing.T) {
	frame := NewRequestFrameWithPGN(0xFE05)

	frame.SetSrcAddr(0x30)
	frame.SetDstAddr(0x24)
	frame.SetPriority(4)

	assert.Equal(t, frame.GetDataLength(), uint32(3))
	assert.Equal(t, frame.GetRequestPGN(), uint32(0xFE05))

	raw := []byte{0x05, 0xFE, 0x00}

	length := 3
	buff := make([]byte, length)

	var identifier uint32 = 0

	err := frame.Encode(&identifier, buff)
	assert.Nil(t, err)

	assert.Equal(t, identifier, uint32(0x10EA2430))
	assert.Equal(t, raw, buff)

	fmt.Println(frame.ToString())

	frame.SetRequestPGN(0xBA00)
	assert.Equal(t, frame.GetRequestPGN(), uint32(0xBA00))

	raw2 := []byte{0x00, 0xBA, 0x00}
	identifier = 0
	err = frame.Encode(&identifier, buff)
	assert.Equal(t, identifier, uint32(0x10EA2430))
	assert.Nil(t, err)

	assert.Equal(t, identifier, uint32(0x10EA2430))
	assert.Equal(t, raw2, buff)
}
