package bam

import (
	"go-j1939/j1939_frame"
)

type EBamError uint8
type FragmentsMap map[uint8]*Fragments
type J1939FrameQueue []*j1939_frame.J1939Frame

const (
	ErrorIncompleteFrame  EBamError = 0
	ErrorUnexpectedFrame  EBamError = 1
	ErrorNotBroadcastAddr EBamError = 2
	ErrorDecoding         EBamError = 3
	ErrorOk               EBamError = 4
)

type Reassembly struct {
	LastError         EBamError
	Fragments         FragmentsMap
	ReassembledFrames J1939FrameQueue
}

func (c *Reassembly) ToBeHandled(_ *j1939_frame.J1939Frame) bool {
	return false
}

func (c *Reassembly) handleFrame(_ *j1939_frame.J1939Frame) uint32 {
	var expectedSize uint32 = 0
	return expectedSize
}

func (c *Reassembly) Clear() {
	c.Fragments = FragmentsMap{}
	c.LastError = ErrorOk
	c.ReassembledFrames = J1939FrameQueue{}
}

func (c *Reassembly) SetError(status EBamError) {
	c.LastError = status
}

func (c *Reassembly) GetLastError() EBamError {
	return c.LastError
}

func (c *Reassembly) ReassembledFramesPending() bool {
	return len(c.ReassembledFrames) != 0
}

func (c *Reassembly) DequeueReassembledFrame() *j1939_frame.J1939Frame {
	length := len(c.ReassembledFrames)
	if length == 0 {
		return nil
	}

	frame := c.ReassembledFrames[0]
	c.ReassembledFrames = c.ReassembledFrames[1:]

	return frame
}
