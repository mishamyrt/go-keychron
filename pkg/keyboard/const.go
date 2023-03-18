package keyboard

const PacketHeader = 0x04
const CustomColorHeader = 0x80

const (
	CommunicationEnd     = 0x02
	ReadEffects          = 0x12
	WriteLEDEffects      = 0x13
	TurnOnCustomization  = 0x18
	TurnOffCustomization = 0x19
	ApplyLEDEffects      = 0xF0
)

const (
	CmdACK  = 0x01
	CmdNACK = 0xFF
)

const (
	EffectPages      = 18
	EffectPageLength = 16
	EffectCRCLow     = 0xAA
	EffectCRCHigh    = 0x55
)

const (
	OffsetCode        = 0
	OffsetR           = 1
	OffsetG           = 2
	OffsetB           = 3
	OffsetRandomColor = 8
	OffsetBrightness  = 9
	OffsetSpeed       = 10
	OffsetDirection   = 11
	OffsetCRCLow      = 14
	OffsetCRCHigh     = 15
)

const ColorBufferLength = 576
