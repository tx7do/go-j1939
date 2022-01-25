package frames

import (
	"errors"
	j1939 "go-j1939"
	"go-j1939/diagnosis"
	"go-j1939/generic_frame"
	"go-j1939/spn"
	"go-j1939/spn/spec"
	"go-j1939/util/math"
)

const (
	Dm1Pgn           = 0x00FECA
	Dm1Name          = "DM1"
	Dm1MinimumLength = 8
)

type DTCSet []*diagnosis.DTC

type DM1 struct {
	generic_frame.GenericFrame
	dtcs DTCSet
}

func NewDM1() DM1 {
	c := DM1{}
	c.SetPGN(Dm1Pgn)
	c.SetName(Dm1Name)
	c.SetDstAddr(j1939.InvalidAddress)
	c.SetSrcAddr(j1939.InvalidAddress)
	c.SPNs = generic_frame.SpnMap{}
	c.registerDefaultStatusSpn()
	return c
}

func NewDM1AndDecode(identifier uint32, buffer []byte) (DM1, error) {
	c := DM1{}
	c.SetPGN(Dm1Pgn)
	c.SetName(Dm1Name)
	c.SPNs = generic_frame.SpnMap{}
	c.registerDefaultStatusSpn()
	err := c.Decode(identifier, buffer)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c *DM1) registerDefaultStatusSpn() {
	var valueToDesc spec.DescMap

	s := spn.NewStatusSPN(1213, "Malfunction indicator Lamp Status", 0, 6, 2, valueToDesc)
	c.RegisterSPN(s)

	s = spn.NewStatusSPN(623, "Red Stop Lamp Status", 0, 4, 2, valueToDesc)
	c.RegisterSPN(s)

	s = spn.NewStatusSPN(624, "Amber Warning Lamp Status", 0, 2, 2, valueToDesc)
	c.RegisterSPN(s)

	s = spn.NewStatusSPN(987, "Protect Lamp Status", 0, 0, 2, valueToDesc)
	c.RegisterSPN(s)

	s = spn.NewStatusSPN(3038, "Flash Malfunction indicator Lamp Status", 1, 6, 2, valueToDesc)
	c.RegisterSPN(s)

	s = spn.NewStatusSPN(3039, "Flash Red Stop Lamp Status", 1, 4, 2, valueToDesc)
	c.RegisterSPN(s)

	s = spn.NewStatusSPN(3040, "Flash Amber Warning Lamp Status", 1, 2, 2, valueToDesc)
	c.RegisterSPN(s)

	s = spn.NewStatusSPN(3041, "Flash Protect Lamp Status", 1, 0, 2, valueToDesc)
	c.RegisterSPN(s)
}

// GetDataLength 获取数据的长度
func (c *DM1) GetDataLength() uint32 {
	dtcLength := uint32(len(c.dtcs))
	//The length of DM1 frame is at least 8 bytes
	return uint32(math.Max(int64(c.GenericFrame.GetDataLength()+diagnosis.DtcFrameSize*dtcLength), Dm1MinimumLength))
}

func (c *DM1) AddDTC(dtc *diagnosis.DTC) {
	c.dtcs = append(c.dtcs, dtc)
}

func (c *DM1) DeleteDTC(pos uint32) bool {
	length := uint32(len(c.dtcs))
	if pos >= length {
		return false
	}

	c.dtcs = append(c.dtcs[:pos], c.dtcs[pos+1:]...)

	return true
}

func (c *DM1) SetDTC(pos uint32, dtc *diagnosis.DTC) bool {
	length := uint32(len(c.dtcs))
	if pos >= length {
		return false
	}
	if dtc == nil {
		return false
	}
	c.dtcs[pos] = dtc
	return true
}

func (c *DM1) GetDTCs() DTCSet {
	return c.dtcs
}
func (c *DM1) GetDTCCount() int {
	return len(c.dtcs)
}

func (c *DM1) ToString() string {
	retStr := c.GenericFrame.ToString()
	length := len(c.dtcs)
	for i := 0; i < length; i++ {
		retStr += c.dtcs[i].ToString()
	}
	return retStr
}

// Decode 解码
func (c *DM1) Decode(identifier uint32, buffer []byte) error {
	lampStatLength := c.GenericFrame.GetDataLength()

	//Decode Lamp Status (SPNs)
	err := c.GenericFrame.Decode(identifier, buffer[:lampStatLength])
	if err != nil {
		return err
	}

	length := uint32(len(buffer))
	for offset := lampStatLength; offset < length; offset = offset + diagnosis.DtcFrameSize {
		dtc := diagnosis.NewDTCAndDecode(buffer[offset : offset+diagnosis.DtcFrameSize])

		//To avoid adding a DTC when there are no faults (a DTC set all to 0s is sent which is not a valid DTC)
		if dtc.GetSpn() != 0 {
			c.AddDTC(dtc)
		}
	}

	return nil
}

// Encode 编码
func (c *DM1) Encode(identifier *uint32, buffer []byte) error {
	//Encode SPNs for bytes 0-1
	lampStatLength := c.GenericFrame.GetDataLength()

	err := c.GenericFrame.Encode(identifier, buffer[:lampStatLength])
	if err != nil {
		return err
	}

	//Must be 2
	if lampStatLength != 2 {
		return errors.New(
			"[DM1::Encode] SPNs are not expected to fit within more than 2 bytes",
		)
	}

	offset := int(lampStatLength)
	lengthDTC := len(c.dtcs)
	for i := 0; i < lengthDTC; i++ {
		_ = c.dtcs[i].Encode(buffer[offset : offset+diagnosis.DtcFrameSize])
		offset += diagnosis.DtcFrameSize
	}

	if lengthDTC == 0 {
		for i := offset; i < diagnosis.DtcFrameSize; i++ {
			buffer[i] = 0x00
		}
		for i := offset + diagnosis.DtcFrameSize; i < 2; i++ {
			buffer[i] = 0xFF
		}
	}

	return nil
}
