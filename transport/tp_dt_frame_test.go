package transport

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTpDtFrameDecoder(t *testing.T) {
	c := NewTPDTFrame()
	identifier := c.GetIdentifier()
	var buf = []byte{0x02, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0xFF}
	err := c.Decode(identifier, buf)
	assert.Nil(t, err)

	fmt.Println(c.ToString())
}

func TestTpDtFrameEncoder(t *testing.T) {
	sq := uint8(0x02)
	data := DataBuffer{0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0xFF}
	c := NewTPDTFrameWithData(sq, data)

	var identifier uint32 = 0
	buff := make([]byte, BamDtSize)
	err := c.Encode(&identifier, buff)
	assert.Nil(t, err)
	fmt.Println(c.ToString())

	fmt.Println(identifier, string(buff))
}
