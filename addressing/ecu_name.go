package addressing

type EcuName struct {
	IdNumber                uint32
	ManufacturerCode        uint16
	EcuInstance             uint8
	FunctionInstance        uint8
	Function                uint8
	VehicleSystem           uint8
	VehicleSystemInstance   uint8
	IndustryGroup           uint8
	ArbitraryAddressCapable bool
}

func NewEcuName() *EcuName {
	c := &EcuName{}
	c.IdNumber = 0
	c.ManufacturerCode = 0
	c.EcuInstance = 0
	c.FunctionInstance = 0
	c.Function = 0
	c.VehicleSystem = 0
	c.VehicleSystemInstance = 0
	c.IndustryGroup = 0
	c.ArbitraryAddressCapable = false
	return c
}
func NewEcuNameWithValue(idNumber uint32, manufacturerCode uint16,
	ecuInstance, functionInstance, function, vehicleSystem, vehicleSystemInstance, industryGroup uint8,
	arbitraryAddressCapable bool) *EcuName {
	c := &EcuName{}
	c.IdNumber = idNumber
	c.ManufacturerCode = manufacturerCode
	c.EcuInstance = ecuInstance
	c.FunctionInstance = functionInstance
	c.Function = function
	c.VehicleSystem = vehicleSystem
	c.VehicleSystemInstance = vehicleSystemInstance
	c.IndustryGroup = industryGroup
	c.ArbitraryAddressCapable = arbitraryAddressCapable
	return c
}

func (c *EcuName) GetValue() uint64 {
	capable := 0
	if c.ArbitraryAddressCapable {
		capable = ArbitraryAddrCapableMask
	}

	return (uint64)(capable<<ArbitraryAddrCapableOffset) |
		((uint64)(c.IndustryGroup&IndustryGroupMask) << IndustryGroupOffset) |
		((uint64)(c.VehicleSystemInstance&VehicleSystemInterfaceMask) << VehicleSystemInterfaceOffset) |
		((uint64)(c.VehicleSystem&VehicleSystemMask) << VehicleSystemOffset) |
		((uint64)(c.Function&FunctionMask) << FunctionOffset) |
		((uint64)(c.FunctionInstance&FunctionInstanceMask) << FunctionInstanceOffset) |
		((uint64)(c.EcuInstance&EcuInstanceMask) << EcuInstanceOffset) |
		((uint64)(c.ManufacturerCode&ManufacturerCodeMask) << ManufacturerCodeOffset) |
		((uint64)(c.IdNumber&IdentityNumberMask) << IdentityNumberOffset)
}

func (c *EcuName) IsArbitraryAddressCapable() bool {
	return c.ArbitraryAddressCapable
}

func (c *EcuName) SetArbitraryAddressCapable(arbitraryAddressCapable bool) {
	c.ArbitraryAddressCapable = arbitraryAddressCapable
}

func (c *EcuName) GetEcuInstance() uint8 {
	return c.EcuInstance
}

func (c *EcuName) SetEcuInstance(ecuInstance uint8) {
	c.EcuInstance = ecuInstance
}

func (c *EcuName) GetFunction() uint8 {
	return c.Function
}

func (c *EcuName) SetFunction(function uint8) {
	c.Function = function
}

func (c *EcuName) GetFunctionInstance() uint8 {
	return c.FunctionInstance
}

func (c *EcuName) SetFunctionInstance(functionInstance uint8) {
	c.FunctionInstance = functionInstance
}

func (c *EcuName) GetIdNumber() uint32 {
	return c.IdNumber
}

func (c *EcuName) SetIdNumber(idNumber uint32) {
	c.IdNumber = idNumber
}

func (c *EcuName) GetIndustryGroup() uint8 {
	return c.IndustryGroup
}

func (c *EcuName) SetIndustryGroup(industryGroup uint8) {
	c.IndustryGroup = industryGroup
}

func (c *EcuName) GetManufacturerCode() uint16 {
	return c.ManufacturerCode
}

func (c *EcuName) SetManufacturerCode(manufacturerCode uint16) {
	c.ManufacturerCode = manufacturerCode
}

func (c *EcuName) GetVehicleSystem() uint8 {
	return c.VehicleSystem
}

func (c *EcuName) SetVehicleSystem(vehicleSystem uint8) {
	c.VehicleSystem = vehicleSystem
}

func (c *EcuName) GetVehicleSystemInstance() uint8 {
	return c.VehicleSystemInstance
}

func (c *EcuName) SetVehicleSystemInstance(vehicleSystemInstance uint8) {
	c.VehicleSystemInstance = vehicleSystemInstance
}
