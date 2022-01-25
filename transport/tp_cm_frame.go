package transport

import (
	"errors"
	"fmt"
	j1939 "go-j1939"
	"go-j1939/j1939_frame"
)

const (
	TPCMPgn  = 0x00EC00
	TPCMName = "Transport Connection Management"
	TPCMSize = 8

	CtrlTPCMRts   = 16
	CtrlTPCMCts   = 17
	CtrlTPCMAck   = 19
	CtrlTPCMBam   = 32
	CtrlTPCMAbort = 255
)

type TPCMFrame struct {
	j1939_frame.J1939FrameImpl

	// Control byte
	CtrlType uint8

	// Total message size, number of bytes
	TotalMsgSize uint16

	// Total number of packets
	TotalPackets uint8

	// Maximum number of packets that can be sent in response to one CTS. FF16 indicates that no limit exists for the originator.
	MaxPackets uint8

	// number of packets that can be sent. This value shall be no larger than the smaller of the two values in byte 4 and byte 5 of the RTS message.
	PacketsToTx uint8

	// Next packet number to be sent
	NextPacket uint8

	// Connection Abort reason
	AbortReason uint8

	// Parameter Group number of the packeted message
	DataPgn uint32
}

func NewTPCMFrame() TPCMFrame {
	c := TPCMFrame{}
	c.SetPGN(TPCMPgn)
	c.SetName(TPCMName)
	c.CtrlType = 0
	c.TotalMsgSize = 0
	c.TotalPackets = 0
	c.MaxPackets = 0
	c.PacketsToTx = 0
	c.NextPacket = 0
	c.AbortReason = 0
	c.DataPgn = 0
	return c
}

func (c *TPCMFrame) Clear() {
	c.CtrlType = 0
	c.TotalMsgSize = 0
	c.TotalPackets = 0
	c.MaxPackets = 0
	c.PacketsToTx = 0
	c.NextPacket = 0
	c.AbortReason = 0

	c.DataPgn = 0

	c.SetPriority(0)
	c.SetSrcAddr(0)
	c.SetDstAddr(j1939.BroadcastAddress)
}

// GetDataLength 获取数据的长度
func (c *TPCMFrame) GetDataLength() uint32 {
	return TPCMSize
}

// Decode 解码
func (c *TPCMFrame) Decode(identifier uint32, buffer []byte) error {
	err := c.PreDecode(identifier)
	if err != nil {
		return err
	}

	length := len(buffer)
	if length != TPCMSize {
		return errors.New(
			fmt.Sprintf("[TPCMFrame::Decode] Buffer length does not match the expected length. Buffer length: %d. Expected length: %d",
				length, TPCMSize))
	}

	c.CtrlType = buffer[0]

	switch c.CtrlType {

	case CtrlTPCMRts:
		c.decodeRTS(buffer[1:])
		break

	case CtrlTPCMCts:
		c.decodeCTS(buffer[1:])
		break
	case CtrlTPCMAck:
		c.decodeEndOfMsgACK(buffer[1:])
		break
	case CtrlTPCMAbort:
		c.decodeConnAbort(buffer[1:])
		break
	case CtrlTPCMBam:
		c.decodeBAM(buffer[1:])
		break
	default:
		return errors.New("[TPCMFrame::Decode] Unknown Ctrl type")
	}

	bit6 := uint32(buffer[5])
	bit7 := uint32(buffer[6])
	bit8 := uint32(buffer[7])
	c.DataPgn = bit6 | (bit7 << 8) | (bit8 << 16)

	return nil
}
func (c *TPCMFrame) decodeRTS(buffer []byte) {
	bit1 := uint16(buffer[0])
	bit2 := uint16(buffer[1])
	c.TotalMsgSize = bit1 | (bit2 << 8)
	c.TotalPackets = buffer[2]
	c.MaxPackets = buffer[3]
}
func (c *TPCMFrame) decodeCTS(buffer []byte) {
	c.PacketsToTx = buffer[0]
	c.NextPacket = buffer[1]
}
func (c *TPCMFrame) decodeEndOfMsgACK(buffer []byte) {
	bit1 := uint16(buffer[0])
	bit2 := uint16(buffer[1])
	c.TotalMsgSize = bit1 | (bit2 << 8)
	c.TotalPackets = buffer[2]
}
func (c *TPCMFrame) decodeConnAbort(buffer []byte) {
	c.AbortReason = buffer[0]
}
func (c *TPCMFrame) decodeBAM(buffer []byte) {
	bit1 := uint16(buffer[0])
	bit2 := uint16(buffer[1])
	c.TotalMsgSize = bit1 | (bit2 << 8)
	c.TotalPackets = buffer[2]
}

