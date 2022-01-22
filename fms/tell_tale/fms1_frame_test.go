package tell_tale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFMS1FrameDecoder(t *testing.T) {
	{
		frame := NewFMS1Frame()

		id := uint32(0x14FD7D35)
		raw := []byte{0xB0, 0xFF, 0x98, 0xFF, 0xFA, 0xFF, 0xFF, 0xFF}

		err := frame.Decode(id, raw)
		assert.Nil(t, err)

		assert.Equal(t, frame.GetSrcAddr(), uint32(0x35))
		assert.Equal(t, frame.GetPriority(), uint8(5))

		assert.Equal(t, frame.GetBlockID(), uint8(0))

		assert.Equal(t, frame.GetTTS(5).GetStatus(), TtsStatusRed)
		assert.Equal(t, frame.GetTTS(4).GetStatus(), TtsStatusOff)
		assert.Equal(t, frame.GetTTS(8).GetStatus(), TtsStatusYellow)
		assert.Equal(t, frame.GetTTS(1).GetStatus(), TtsStatusInfo)
	}
	{
		frame := NewFMS1Frame()

		id := uint32(0x04FD7D44)
		raw := []byte{0xF1, 0xFA, 0x9F, 0xF8, 0xFF, 0xFF, 0xFF, 0xBF}

		err := frame.Decode(id, raw)
		assert.Nil(t, err)

		assert.Equal(t, frame.GetSrcAddr(), uint32(0x44))
		assert.Equal(t, frame.GetPriority(), uint8(1))

		assert.Equal(t, frame.GetBlockID(), uint8(1))

		assert.Equal(t, frame.GetTTS(20).GetStatus(), TtsStatusRed)
		assert.Equal(t, frame.GetTTS(21).GetStatus(), TtsStatusOff)
		assert.Equal(t, frame.GetTTS(17).GetStatus(), TtsStatusYellow)
		assert.Equal(t, frame.GetTTS(30).GetStatus(), TtsStatusInfo)
	}
	{
		frame := NewFMS1Frame()

		id := uint32(0x1CFD7D36)
		raw := []byte{0xF2, 0xFF, 0x9F, 0xFF, 0xFF, 0xA8, 0xFF, 0xBF}

		err := frame.Decode(id, raw)
		assert.Nil(t, err)

		assert.Equal(t, frame.GetSrcAddr(), uint32(0x36))
		assert.Equal(t, frame.GetPriority(), uint8(7))

		assert.Equal(t, frame.GetBlockID(), uint8(2))

		assert.Equal(t, frame.GetTTS(35).GetStatus(), TtsStatusRed)
		assert.Equal(t, frame.GetTTS(40).GetStatus(), TtsStatusOff)
		assert.Equal(t, frame.GetTTS(41).GetStatus(), TtsStatusYellow)
		assert.Equal(t, frame.GetTTS(45).GetStatus(), TtsStatusInfo)
	}
	{
		frame := NewFMS1Frame()

		id := uint32(0x14FD7D25)
		raw := []byte{0x93, 0xFF, 0xA8, 0xFF, 0xFF, 0xFF, 0xFF, 0xBF}

		err := frame.Decode(id, raw)
		assert.Nil(t, err)

		assert.Equal(t, frame.GetSrcAddr(), uint32(0x25))
		assert.Equal(t, frame.GetPriority(), uint8(5))

		assert.Equal(t, frame.GetBlockID(), uint8(3))

		assert.Equal(t, frame.GetTTS(46).GetStatus(), TtsStatusRed)
		assert.Equal(t, frame.GetTTS(49).GetStatus(), TtsStatusOff)
		assert.Equal(t, frame.GetTTS(50).GetStatus(), TtsStatusYellow)
		assert.Equal(t, frame.GetTTS(60).GetStatus(), TtsStatusInfo)
	}
}

