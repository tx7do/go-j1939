package generic_frame

import (
	"errors"
	"fmt"
	"go-j1939/j1939_frame"
	"go-j1939/spn"
	"go-j1939/util/math"
)

type SpnMap map[uint32]spn.SPN

type GenericFrame struct {
	j1939_frame.J1939FrameImpl

	Length uint32
	SPNs   SpnMap
}

func NewGenericFrame(pgn uint32) *GenericFrame {
	c := &GenericFrame{}
	c.SetPGN(pgn)
	c.SPNs = SpnMap{}
	c.Length = 0
	return c
}

func (c *GenericFrame) SetName(name string) {
	c.Name = name
}

func (c *GenericFrame) SetLength(length uint32) {
	c.Length = length
}

func (c *GenericFrame) IsGenericFrame() bool {
	return true
}

func (c *GenericFrame) GetDataLength() uint32 {
	var maxOffset uint32 = 0
	var sizeLastSpn uint8 = 1

	for _, v := range c.SPNs {
		var offset uint32 = 0
		var byteSize uint8 = 0
		offset = v.GetOffset()
		byteSize = v.GetByteSize()
		if maxOffset > offset {
			continue
		}

		maxOffset = offset
		sizeLastSpn = byteSize
	}

	//If we have specified a length, return the maximum value between the real size and the specified one.
	//This is done if not all spns for this frame are defined, and we can have a length smaller than the one necessary to transmit the frame.
	return uint32(math.Max(int64(c.Length), int64(maxOffset+uint32(sizeLastSpn))))
}

func (c *GenericFrame) ToString() string {
	retStr := c.J1939FrameImpl.ToString()
	for _, v := range c.SPNs {
		retStr += v.ToString()
	}
	return retStr
}

// RegisterSPN 注册一个spn
func (c *GenericFrame) RegisterSPN(_spn spn.SPN) *spn.SPN {
	f, ok := c.SPNs[_spn.GetSpnNumber()]
	if ok {
		f = _spn
		return &f
	}

	c.SPNs[_spn.GetSpnNumber()] = _spn

	return &_spn
}

func (c *GenericFrame) RecalculateStringOffsets() {
	var preSpn spn.SPN = nil
	for _, v := range c.SPNs {
		if v == nil {
			continue
		}
		if v.GetType() != spn.StringType {
			continue
		}
		if v != nil && preSpn != nil {
			v.SetOffset(preSpn.GetOffset() + uint32(preSpn.GetByteSize()))
		}
		if v != nil {
			preSpn = v
		}
	}
}

// DeleteSPN 删除一个spn
func (c *GenericFrame) DeleteSPN(number uint32) {
	delete(c.SPNs, number)
}

// GetSPNNumbers 获取spn编号的集合
func (c *GenericFrame) GetSPNNumbers() []uint32 {
	length := len(c.SPNs)
	if length == 0 {
		return nil
	}

	numSet := make([]uint32, length)
	for k := range c.SPNs {
		numSet = append(numSet, k)
	}

	return numSet
}

// GetSPN 获取一个spn
func (c *GenericFrame) GetSPN(number uint32) spn.SPN {
	f, ok := c.SPNs[number]
	if ok {
		return f
	} else {
		return nil
	}
}

// HasSPN 获取一个spn
func (c *GenericFrame) HasSPN(number uint32) bool {
	_, ok := c.SPNs[number]
	return ok
}

// Decode 解码
func (c *GenericFrame) Decode(identifier uint32, buffer []byte) error {
	err := c.PreDecode(identifier)
	if err != nil {
		return err
	}

	length := uint32(len(buffer))

	for _, v := range c.SPNs {
		var e error = nil
		var offset uint32 = 0
		offset = v.GetOffset()

		if offset >= length {
			return errors.New("[GenericFrame::Decode] offset of spn is higher than frame length")
		}

		e = v.Decode(buffer[offset:])

		if e != nil {
			return errors.New("[GenericFrame::Decode] spn decode failed")
		}
	}

	return nil
}

// Encode 编码
func (c *GenericFrame) Encode(identifier *uint32, buffer []byte) error {
	err := c.PreEncode(identifier)
	if err != nil {
		return err
	}
	length := uint32(len(buffer))
	var spnNumber uint32 = 0
	var offset uint32 = 0

	for _, v := range c.SPNs {
		spnNumber = v.GetSpnNumber()
		offset = v.GetOffset()

		if offset >= length {
			return errors.New(
				fmt.Sprintf("[GenericFrame::Encode] Offset of spn is higher than frame length: SPN number: %d offset: %d length: %d",
					spnNumber, offset, c.GetDataLength()),
			)
		}

		var e error = nil
		e = v.Encode(buffer[offset : offset+uint32(v.GetByteSize())])
		if e != nil {
			return e
		}
	}

	return nil
}
