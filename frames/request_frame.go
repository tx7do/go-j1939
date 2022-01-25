package frames

import (
	"errors"
	"fmt"
	j1939 "go-j1939"
	"go-j1939/j1939_frame"
)

const (
	RequestPgn         = 0xEA00
	RequestFrameName   = "Request"
	RequestFrameLength = 3
)

type RequestFrame struct {
	j1939_frame.J1939FrameImpl

	requestPGN uint32
}

func NewRequestFrame() RequestFrame {
	c := RequestFrame{}
	c.requestPGN = 0
	c.SetPGN(RequestPgn)
	c.SetName(RequestFrameName)
	return c
}

func NewRequestFrameWithPGN(pgn uint32) RequestFrame {
	c := RequestFrame{}
	c.requestPGN = pgn
	c.SetPGN(RequestPgn)
	c.SetName(RequestFrameName)
	return c
}

// GetDataLength 获取数据的长度
func (c *RequestFrame) GetDataLength() uint32 {
	return RequestFrameLength
}

func (c *RequestFrame) GetRequestPGN() uint32 {
	return c.requestPGN
}

func (c *RequestFrame) SetRequestPGN(requestPGN uint32) {
	c.requestPGN = requestPGN
}

// Decode 解码
func (c *RequestFrame) Decode(identifier uint32, buffer []byte) error {
	err := c.PreDecode(identifier)
	if err != nil {
		return err
	}

	length := len(buffer)
	if length != RequestFrameLength {
		return errors.New(
			fmt.Sprintf("[RequestFrame::Decode] Buffer length does not match the expected length. Buffer length: %d. Expected length: %d",
				length, RequestFrameLength))
	}

	bit1 := uint32(buffer[0])
	bit2 := uint32(buffer[1])
	bit3 := uint32(buffer[2])

	c.requestPGN = bit1
	c.requestPGN |= bit2 << 8
	c.requestPGN |= bit3 << 16
	c.requestPGN &= j1939.PgnMask

	return nil
}

// Encode 编码
func (c *RequestFrame) Encode(identifier *uint32, buffer []byte) error {

	err := c.PreEncode(identifier)
	if err != nil {
		return err
	}

	length := uint32(len(buffer))
	if length < c.GetDataLength() {
		return errors.New("[J1939Frame::Encode] Length smaller than expected")
	}

	buffer[0] = byte(c.requestPGN & 0xFF)
	buffer[1] = byte((c.requestPGN >> 8) & 0xFF)
	buffer[2] = byte((c.requestPGN >> 16) & (j1939.PgnMask >> 16))

	return nil
}

func (c *RequestFrame) ToString() string {
	retVal := c.J1939FrameImpl.ToString()

	content := fmt.Sprintf("Request PGN: %x \n", c.requestPGN)

	return retVal + content
}
