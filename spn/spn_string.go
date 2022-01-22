package spn

import (
	"errors"
	"fmt"
	j1939 "go-j1939"
	"go-j1939/spn/spec"
	"strings"
)

type StringSPN struct {
	spnImpl
	Offset uint32
	Value  string
}

func NewStringSPN(number uint32, name string) *StringSPN {
	c := &StringSPN{}
	c.spnImpl.Spec = spec.NewSpec(number, name, 0)
	c.Offset = 0
	//c.Owner = nil
	return c
}

func (c *StringSPN) GetType() EType {
	return StringType
}

func (c *StringSPN) GetByteSize() uint8 {
	return uint8(len(c.Value)) + 1
}

func (c *StringSPN) GetOffset() uint32 {
	return c.Offset
}

func (c *StringSPN) SetOffset(offset uint32) {
	c.Offset = offset
}

func (c *StringSPN) SetStringValue(value string) {
	for _, c := range value {
		if (c & 0x80) != 0 {
			return //This is not ASCII. We cannot assign a string that is not an ASCII string...
		}
	}
	c.Value = value

	//if c.Owner != nil { //The offsets for SPN of type string may have changed.
	//	c.Owner.RecalculateStringOffsets()
	//}
}

func (c *StringSPN) GetStringValue() string {
	return c.Value
}

func (c *StringSPN) ToString() string {
	retVal := c.spnImpl.ToString()
	return fmt.Sprintf("%s -> Value: %s\n",
		retVal, c.Value)
}

// Decode 解码
func (c *StringSPN) Decode(buffer []byte) error {
	c.Value = ""

	bufStr := string(buffer)
	length := len(bufStr)
	if length == 0 {
		return errors.New("[StringSPN::Decode] buffer is empty")
	}

	terminatorIdx := strings.IndexAny(bufStr, string(j1939.StringTerminator))

	if terminatorIdx == -1 {
		return errors.New("[StringSPN::Decode] '*' terminator not found")
	}

	for i := 0; i < length; i++ {
		c := bufStr[i]
		if c == j1939.StringTerminator {
			break
		}
		if (c & 0x80) != 0 {
			return errors.New("[StringSPN::Decode] StringSPN is not ASCII")
		}
	}

	c.Value = bufStr[:terminatorIdx]

	//if c.Owner != nil { //The size of this SPN has changed, the offsets must be recalculated
	//	c.Owner.RecalculateStringOffsets()
	//}

	return nil
}

// Encode 编码
func (c *StringSPN) Encode(buffer []byte) error {
	lenBuffer := len(buffer)
	lenValue := len(c.Value)
	if lenValue >= lenBuffer {
		return errors.New("[StringSPN::Encode] Not enough length to encode the string")
	}

	//Copy string to the buffer
	for i := 0; i < lenValue; i++ {
		buffer[i] = c.Value[i]
	}

	//Add string terminator to need of the string
	buffer[lenValue] = j1939.StringTerminator

	return nil
}
