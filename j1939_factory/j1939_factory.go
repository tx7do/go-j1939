package j1939_factory

import (
	"encoding/json"
	"errors"
	j1939 "go-j1939"
	"go-j1939/addressing"
	diagnosis "go-j1939/diagnosis/frames"
	fms "go-j1939/fms/tell_tale"
	"go-j1939/frames"
	"go-j1939/generic_frame"
	"go-j1939/j1939_frame"
	"go-j1939/spn"
	"go-j1939/spn/spec"
	"go-j1939/transport"
	"io/ioutil"
	"os"
)

type J1939FrameMap map[uint32]j1939_frame.J1939Frame

var j1939Factory J1939Factory

type J1939Factory struct {
	frames J1939FrameMap
}

// GetJ1939Factory 获取工厂实例
func GetJ1939Factory() J1939Factory {
	if j1939Factory.frames == nil {
		j1939Factory.frames = J1939FrameMap{}
	}
	if len(j1939Factory.frames) == 0 {
		j1939Factory.registerPredefinedFrames()
	}
	return j1939Factory
}

// registerPredefinedFrames 注册预先定义的数据帧
func (c *J1939Factory) registerPredefinedFrames() {
	{
		frame := transport.NewTPCMFrame()
		c.RegisterFrame(&frame)
	}

	{
		frame := transport.NewTPDTFrame()
		c.RegisterFrame(&frame)
	}

	{
		frame := fms.NewFMS1Frame()
		c.RegisterFrame(&frame)
	}

	{
		frame := diagnosis.NewDM1()
		c.RegisterFrame(&frame)
	}

	{
		frame := addressing.NewAddressClaimFrame()
		c.RegisterFrame(&frame)
	}

	{
		frame := frames.NewRequestFrame()
		c.RegisterFrame(&frame)
	}
}

// RegisterFrame 注册一个数据帧
func (c *J1939Factory) RegisterFrame(frame j1939_frame.J1939Frame) bool {
	_, ok := c.frames[frame.GetPGN()]
	if ok {
		return false
	} else {
		c.frames[frame.GetPGN()] = frame
		return true
	}
}

// UnregisterFrame  注销一个数据帧
func (c *J1939Factory) UnregisterFrame(pgn uint32) {
	delete(c.frames, pgn)
}

func (c *J1939Factory) RegisterDatabaseFrames(dbFile string) error {
	f, err := os.Open(dbFile)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	//fmt.Print(string(b))

	type FrameItem struct {
		Name   *string `json:"name,omitempty"`
		Pgn    uint32  `json:"pgn"`
		Length *uint32 `json:"length,omitempty"`
		SPNs   []struct {
			Name   string `json:"name"`
			Number uint32 `json:"number"`
			Type   uint8  `json:"type"`

			Offset    *uint32 `json:"offset,omitempty"`
			BitSize   *uint8  `json:"bitSize,omitempty"`
			BitOffset *uint8  `json:"bitOffset,omitempty"`
			ByteSize  *uint8  `json:"byteSize,omitempty"`

			FormatGain   *float64 `json:"formatGain,omitempty"`
			FormatOffset *float64 `json:"formatOffset,omitempty"`
			Units        *string  `json:"units,omitempty"`

			Descriptions []string `json:"descriptions,omitempty"`
		} `json:"spns"`
	}

	var results []FrameItem
	err = json.Unmarshal(b, &results)
	if err != nil {
		return err
	}

	for _, v := range results {
		switch v.Pgn {
		case diagnosis.Dm1Pgn:
			//Not permitted this token. Dangerous for the good performance of the software.
			return errors.New("A token that should not be present in the database which is reserved for internal use in the software")
		default:
			break
		}

		frame := generic_frame.NewGenericFrame(v.Pgn)
		if v.Name != nil {
			frame.SetName(*v.Name)
		}
		if v.Length != nil {
			frame.SetLength(*v.Length)
		}

		if v.SPNs != nil {
			for _, _spn := range v.SPNs {
				switch spn.EType(_spn.Type) {
				case spn.NumericType:
					_s := spn.NewNumericSPN(_spn.Number, _spn.Name, *_spn.Offset, *_spn.FormatGain, *_spn.FormatOffset, *_spn.ByteSize, *_spn.Units)
					frame.RegisterSPN(_s)
					break
				case spn.StatusType:
					var descMap = spec.DescMap{}
					for i, v := range _spn.Descriptions {
						descMap[uint8(i)] = v
					}
					_s := spn.NewStatusSPN(_spn.Number, _spn.Name, *_spn.Offset, *_spn.BitOffset, *_spn.BitSize, descMap)
					frame.RegisterSPN(_s)
					break
				case spn.StringType:
					_s := spn.NewStringSPN(_spn.Number, _spn.Name)
					frame.RegisterSPN(_s)
					break
				default:
					return errors.New("the type of spn is not recognized")
				}
			}
		}

		c.RegisterFrame(frame)
	}

	return nil
}

func (c *J1939Factory) unregisterAllFrames() {
	c.frames = J1939FrameMap{}
}

// GetAllRegisteredPGNs 获取注册表中的数据帧的PGN
func (c *J1939Factory) GetAllRegisteredPGNs() []uint32 {
	var pgnSet []uint32

	for _, frame := range c.frames {
		pgnSet = append(pgnSet, frame.GetPGN())
	}

	return pgnSet
}

// GetJ1939FrameFromData 获取注册表中的数据帧 - id和数据
func (c *J1939Factory) GetJ1939FrameFromData(identifier uint32, buffer []byte) j1939_frame.J1939Frame {
	pgn := (identifier >> j1939.PgnOffset) & j1939.PgnMask

	//Check if PDU format belongs to the first group
	if ((pgn >> j1939.PduFmtOffset) & j1939.PduFmtMask) < j1939.PduFmtDelimiter {
		pgn &= j1939.PduFmtMask << j1939.PduFmtOffset
	}

	frame, ok := c.frames[pgn]
	if !ok {
		return nil
	}

	err := frame.Decode(identifier, buffer)
	if err != nil {
		return nil
	}

	return frame
}

// GetJ1939FrameByPgn 获取注册表中的数据帧 - 帧PGN
func (c *J1939Factory) GetJ1939FrameByPgn(pgn uint32) j1939_frame.J1939Frame {
	frame, ok := c.frames[pgn]
	if ok {
		return frame
	} else {
		return nil
	}
}

// GetJ1939FrameByName 获取注册表中的数据帧 - 帧名字
func (c *J1939Factory) GetJ1939FrameByName(name string) j1939_frame.J1939Frame {
	for _, frame := range c.frames {
		if frame.GetName() == name {
			return frame
		}
	}
	return nil
}
