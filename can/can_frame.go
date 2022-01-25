package can

import "fmt"

const (
	MAX_CAN_DATA_SIZE = 8
)

type Frame struct {
	ExtendedFormat bool
	Id             uint32
	Data           string
}

func NewCanFrame() Frame {
	c := Frame{}
	c.Id = 0
	c.ExtendedFormat = false
	c.Data = ""
	return c
}

func NewCanFrameWithId(extFormat bool, id uint32) Frame {
	c := Frame{}
	c.Id = id
	c.ExtendedFormat = extFormat
	c.Data = ""
	return c
}

func NewCanFrameWithData(extFormat bool, id uint32, data string) Frame {
	c := Frame{}
	c.Id = id
	c.ExtendedFormat = extFormat
	c.SetData(data)
	return c
}

func (c *Frame) GetData() string {
	return c.Data
}
func (c *Frame) SetData(data string) bool {
	if len(data) > MAX_CAN_DATA_SIZE {
		return false
	}
	c.Data = data
	return true
}

func (c *Frame) GetId() uint32 {
	return c.Id
}
func (c *Frame) SetId(id uint32) {
	c.Id = id
}

func (c *Frame) Clear() {
	c.Id = 0
	c.Data = ""
}

func (c *Frame) IsExtendedFormat() bool {
	return c.ExtendedFormat
}
func (c *Frame) SetExtendedFormat(extendedFormat bool) {
	c.ExtendedFormat = extendedFormat
}

// HexDump 解析为十六进制数据
func (c *Frame) HexDump() string {
	var str string
	for _, v := range c.Data {
		str += fmt.Sprintf("0%.2d ", v&0xFF)
	}
	return str
}
