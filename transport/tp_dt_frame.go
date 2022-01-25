package transport

import (
	"errors"
	"fmt"
	"go-j1939"
	"go-j1939/j1939_frame"
)

const (
	TPDTPgn   = 0x00EB00
	TPDTName  = "Transport Data"
	BamDtSize = 8
)

type DataBuffer [j1939.TpDtPacketSize]byte

type TPDTFrame struct {
	j1939_frame.J1939FrameImpl
	SQ   uint8
	Data DataBuffer
}

func NewTPDTFrame() TPDTFrame {
	c := TPDTFrame{}
	c.SetPGN(TPDTPgn)
	c.SetName(TPDTName)
	c.SetSq(0)
	c.SetData(DataBuffer{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	return c
}

func NewTPDTFrameWithData(sq uint8, data DataBuffer) TPDTFrame {
	c := TPDTFrame{}
	c.SetPGN(TPDTPgn)
	c.SetName(TPDTName)
	c.SetSq(sq)
	c.SetData(data)
	return c
}

// GetDataLength 获取数据的长度
func (c *TPDTFrame) GetDataLength() uint32 {
	return BamDtSize
}

// Decode 解码
func (c *TPDTFrame) Decode(identifier uint32, buffer []byte) error {
	err := c.PreDecode(identifier)
	if err != nil {
		return err
	}

	length := len(buffer)
	if length != BamDtSize {
		return errors.New(
			fmt.Sprintf("[TPDTFrame::Decode] Buffer length does not match the expected length. Buffer length: %d. Expected length: %d",
				length, j1939.TpDtPacketSize))
	}

	c.SQ = buffer[0]

	for i := 0; i < int(j1939.TpDtPacketSize); i++ {
		c.Data[i] = buffer[i+1]
	}

	return nil
}

// Encode 编码
func (c *TPDTFrame) Encode(identifier *uint32, buffer []byte) error {
	err := c.PreEncode(identifier)
	if err != nil {
		return err
	}

	//buffer := make([]byte, BamDtSize)

	buffer[0] = c.SQ

	for i := 0; i < int(j1939.TpDtPacketSize); i++ {
		buffer[i+1] = c.Data[i]
	}

	return nil
}

func (c *TPDTFrame) GetData() DataBuffer {
	return c.Data
}
func (c *TPDTFrame) SetData(data DataBuffer) {
	c.Data = data
}

func (c *TPDTFrame) GetSq() uint8 {
	return c.SQ
}
func (c *TPDTFrame) SetSq(sq uint8) {
	c.SQ = sq
}
