package j1939

const (
	MaxSize uint32 = 8

	PgnOffset uint32 = 8
	PgnMask   uint32 = 0x3FFFF

	PduSpecificMask uint32 = 0xFF

	SrcAddrMask uint32 = 0xFF
	DstAddrMask uint32 = 0xFF

	SrcAddrOffset uint32 = 0
	DstAddrOffset uint32 = 0

	PduFmtMask   uint32 = 0xFF
	PduFmtOffset uint32 = 8

	DataPageMask   uint32 = 1
	DataPageOffset uint32 = 16

	ExtDataPageMask   uint32 = 1
	ExtDataPageOffset uint32 = 17

	PriorityOffset uint32 = 26
	PriorityMask   uint32 = 7

	StatusMask uint32 = 3

	StringTerminator byte  = '*'
	NULL_TERMINATOR  uint8 = 0

	TpDtPacketSize uint8 = 7

	PduFmtDelimiter uint32 = 0xF0

	InvalidAddress   uint32 = 0xFE
	BroadcastAddress uint32 = 0xFF

	SPN_NUMBER_MAX_BITS uint8 = 19

	SpnNumericMaxByteSize uint8 = 4
)