// Encode 编码
func (c *TPCMFrame) Encode(identifier *uint32, buffer []byte) error {
	err := c.PreEncode(identifier)
	if err != nil {
		return err
	}

	//buffer := make([]byte, TPCMSize)

	buffer[0] = c.CtrlType

	switch c.CtrlType {
	case CtrlTPCMRts:
		c.encodeRTS(buffer[1:])
		break

	case CtrlTPCMCts:
		c.encodeCTS(buffer[1:])
		break
	case CtrlTPCMAck:
		c.encodeEndOfMsgACK(buffer[1:])
		break
	case CtrlTPCMAbort:
		c.encodeConnAbort(buffer[1:])
		break
	case CtrlTPCMBam:
		c.encodeBAM(buffer[1:])
		break
	default:
		return errors.New("[TPCMFrame::Encode] Unknown Ctrl Type")
	}

	buffer[5] = byte(c.DataPgn & 0xFF)
	buffer[6] = byte((c.DataPgn >> 8) & 0xFF)
	buffer[7] = byte((c.DataPgn >> 16) & 0xFF)

	return nil
}

func (c *TPCMFrame) encodeRTS(buffer []byte) {
	buffer[0] = byte(c.TotalMsgSize & 0xFF)
	buffer[1] = byte((c.TotalMsgSize >> 8) & 0xFF)
	buffer[2] = c.TotalPackets
	buffer[3] = c.MaxPackets
}
func (c *TPCMFrame) encodeCTS(buffer []byte) {
	buffer[0] = c.PacketsToTx
	buffer[1] = c.NextPacket
}
func (c *TPCMFrame) encodeEndOfMsgACK(buffer []byte) {
	buffer[0] = byte(c.TotalMsgSize & 0xFF)
	buffer[1] = byte((c.TotalMsgSize >> 8) & 0xFF)
	buffer[2] = c.TotalPackets
}
func (c *TPCMFrame) encodeConnAbort(buffer []byte) {
	buffer[0] = c.AbortReason
}
func (c *TPCMFrame) encodeBAM(buffer []byte) {
	buffer[0] = byte(c.TotalMsgSize & 0xFF)
	buffer[1] = byte((c.TotalMsgSize >> 8) & 0xFF)
	buffer[2] = c.TotalPackets
}

func (c *TPCMFrame) GetAbortReason() uint8 {
	return c.AbortReason
}

func (c *TPCMFrame) GetCtrlType() uint8 {
	return c.CtrlType
}

func (c *TPCMFrame) GetDataPgn() uint32 {
	return c.DataPgn
}

func (c *TPCMFrame) GetMaxPackets() uint8 {
	return c.MaxPackets
}

func (c *TPCMFrame) GetNextPacket() uint8 {
	return c.NextPacket
}

func (c *TPCMFrame) GetPacketsToTx() uint8 {
	return c.PacketsToTx
}

func (c *TPCMFrame) GetTotalMsgSize() uint16 {
	return c.TotalMsgSize
}

func (c *TPCMFrame) GetTotalPackets() uint8 {
	return c.TotalPackets
}

func (c *TPCMFrame) SetAbortReason(abortReason uint8) {
	c.AbortReason = abortReason
}

func (c *TPCMFrame) SetCtrlType(ctrlType uint8) {
	c.CtrlType = ctrlType
}

func (c *TPCMFrame) SetDataPgn(dataPgn uint32) {
	c.DataPgn = dataPgn
}

func (c *TPCMFrame) SetMaxPackets(maxPackets uint8) {
	c.MaxPackets = maxPackets
}

func (c *TPCMFrame) SetNextPacket(nextPacket uint8) {
	c.NextPacket = nextPacket
}

func (c *TPCMFrame) SetPacketsToTx(packetsToTx uint8) {
	c.PacketsToTx = packetsToTx
}

func (c *TPCMFrame) SetTotalMsgSize(totalMsgSize uint16) {
	c.TotalMsgSize = totalMsgSize
}

func (c *TPCMFrame) SetTotalPackets(totalPackets uint8) {
	c.TotalPackets = totalPackets
}
