package bam

import "go-j1939/transport"

//type TPDTFrameSet []*transport.TPDTFrame

type Fragments struct {
	CMFrame  transport.TPCMFrame
	DTFrames TPDTFrameSet
}

func (c *Fragments) GetCmFrame() transport.TPCMFrame {
	return c.CMFrame
}

func (c *Fragments) SetCmFrame(cmFrame transport.TPCMFrame) {
	c.CMFrame = cmFrame
}

func (c *Fragments) GetDtFrames() TPDTFrameSet {
	return c.DTFrames
}

func (c *Fragments) AddDtFrame(dtFrame *transport.TPDTFrame) {
	c.DTFrames = append(c.DTFrames, dtFrame)
}

func (c *Fragments) GetLastSQ() uint8 {
	length := len(c.DTFrames)
	if length == 0 {
		return 0
	}

	return c.DTFrames[length-1].GetSq()
}
