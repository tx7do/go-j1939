package j1939_factory

import (
	"github.com/stretchr/testify/assert"
	"go-j1939/addressing"
	diagnosis "go-j1939/diagnosis/frames"
	fms "go-j1939/fms/tell_tale"
	"go-j1939/frames"
	"go-j1939/generic_frame"
	"go-j1939/spn"
	"go-j1939/transport"
	"testing"
)

func TestRegisterPredefinedFrames(t *testing.T) {
	fac := GetJ1939Factory()

	assert.NotNil(t, fac.GetJ1939FrameByPgn(transport.TPCMPgn))
	assert.NotNil(t, fac.GetJ1939FrameByPgn(transport.TPDTPgn))
	assert.NotNil(t, fac.GetJ1939FrameByPgn(fms.FMS1Pgn))
	assert.NotNil(t, fac.GetJ1939FrameByPgn(diagnosis.Dm1Pgn))
	assert.NotNil(t, fac.GetJ1939FrameByPgn(addressing.AddressClaimPgn))
	assert.NotNil(t, fac.GetJ1939FrameByPgn(frames.RequestPgn))

	assert.NotNil(t, fac.GetJ1939FrameByName(transport.TPCMName))
	assert.NotNil(t, fac.GetJ1939FrameByName(transport.TPDTName))
	assert.NotNil(t, fac.GetJ1939FrameByName(fms.FMS1Name))
	assert.NotNil(t, fac.GetJ1939FrameByName(diagnosis.Dm1Name))
	assert.NotNil(t, fac.GetJ1939FrameByName(addressing.AddressClaimName))
	assert.NotNil(t, fac.GetJ1939FrameByName(frames.RequestFrameName))
}

func TestRegisterDatabaseFrames(t *testing.T) {
	fac := GetJ1939Factory()

	assert.NotNil(t, fac.RegisterDatabaseFrames("test_not_found.json"))
	assert.Nil(t, fac.RegisterDatabaseFrames("D:\\GoProject\\go-j1939\\database\\frames.json"))

	{
		assert.NotNil(t, fac.GetJ1939FrameByPgn(65260))
		frame := fac.GetJ1939FrameByName("VI")
		assert.NotNil(t, frame)
		assert.Equal(t, frame.GetName(), "VI")
		
		gf := frame.(*generic_frame.GenericFrame)
		assert.NotNil(t, gf.GetSPN(237))
		assert.Equal(t, gf.GetSPN(237).GetName(), "Vehicle Number Identifier")
		assert.Equal(t, gf.GetSPN(237).GetSpnNumber(), uint32(237))
		assert.Equal(t, gf.GetSPN(237).GetType(), spn.StringType)
	}

	{
		assert.NotNil(t, fac.GetJ1939FrameByPgn(65131))
		frame := fac.GetJ1939FrameByName("DI")
		assert.NotNil(t, frame)
		assert.Equal(t, frame.GetName(), "DI")

		gf := frame.(*generic_frame.GenericFrame)

		assert.NotNil(t, gf.GetSPN(1625))
		assert.Equal(t, gf.GetSPN(1625).GetName(), "Driver 1 Identification")
		assert.Equal(t, gf.GetSPN(1625).GetSpnNumber(), uint32(1625))
		assert.Equal(t, gf.GetSPN(1625).GetType(), spn.StringType)

		assert.NotNil(t, gf.GetSPN(1626))
		assert.Equal(t, gf.GetSPN(1626).GetName(), "Driver 2 Identification")
		assert.Equal(t, gf.GetSPN(1626).GetSpnNumber(), uint32(1626))
		assert.Equal(t, gf.GetSPN(1626).GetType(), spn.StringType)
	}

}
