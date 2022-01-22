package spn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSpnBasic(t *testing.T) {
	c := NewStringSPN(50, "string_example")
	assert.Equal(t, c.GetSpnNumber(), uint32(50))
	assert.Equal(t, c.GetName(), "string_example")
	assert.Equal(t, c.GetType(), StringType)
	assert.Equal(t, c.GetStringValue(), "")
}

func TestStringSpnOffset(t *testing.T) {
	c := NewStringSPN(50, "string_example")
	c.SetOffset(5)
	assert.Equal(t, c.GetOffset(), uint32(5))
}

func TestStringSpnEncoder(t *testing.T) {
	c := NewStringSPN(21, "string_example")

	testStr := "abcdefghijklmnopqrstuvwxyz"
	c.SetStringValue(testStr)

	lenTestStr := len(testStr) + 1
	length := int(c.GetByteSize())
	assert.Equal(t, length, lenTestStr)

	buff := make([]byte, length)
	err := c.Encode(buff)
	assert.Nil(t, err)

	strBuff := string(buff)
	strDst := testStr + "*"
	assert.Equal(t, strBuff, strDst)
}

func TestStringSpnDecoder(t *testing.T) {
	testStr := "zyxwvutsrqponmlkjihgfedcba"
	c := NewStringSPN(14, "string_example")

	var buffer = []byte(testStr + "*")
	err := c.Decode(buffer)
	assert.Nil(t, err)
	assert.Equal(t, testStr, c.GetStringValue())
}
