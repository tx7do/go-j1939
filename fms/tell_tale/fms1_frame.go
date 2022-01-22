package tell_tale

import (
	"errors"
	"fmt"
	"go-j1939/j1939_frame"
)

const (
	BlockIdMask = 0xF

	TtsMask                = 0x7
	TtsHighPartShift       = 4
	TtssPerBlock     uint8 = 15
	TtsEncodingMask  uint8 = 0x8

	NumberOfBlocks = 4

	Fms1FrameLength = 8
	Fms1Pgn         = 0xFD7D
	Fms1Name        = "FMS1"
)

type TellTaleMap map[uint8]*TellTale

type FMS1Frame struct {
	j1939_frame.J1939FrameImpl
	BlockID uint8
	TTSs    TellTaleMap
}

func NewFMS1Frame() *FMS1Frame {
	c := &FMS1Frame{}
	c.SetPGN(Fms1Pgn)
	c.SetName(Fms1Name)
	c.BlockID = NumberOfBlocks
	c.TTSs = TellTaleMap{}
	return c
}

func NewFMS1FrameWithBlockID(blockID uint8) *FMS1Frame {
	c := &FMS1Frame{}
	c.SetPGN(Fms1Pgn)
	c.SetName(Fms1Name)
	c.BlockID = blockID
	c.TTSs = TellTaleMap{}

	for i := c.BlockID*TtssPerBlock + 1; i < (c.BlockID+1)*TtssPerBlock+1; i++ {
		c.TTSs[i] = NewTellTaleWithValue(i, TtsStatusNotAvailable)
	}

	return c
}

func (c *FMS1Frame) GetDataLength() uint32 {
	return Fms1FrameLength
}

func (c *FMS1Frame) HasTTS(number uint8) bool {
	_, ok := c.TTSs[number]
	return ok
}

func (c *FMS1Frame) GetTTS(number uint8) *TellTale {
	ret, _ := c.TTSs[number]
	return ret
}

func (c *FMS1Frame) SetTTS(number uint8, status TtsStatusType) bool {
	ret, ok := c.TTSs[number]
	if ok {
		ret.SetStatus(status)
		return true
	} else {
		ret = NewTellTaleWithValue(number, status)
		c.TTSs[number] = ret
		return false
	}
}

func (c *FMS1Frame) GetBlockID() uint8 {
	return c.BlockID
}

func (c *FMS1Frame) ToString() string {
	retStr := c.J1939FrameImpl.ToString()

	content := fmt.Sprintf("Block ID: %d\n", c.BlockID)

	for _, v := range c.TTSs {
		content += v.ToString()
	}

	return retStr + content
}

// Decode 解码
func (c *FMS1Frame) Decode(identifier uint32, buffer []byte) error {
	err := c.PreDecode(identifier)
	if err != nil {
		return err
	}

	length := len(buffer)
	if length != Fms1FrameLength { //Check the length first
		return errors.New(
			fmt.Sprintf(
				"[FMS1Frame::Decode] Buffer length does not match the expected length. Buffer length:%d. Expected length: %d",
				length, Fms1FrameLength),
		)
	}

	blockID := buffer[0] & BlockIdMask
	if blockID >= NumberOfBlocks {
		return errors.New(
			fmt.Sprintf("[FMS1Frame::Decode] Block ID higher than the maximum permitted. Max: %d",
				NumberOfBlocks-1),
		)
	}

	//If block ID changes, clear mTTSs to not accumulate the previous decoded TTSs
	//related to the previous block.

	if c.BlockID != blockID {
		c.TTSs = TellTaleMap{}
	}

	c.BlockID = blockID

	tts1Number := TtssPerBlock*c.BlockID + 1
	tts1 := NewTellTaleWithValue(tts1Number, TtsStatusType((buffer[0]>>TtsHighPartShift)&TtsMask))
	c.TTSs[tts1Number] = tts1

	var i uint8 = 0
	for i = 1; i < Fms1FrameLength; i++ {
		ttsLowPartNumber := (TtssPerBlock * c.BlockID) + 2*i
		ttsHighPartNumber := (TtssPerBlock * c.BlockID) + 2*i + 1

		ttsLowPartStatus := buffer[i] & TtsMask
		ttsHighPartStatus := (buffer[i] >> TtsHighPartShift) & TtsMask

		c.TTSs[ttsLowPartNumber] = NewTellTaleWithValue(ttsLowPartNumber, TtsStatusType(ttsLowPartStatus))
		c.TTSs[ttsHighPartNumber] = NewTellTaleWithValue(ttsHighPartNumber, TtsStatusType(ttsHighPartStatus))
	}

	return nil
}

// Encode 编码
func (c *FMS1Frame) Encode(identifier *uint32, buffer []byte) error {
	err := c.PreEncode(identifier)
	if err != nil {
		return err
	}

	length := len(c.TTSs)

	if length != int(TtssPerBlock) {
		return errors.New(
			fmt.Sprintf(
				"[FMS1Frame::Encode] There are not %d defined",
				TtssPerBlock),
		)
	}

	var maxKey uint8 = 0
	var minKey uint8 = 0
	for k := range c.TTSs {
		if k < minKey || minKey == 0 {
			minKey = k
		}
		if k > maxKey {
			maxKey = k
		}
	}
	//Check if the number for every TTS is the right one.
	if minKey <= c.BlockID*TtssPerBlock || maxKey > (c.BlockID+1)*TtssPerBlock {
		return errors.New("[FMS1Frame::Encode] TTS numbers are not the proper ones for this block")
	}

	tts1Number := TtssPerBlock*c.BlockID + 1
	tts1, ok := c.TTSs[tts1Number]
	if ok {
		buffer[0] = (c.BlockID & BlockIdMask) |
			((uint8(tts1.GetStatus()) | TtsEncodingMask) << TtsHighPartShift)
	}

	var i uint8 = 0
	for i = 1; i < Fms1FrameLength; i++ {
		ttsLowPartNumber := TtssPerBlock*c.BlockID + 2*i
		ttsHighPartNumber := TtssPerBlock*c.BlockID + 2*i + 1

		ttsLow, okLow := c.TTSs[ttsLowPartNumber]
		ttsHigh, okHigh := c.TTSs[ttsHighPartNumber]

		if okLow && okHigh {
			buffer[i] = (uint8(ttsLow.GetStatus()) | TtsEncodingMask) |
				((uint8(ttsHigh.GetStatus()) | TtsEncodingMask) << TtsHighPartShift)
		}
	}

	return nil
}
