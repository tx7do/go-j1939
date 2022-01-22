package bam

import (
	j1939 "go-j1939"
	"go-j1939/j1939_frame"
	"go-j1939/transport"
	"go-j1939/util/math"
)

type TPDTFrameMap map[uint8]*transport.TPDTFrame
type TPDTFrameSet []*transport.TPDTFrame

type Fragmenter struct {
	CMFrame  transport.TPCMFrame
	DTFrames TPDTFrameMap
}

func (c *Fragmenter) Clear() {
	c.DTFrames = TPDTFrameMap{}
	c.CMFrame.Clear()
}

func (c *Fragmenter) Fragment(frame j1939_frame.J1939Frame) bool {
	length := frame.GetDataLength()
	if length <= j1939.MaxSize {
		//Not necessary to fragment the frame
		return false
	}

	//Clear previous fragmented frames
	c.Clear()

	//Bam type
	c.CMFrame.SetCtrlType(transport.CtrlTpcmBam)

	//Same source and priority as the original frame
	c.CMFrame.SetPriority(frame.GetPriority())
	c.CMFrame.SetSrcAddr(frame.GetSrcAddr())

	//Encoded pgn
	c.CMFrame.SetDataPgn(frame.GetPGN())

	//Always broadcast
	c.CMFrame.SetDstAddr(j1939.BroadcastAddress)

	var identifier uint32 = 0
	frameData := make([]byte, length)

	err := frame.Encode(&identifier, frameData)
	if err != nil {
		return false
	}

	var dataBuffer transport.DataBuffer
	var sq uint8 = 0
	var offset uint32 = 0
	for offset = 0; offset < length; offset += uint32(j1939.TpDtPacketSize) {
		sq++

		dataBuffer = transport.DataBuffer{}
		_data := frameData[offset:math.Min(int64(j1939.TpDtPacketSize), int64(length-offset))]
		for i := 0; i < len(_data); i++ {
			dataBuffer[i] = _data[i]
		}
		dataFrame := transport.NewTPDTFrameWithData(sq, dataBuffer)

		//Same source and priority as the original frame
		dataFrame.SetPriority(frame.GetPriority())
		dataFrame.SetSrcAddr(frame.GetSrcAddr())

		//Always broadcast
		dataFrame.SetDstAddr(j1939.BroadcastAddress)

		c.DTFrames[sq] = dataFrame
	}

	c.CMFrame.SetTotalPackets(uint8(len(c.DTFrames)))
	c.CMFrame.SetTotalMsgSize(uint16(length))

	return true
}

func (c *Fragmenter) GetConnFrame() transport.TPCMFrame {
	return c.CMFrame
}

func (c *Fragmenter) GetDataFrames() TPDTFrameSet {
	frames := make(TPDTFrameSet, len(c.DTFrames))
	for _, v := range c.DTFrames {
		frames = append(frames, v)
	}
	return frames
}
