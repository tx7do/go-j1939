package j1939_frame

import (
	"errors"
	"fmt"
	j1939 "go-j1939"
)

type EJ1939Status uint8
type EJ1939PduFormat uint8

const (
	UnknownFrame = "Unknown"

	StatusOff          EJ1939Status = 0
	StatusOn           EJ1939Status = 1
	StatusError        EJ1939Status = 2
	StatusNotAvailable EJ1939Status = 3

	PduFormat1 EJ1939PduFormat = 0
	PduFormat2 EJ1939PduFormat = 1
)

type J1939Frame interface {
	GetPriority() uint8
	SetPriority(priority uint8) bool

	GetExtDataPage() uint32
	GetDataPage() uint32

	GetPDUFormat() uint32
	GetPDUSpecific() uint32

	GetPDUFormatGroup() EJ1939PduFormat

	GetSrcAddr() uint32
	SetSrcAddr(src uint32) bool

	GetDstAddr() uint32
	SetDstAddr(dst uint32) bool

	GetPGN() uint32
	SetPGN(pgn uint32)

	GetName() string
	SetName(name string)

	GetIdentifier() uint32

	GetHeader() string
	ToString() string

	IsGenericFrame() bool

	// PreDecode 解码
	PreDecode(identifier uint32) error
	// PreEncode 编码
	PreEncode(identifier *uint32) error

	// GetDataLength 获取数据的长度
	GetDataLength() uint32

	// Decode 解码
	Decode(identifier uint32, buffer []byte) error
	// Encode 编码
	Encode(identifier *uint32, buffer []byte) error
}

type J1939FrameImpl struct {
	Priority uint8
	Pgn      uint32
	SrcAddr  uint32
	DstAddr  uint32

	Name string
}

//func newJ1939Frame(pgn uint32) J1939Frame {
//	c := &J1939FrameImpl{}
//	c.Priority = 0
//	c.SrcAddr = J1939_INVALID_ADDRESS
//	c.Pgn = pgn
//	c.DstAddr = J1939_INVALID_ADDRESS
//	c.name = UnknownFrame
//	return c
//}

func (c *J1939FrameImpl) GetPriority() uint8 {
	return c.Priority
}
func (c *J1939FrameImpl) SetPriority(priority uint8) bool {
	c.Priority = uint8(uint32(priority) & j1939.PriorityMask)
	return c.Priority == priority
}

func (c *J1939FrameImpl) GetExtDataPage() uint32 {
	return (c.Pgn >> j1939.ExtDataPageOffset) & j1939.ExtDataPageMask
}

func (c *J1939FrameImpl) GetDataPage() uint32 {
	return (c.Pgn >> j1939.DataPageOffset) & j1939.DataPageMask
}

func (c *J1939FrameImpl) GetPDUFormat() uint32 {
	return (c.Pgn >> j1939.PduFmtOffset) & j1939.PduFmtMask
}

func (c *J1939FrameImpl) GetPDUSpecific() uint32 {
	return c.Pgn & (j1939.PduSpecificMask)
}

func (c *J1939FrameImpl) GetSrcAddr() uint32 {
	return c.SrcAddr
}
func (c *J1939FrameImpl) SetSrcAddr(src uint32) bool {
	c.SrcAddr = src
	return true
}

func (c *J1939FrameImpl) GetPDUFormatGroup() EJ1939PduFormat {
	if c.GetPDUFormat() < j1939.PduFmtDelimiter {
		return PduFormat1
	}
	return PduFormat2
}

func (c *J1939FrameImpl) GetDstAddr() uint32 {
	return c.DstAddr
}

func (c *J1939FrameImpl) SetDstAddr(dst uint32) bool {
	if c.GetPDUFormatGroup() == PduFormat1 {
		c.DstAddr = dst
		return true
	}
	return false
}

func (c *J1939FrameImpl) GetIdentifier() uint32 {
	identifier := c.SrcAddr

	aux := c.Pgn

	if c.GetPDUFormatGroup() == PduFormat1 {
		aux = c.Pgn | ((c.DstAddr & j1939.DstAddrMask) << j1939.DstAddrOffset)
	}

	identifier |= (aux & j1939.PgnMask) << j1939.PgnOffset
	identifier |= uint32(c.Priority) << j1939.PriorityOffset

	return identifier
}

func (c *J1939FrameImpl) GetPGN() uint32 {
	return c.Pgn
}
func (c *J1939FrameImpl) SetPGN(pgn uint32) {
	c.Pgn = pgn
}

func (c *J1939FrameImpl) GetName() string {
	return c.Name
}
func (c *J1939FrameImpl) SetName(name string) {
	c.Name = name
}

func (c *J1939FrameImpl) IsGenericFrame() bool {
	return false
}

func (c *J1939FrameImpl) GetHeader() string {
	title := "name\tPGN\tSource Address\tPDU format\t"

	if c.GetPDUFormatGroup() == PduFormat1 {
		title += "Dest Address\t"
	}

	title += "Priority\t\n"

	group := "1"
	if c.GetPDUFormatGroup() != PduFormat1 {
		group = "2"
	}
	content := fmt.Sprintf("%s\t%x\t%d\t\t%s\t\t",
		c.Name, c.Pgn, c.SrcAddr, group)

	if c.GetPDUFormatGroup() == PduFormat1 {
		content += fmt.Sprintf("%d\t", c.DstAddr)
	}

	content += fmt.Sprintf("%d\t\n", c.Priority)

	return title + content
}

func (c *J1939FrameImpl) ToString() string {
	return c.GetHeader()
}

// PreDecode 解码
func (c *J1939FrameImpl) PreDecode(identifier uint32) error {
	pgn := (identifier >> j1939.PgnOffset) & j1939.PgnMask

	//Check if PDU format belongs to the first group
	if ((pgn >> j1939.PduFmtOffset) & j1939.PduFmtMask) < j1939.PduFmtDelimiter {

		c.DstAddr = (pgn >> j1939.DstAddrOffset) & j1939.DstAddrMask
		pgn &= j1939.PduFmtMask << j1939.PduFmtOffset
	}

	if pgn != c.Pgn {
		return errors.New("[J1939Frame::PreDecode] Pgn does not match")
	}

	c.SrcAddr = identifier & j1939.SrcAddrMask
	identifier >>= j1939.PriorityOffset

	c.Priority = uint8(identifier & j1939.PriorityMask)

	return nil
}

// PreEncode 编码
func (c *J1939FrameImpl) PreEncode(identifier *uint32) error {
	prior := uint32(c.Priority) & j1939.PriorityMask

	if prior != uint32(c.Priority) {
		return errors.New("[J1939Frame::PreEncode] Priority exceeded its range")
	}

	//if length < c.GetDataLength() {
	//	return errors.New("[J1939Frame::PreEncode] Length smaller than expected"), 0
	//}

	*identifier = c.SrcAddr

	aux := c.Pgn

	if c.GetPDUFormatGroup() == PduFormat1 { //Group 1
		aux = c.Pgn | ((c.DstAddr & j1939.DstAddrMask) << j1939.DstAddrOffset)
	}

	*identifier |= (aux & j1939.PgnMask) << j1939.PgnOffset
	*identifier |= prior << j1939.PriorityOffset

	return nil
}