func TestFMS1FrameEncoder(t *testing.T) {
	{
		frame := NewFMS1FrameWithBlockID(0)

		frame.SetSrcAddr(0x35)
		frame.SetPriority(5)

		assert.Equal(t, frame.GetDataLength(), uint32(8))

		assert.Equal(t, frame.GetBlockID(), uint8(0))

		assert.True(t, frame.HasTTS(5))
		assert.False(t, frame.HasTTS(16))
		assert.False(t, frame.HasTTS(35))
		assert.False(t, frame.HasTTS(50))

		assert.True(t, frame.SetTTS(5, TtsStatusRed))
		assert.True(t, frame.SetTTS(4, TtsStatusOff))
		assert.True(t, frame.SetTTS(8, TtsStatusYellow))
		assert.True(t, frame.SetTTS(1, TtsStatusInfo))

		length := frame.GetDataLength()
		raw := []byte{0xB0, 0xFF, 0x98, 0xFF, 0xFA, 0xFF, 0xFF, 0xFF}

		buff := make([]byte, length)
		var identifier uint32 = 0

		err := frame.Encode(&identifier, buff)
		assert.Nil(t, err)

		assert.Equal(t, identifier, uint32(0x14FD7D35))
		assert.Equal(t, raw, buff)
	}

	{
		frame := NewFMS1FrameWithBlockID(1)

		frame.SetSrcAddr(0x44)
		frame.SetPriority(1)

		assert.Equal(t, frame.GetDataLength(), uint32(8))

		assert.Equal(t, frame.GetBlockID(), uint8(1))

		assert.False(t, frame.HasTTS(5))
		assert.True(t, frame.HasTTS(16))
		assert.False(t, frame.HasTTS(35))
		assert.False(t, frame.HasTTS(50))

		assert.True(t, frame.SetTTS(20, TtsStatusRed))
		assert.True(t, frame.SetTTS(21, TtsStatusOff))
		assert.True(t, frame.SetTTS(17, TtsStatusYellow))
		assert.True(t, frame.SetTTS(30, TtsStatusInfo))

		length := frame.GetDataLength()
		raw := []byte{0xF1, 0xFA, 0x9F, 0xF8, 0xFF, 0xFF, 0xFF, 0xBF}

		buff := make([]byte, length)
		var identifier uint32 = 0

		err := frame.Encode(&identifier, buff)
		assert.Nil(t, err)

		assert.Equal(t, identifier, uint32(0x04FD7D44))
		assert.Equal(t, raw, buff)
	}

	{
		frame := NewFMS1FrameWithBlockID(2)

		frame.SetSrcAddr(0x36)
		frame.SetPriority(7)

		assert.Equal(t, frame.GetDataLength(), uint32(8))

		assert.Equal(t, frame.GetBlockID(), uint8(2))

		assert.False(t, frame.HasTTS(5))
		assert.False(t, frame.HasTTS(16))
		assert.True(t, frame.HasTTS(35))

		assert.True(t, frame.SetTTS(35, TtsStatusRed))
		assert.True(t, frame.SetTTS(40, TtsStatusOff))
		assert.True(t, frame.SetTTS(41, TtsStatusYellow))
		assert.True(t, frame.SetTTS(45, TtsStatusInfo))

		length := frame.GetDataLength()
		raw := []byte{0xF2, 0xFF, 0x9F, 0xFF, 0xFF, 0xA8, 0xFF, 0xBF}

		buff := make([]byte, length)
		var identifier uint32 = 0

		err := frame.Encode(&identifier, buff)
		assert.Nil(t, err)

		assert.Equal(t, identifier, uint32(0x1CFD7D36))
		assert.Equal(t, raw, buff)
	}

	{
		frame := NewFMS1FrameWithBlockID(3)

		frame.SetSrcAddr(0x25)
		frame.SetPriority(5)

		assert.Equal(t, frame.GetDataLength(), uint32(8))

		assert.Equal(t, frame.GetBlockID(), uint8(3))

		assert.False(t, frame.HasTTS(5))
		assert.False(t, frame.HasTTS(16))
		assert.False(t, frame.HasTTS(35))
		assert.True(t, frame.HasTTS(50))

		assert.True(t, frame.SetTTS(46, TtsStatusRed))
		assert.True(t, frame.SetTTS(49, TtsStatusOff))
		assert.True(t, frame.SetTTS(50, TtsStatusYellow))
		assert.True(t, frame.SetTTS(60, TtsStatusInfo))

		length := frame.GetDataLength()
		raw := []byte{0x93, 0xFF, 0xA8, 0xFF, 0xFF, 0xFF, 0xFF, 0xBF}

		buff := make([]byte, length)
		var identifier uint32 = 0

		err := frame.Encode(&identifier, buff)
		assert.Nil(t, err)

		assert.Equal(t, identifier, uint32(0x14FD7D25))
		assert.Equal(t, raw, buff)
	}
}
