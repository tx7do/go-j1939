package addressing

const (
	AddressClaimName   = "Address Claim"
	AddressClaimPgn    = 0x00EE00
	AddressFrameLength = 8

	ArbitraryAddrCapableMask   = 0x1
	ArbitraryAddrCapableOffset = 63

	IndustryGroupMask   = 0x7
	IndustryGroupOffset = 60

	VehicleSystemInterfaceMask   = 0xF
	VehicleSystemInterfaceOffset = 56
	VehicleSystemMask            = 0x7F
	VehicleSystemOffset          = 49

	FunctionMask           = 0xFF
	FunctionOffset         = 40
	FunctionInstanceMask   = 0x1F
	FunctionInstanceOffset = 35

	EcuInstanceMask   = 0x7
	EcuInstanceOffset = 32

	ManufacturerCodeMask   = 0x7FF
	ManufacturerCodeOffset = 21

	IdentityNumberMask   = 0x1FFFFF
	IdentityNumberOffset = 0
)
