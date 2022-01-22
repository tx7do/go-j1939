package addressing

import (
	"errors"
	"fmt"
	"go-j1939/j1939_frame"
)

type AddressClaimFrame struct {
	j1939_frame.J1939FrameImpl
	ecuName *EcuName
}

func NewAddressClaimFrame() *AddressClaimFrame {
	c := &AddressClaimFrame{}
	c.SetPGN(AddressClaimPgn)
	c.SetName(AddressClaimName)
	c.ecuName = NewEcuName()

	return c
}
func NewAddressClaimFrameWithEcuName(name *EcuName) *AddressClaimFrame {
	c := &AddressClaimFrame{}
	c.SetPGN(AddressClaimPgn)
	c.SetName(AddressClaimName)
	c.ecuName = name
	return c
}

func (c *AddressClaimFrame) GetEcuName() *EcuName {
	return c.ecuName
}

// GetDataLength 获取数据的长度
func (c *AddressClaimFrame) GetDataLength() uint32 {
	return AddressFrameLength
}

func (c *AddressClaimFrame) ToString() string {
	retVal := c.J1939FrameImpl.ToString()

	capable := "yes"
	if !c.ecuName.IsArbitraryAddressCapable() {
		capable = "no"
	}

	content := fmt.Sprintf("Id number: %d \nManufacturer Code: %d\nECU Intance: %d\nFunction Instance: %d\nFunction: %d\nVehicle System: %d\nVehicle System Instance: %d\nIndustry Group: %d\nAddress Capable: %s\n",
		c.ecuName.GetIdNumber(),
		c.ecuName.GetManufacturerCode(),
		c.ecuName.GetEcuInstance(),
		c.ecuName.GetFunctionInstance(),
		c.ecuName.GetFunction(),
		c.ecuName.GetVehicleSystem(),
		c.ecuName.GetVehicleSystemInstance(),
		c.ecuName.GetIndustryGroup(),
		capable)

	return retVal + content
}

// Decode 解码
func (c *AddressClaimFrame) Decode(identifier uint32, buffer []byte) error {
	err := c.PreDecode(identifier)
	if err != nil {
		return err
	}

	length := len(buffer)
	if length != AddressFrameLength {
		return errors.New(
			fmt.Sprintf("[AdressClaimFrame::Decode] Buffer length does not match the expected length. Buffer length: %d. Expected length: %d",
				length, AddressFrameLength))
	}

	bit1 := uint32(buffer[0])
	bit2 := uint32(buffer[1])
	bit3 := uint32(buffer[2])
	bit4 := uint32(buffer[3])

	idNumber := bit1 | (bit2 << 8) | ((bit3 & 0x1F) << 16)

	manufacturerCode := uint16(((bit3 & 0xE0) >> 5) | (bit4 << 3))

	ecuInstance := buffer[4] & 0x07

	functionInstance := (buffer[4] & 0xF8) >> 3

	function := buffer[5]

	vehicleSystem := buffer[6] >> 1

	vehicleSystemInstance := buffer[7] & 0x0F

	industryGroup := (buffer[7] >> 4) & 0x07

	capable := buffer[7] >> 7
	arbitraryAddressCapable := false
	if capable != 0 {
		arbitraryAddressCapable = true
	}

	c.ecuName = NewEcuNameWithValue(idNumber, manufacturerCode, ecuInstance, functionInstance, function, vehicleSystem,
		vehicleSystemInstance, industryGroup, arbitraryAddressCapable)

	return nil
}

// Encode 编码
func (c *AddressClaimFrame) Encode(identifier *uint32, buffer []byte) error {
	err := c.PreEncode(identifier)
	if err != nil {
		return err
	}

	length := uint32(len(buffer))
	if length > c.GetDataLength() {
		return errors.New("[AddressClaimFrame::Encode] Length smaller than expected")
	}

	buffer[0] = byte(c.ecuName.GetIdNumber() & 0xFF)
	buffer[1] = byte((c.ecuName.GetIdNumber() >> 8) & 0xFF)
	buffer[2] = byte(c.ecuName.GetIdNumber()>>16)&0x1F | byte((c.ecuName.GetManufacturerCode()<<5)&0xFF)
	buffer[3] = byte((c.ecuName.GetManufacturerCode() >> 3) & 0xFF)
	buffer[4] = (c.ecuName.GetEcuInstance() & EcuInstanceMask) | ((c.ecuName.GetFunctionInstance() << 3) & 0xFF)

	buffer[5] = c.ecuName.GetFunction()

	buffer[6] = c.ecuName.GetVehicleSystem() << 1

	capable := 0
	if c.ecuName.IsArbitraryAddressCapable() {
		capable = 1
	}
	buffer[7] = c.ecuName.GetVehicleSystemInstance()&VehicleSystemInterfaceMask |
		(c.ecuName.GetIndustryGroup()&IndustryGroupMask)<<4 |
		byte(capable<<7)

	return nil
}
