package generic_frame

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-j1939/spn"
	"go-j1939/spn/spec"
	"testing"
)

var ccvs *GenericFrame = nil
var vin *GenericFrame = nil

func init() {
	initCCVS()
	initVIN()
}

func initCCVS() {
	//Jl939 Cruise Control Vehicle Speed (CCVS)
	ccvs = NewGenericFrame(0xFEF1)
	ccvs.SetName("CCVS")
	ccvs.SetLength(8)

	spnNum := spn.NewNumericSPN(84, "Wheel Speed", 1, 0.00390625, 0, 2, "km/h")
	ccvs.RegisterSPN(spnNum)

	{
		var valueToDesc = spec.DescMap{}
		valueToDesc[0] = "Pedal Released"
		valueToDesc[1] = "Pedal Depressed"
		valueToDesc[2] = "Error"
		valueToDesc[3] = "Not Available"

		spnStat := spn.NewStatusSPN(597, "Brake Switch", 3, 4, 2, valueToDesc)

		ccvs.RegisterSPN(spnStat)
	}

	{
		var valueToDesc = spec.DescMap{}
		valueToDesc[0] = "Pedal Released"
		valueToDesc[1] = "Pedal Depressed"
		valueToDesc[2] = "Error"
		valueToDesc[3] = "Not Available"

		spnStat := spn.NewStatusSPN(598, "Clutch Switch", 3, 6, 2, valueToDesc)

		ccvs.RegisterSPN(spnStat)
	}

	{
		var valueToDesc = spec.DescMap{}
		valueToDesc[0] = "Off"
		valueToDesc[5] = "Set"

		spnStat := spn.NewStatusSPN(976, "PTO State", 6, 0, 5, valueToDesc)

		ccvs.RegisterSPN(spnStat)
	}
}

func initVIN() {
	vin = NewGenericFrame(0xFEEC)
	vin.SetName("VIN")

	vinSpn := spn.NewStringSPN(237, "Vehicle Number Identifier")
	vin.RegisterSPN(vinSpn)
}

func TestSPN(t *testing.T) {

	assert.True(t, ccvs.HasSPN(84))
	assert.NotNil(t, ccvs.GetSPN(84))
	assert.Equal(t, ccvs.GetSPN(84).GetType(), spn.NumericType)

	assert.True(t, ccvs.HasSPN(597))
	assert.NotNil(t, ccvs.GetSPN(597))
	assert.Equal(t, ccvs.GetSPN(597).GetType(), spn.StatusType)

	assert.True(t, ccvs.HasSPN(598))
	assert.NotNil(t, ccvs.GetSPN(598))
	assert.Equal(t, ccvs.GetSPN(598).GetType(), spn.StatusType)

	assert.True(t, ccvs.HasSPN(976))
	assert.NotNil(t, ccvs.GetSPN(976))
	assert.Equal(t, ccvs.GetSPN(976).GetType(), spn.StatusType)

	assert.True(t, vin.HasSPN(237))
	assert.NotNil(t, vin.GetSPN(237))
	assert.Equal(t, vin.GetSPN(237).GetType(), spn.StringType)

	fmt.Println(ccvs.ToString())
	fmt.Println(vin.ToString())
}

func TestEncode(t *testing.T) {
	wheelSpeed := ccvs.GetSPN(84)
	wheelSpeed.SetFormattedValue(50) //50 kph

	{
		brakeSwitch := ccvs.GetSPN(597)

		brakeSwitch.SetStatusValue(2) //Error
	}

	{
		clutchSwitch := ccvs.GetSPN(598)

		clutchSwitch.SetStatusValue(1) //Pedal Depressed
	}

	{
		ptoState := ccvs.GetSPN(976)

		ptoState.SetStatusValue(5) //Set
	}

	ccvs.SetSrcAddr(0x50)
	ccvs.SetPriority(7)

	length := ccvs.GetDataLength()
	assert.Equal(t, length, uint32(8))

	buffer := make([]byte, length)
	for i := 0; i < int(length); i++ {
		buffer[i] = 0xFF
	}

	var identifier uint32 = 0

	err := ccvs.Encode(&identifier, buffer)
	assert.Nil(t, err)
	assert.Equal(t, identifier, uint32(0x1CFEF150))

	encodedCCVS := []byte{0xFF, 0x00, 0x32, 0x6F, 0xFF, 0xFF, 0xE5, 0xFF}
	confirm := true
	for i := 0; i < 8; i++ {
		if encodedCCVS[i] != buffer[i] {
			confirm = false
		}
	}
	assert.Equal(t, confirm, true)

	vinSpn := vin.GetSPN(237)
	vinSpn.SetStringValue("abcdefghjk1234")

	length = vin.GetDataLength()
	assert.Equal(t, length, uint32(15))

	vin.SetSrcAddr(0x30)
	vin.SetPriority(6)

	buffer = make([]byte, length)

	err = vin.Encode(&identifier, buffer)
	assert.Nil(t, err)
	assert.Equal(t, identifier, uint32(0x18FEEC30))

	confirm = true
	encodedVIN := "abcdefghjk1234*"
	for i := 0; i < len(encodedVIN); i++ {
		if encodedVIN[i] != buffer[i] {
			confirm = false
		}
	}
	assert.Equal(t, confirm, true)
}

func TestDecode(t *testing.T) {
	encodedCCVS := []byte{0xFF, 0x00, 0x50, 0x9F, 0xFF, 0xFF, 0x1F, 0xFF}
	{
		id := uint32(0x18FEF120)

		err := ccvs.Decode(id, encodedCCVS)
		assert.Nil(t, err)

		assert.Equal(t, ccvs.GetSrcAddr(), uint32(0x20))
		assert.Equal(t, ccvs.GetPriority(), uint8(6))

		wheelSpeed := ccvs.GetSPN(84)
		assert.Equal(t, wheelSpeed.GetFormattedValue(), float64(80))

		brakeSwitch := ccvs.GetSPN(597)
		assert.Equal(t, brakeSwitch.GetStatusValue(), uint8(1))

		clutchSwitch := ccvs.GetSPN(598)
		assert.Equal(t, clutchSwitch.GetStatusValue(), uint8(2))

		ptoState := ccvs.GetSPN(976)
		assert.Equal(t, ptoState.GetStatusValue(), uint8(0x1F))
	}

	{
		id := uint32(0x18FEF320)
		err := ccvs.Decode(id, encodedCCVS)
		assert.NotNil(t, err)
	}

	{
		id := uint32(0x04FEEC15)
		strVinTest := "ghijklmnopqrs*"
		bufVinTest := []byte(strVinTest)
		err := vin.Decode(id, bufVinTest)
		assert.Nil(t, err)

		assert.Equal(t, vin.GetSrcAddr(), uint32(0x15))
		assert.Equal(t, vin.GetPriority(), uint8(1))

		vinSpn := vin.GetSPN(237)
		assert.Equal(t, vinSpn.GetStringValue(), "ghijklmnopqrs")
	}
}
