package spn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-j1939/spn/spec"
	"testing"
)

var brakeSwitch *StatusSPN = nil
var clutchSwitch *StatusSPN = nil
var ptoState *StatusSPN = nil
var testState *StatusSPN = nil

func init() {
	{
		var valueToDesc = spec.DescMap{}
		valueToDesc[0] = "Desc 0"
		valueToDesc[1] = "Desc 1"
		valueToDesc[2] = "Desc 2"
		valueToDesc[3] = "Desc 3"

		testState = NewStatusSPN(100, "test_status", 4, 2, 2, valueToDesc)
	}

	{
		var valueToDesc = spec.DescMap{}
		valueToDesc[0] = "Pedal Released"
		valueToDesc[1] = "Pedal Depressed"
		valueToDesc[2] = "Error"
		valueToDesc[3] = "Not Available"

		brakeSwitch = NewStatusSPN(597, "Brake Switch", 3, 4, 2, valueToDesc)
	}

	{
		var valueToDesc = spec.DescMap{}
		valueToDesc[0] = "Pedal Released"
		valueToDesc[1] = "Pedal Depressed"
		valueToDesc[2] = "Error"
		valueToDesc[3] = "Not Available"

		clutchSwitch = NewStatusSPN(598, "Clutch Switch", 3, 6, 2, valueToDesc)
	}

	{
		var valueToDesc = spec.DescMap{}
		valueToDesc[0] = "Off"
		valueToDesc[5] = "Set"

		ptoState = NewStatusSPN(976, "PTO State", 6, 0, 5, valueToDesc)
	}
}

func TestStatusSpnBaseValue(t *testing.T) {

	assert.Equal(t, testState.GetSpnNumber(), uint32(100))
	assert.Equal(t, testState.GetName(), "test_status")
	assert.Equal(t, testState.GetOffset(), uint32(4))
	assert.Equal(t, testState.GetBitOffset(), uint8(2))
	assert.Equal(t, testState.GetBitSize(), uint8(2))
	assert.Equal(t, testState.GetType(), StatusType)

	assert.Equal(t, testState.GetValueDescription(0), "Desc 0")
	assert.Equal(t, testState.GetValueDescription(1), "Desc 1")
	assert.Equal(t, testState.GetValueDescription(2), "Desc 2")
	assert.Equal(t, testState.GetValueDescription(3), "Desc 3")
}

func TestStatusSpnSetValue(t *testing.T) {
	brakeSwitch.SetStatusValue(2)
	assert.Equal(t, brakeSwitch.GetSpnNumber(), uint32(597))
	assert.Equal(t, brakeSwitch.GetOffset(), uint32(3))
	assert.Equal(t, brakeSwitch.GetStatusValue(), uint8(2))
	assert.Equal(t, brakeSwitch.GetBitOffset(), uint8(4))
	assert.Equal(t, brakeSwitch.GetBitSize(), uint8(2))
	assert.Equal(t, brakeSwitch.GetType(), StatusType)
	fmt.Println(brakeSwitch.ToString())

	clutchSwitch.SetStatusValue(1)
	assert.Equal(t, clutchSwitch.GetSpnNumber(), uint32(598))
	assert.Equal(t, clutchSwitch.GetOffset(), uint32(3))
	assert.Equal(t, clutchSwitch.GetStatusValue(), uint8(1))
	assert.Equal(t, clutchSwitch.GetBitOffset(), uint8(6))
	assert.Equal(t, clutchSwitch.GetBitSize(), uint8(2))
	assert.Equal(t, clutchSwitch.GetType(), StatusType)
	fmt.Println(clutchSwitch.ToString())

	ptoState.SetStatusValue(5)
	assert.Equal(t, ptoState.GetSpnNumber(), uint32(976))
	assert.Equal(t, ptoState.GetOffset(), uint32(6))
	assert.Equal(t, ptoState.GetStatusValue(), uint8(5))
	assert.Equal(t, ptoState.GetBitOffset(), uint8(0))
	assert.Equal(t, ptoState.GetBitSize(), uint8(5))
	assert.Equal(t, ptoState.GetByteSize(), uint8(1))
	assert.Equal(t, ptoState.GetType(), StatusType)
	fmt.Println(ptoState.ToString())
}

func TestStatusSpnDecoder(t *testing.T) {

	c := brakeSwitch

	var buf = []byte{0x20}
	err := c.Decode(buf)
	assert.Nil(t, err)

	fmt.Println(c.ToString())
}

func TestStatusSpnEncoder(t *testing.T) {
	var valueToDesc = spec.DescMap{}

	{
		status1 := NewStatusSPN(1, "test_status1", 4, 0, 2, valueToDesc)
		status2 := NewStatusSPN(2, "test_status2", 4, 2, 2, valueToDesc)
		status3 := NewStatusSPN(3, "test_status3", 4, 4, 2, valueToDesc)
		status4 := NewStatusSPN(4, "test_status4", 4, 6, 2, valueToDesc)

		//Test encoding
		status1.SetStatusValue(0)
		status2.SetStatusValue(1)
		status3.SetStatusValue(2)
		status4.SetStatusValue(3)

		buf := make([]byte, 1)

		err := status1.Encode(buf)
		assert.Nil(t, err)

		err = status2.Encode(buf)
		assert.Nil(t, err)

		err = status3.Encode(buf)
		assert.Nil(t, err)

		err = status4.Encode(buf)
		assert.Nil(t, err)

		assert.Equal(t, buf[0], byte(0xE4))
	}

	{
		status1 := NewStatusSPN(10, "test_status1", 1, 0, 2, valueToDesc)
		status2 := NewStatusSPN(20, "test_status2", 1, 2, 2, valueToDesc)
		status3 := NewStatusSPN(30, "test_status3", 1, 4, 4, valueToDesc)

		status1.SetStatusValue(2)
		status2.SetStatusValue(1)
		status3.SetStatusValue(7)

		buf := make([]byte, 1)

		err := status1.Encode(buf)
		assert.Nil(t, err)

		err = status2.Encode(buf)
		assert.Nil(t, err)

		err = status3.Encode(buf)
		assert.Nil(t, err)

		assert.Equal(t, buf[0], byte(0x76))
	}

	{
		status1 := NewStatusSPN(10, "test_status1", 1, 0, 5, valueToDesc)
		status2 := NewStatusSPN(20, "test_status2", 1, 5, 3, valueToDesc)

		status1.SetStatusValue(30)
		status2.SetStatusValue(6)

		buf := make([]byte, 1)

		err := status1.Encode(buf)
		assert.Nil(t, err)

		err = status2.Encode(buf)
		assert.Nil(t, err)

		assert.Equal(t, buf[0], byte(0xDE))
	}

	//fmt.Println(c.ToString())
	//fmt.Println(fmt.Sprintf("%x\n", buf))
}
