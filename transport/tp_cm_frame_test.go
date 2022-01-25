package transport

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTpCmFrameDecoder(t *testing.T) {
	c := NewTPCMFrame()
	identifier := c.GetIdentifier()
	var buf = []byte{0x12, 0x13}
	err := c.Decode(identifier, buf)
	assert.Nil(t, err)

	fmt.Println(c.ToString())
}

func TestTpCmFrameEncoder(t *testing.T) {
	c := NewTPCMFrame()
	var identifier uint32 = 0
	buff := make([]byte, TPCMSize)
	err := c.Encode(&identifier, buff)
	assert.Nil(t, err)
	fmt.Println(c.ToString())

	fmt.Println(identifier, string(buff))
}
