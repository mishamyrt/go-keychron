package keychron

const ReportID = 0x00
const VendorID = 0x05AC

const PacketLength = 64
const EffectPageLength = 16
const ColorBufferLength = 576

const DirectionLTR = 0x00
const DirectionRTL = 0x01
const DirectionDTU = 0x02
const DirectionUTD = 0x03

const MinSpeed = 0x00
const MaxSpeed = 0x0F

const MinBrightness = 0x00
const MaxBrightness = 0x0F

const EffectPageCRCLow = 0xAA
const EffectPageCRCHigh = 0x55

const PacketHeader = 0x04
const LEDSpecialEffectPacket = 0x12

const CustomModeValue = 0x00
const StaticModeValue = 0x01
const KeystrokeLightUpModeValue = 0x02
const KeystrokeDimModeValue = 0x03
const SparkleModeValue = 0x04
const RainModeValue = 0x05
const RandomColorsModeValue = 0x06
const BreathingModeValue = 0x07
const SpectrumCycleModeValue = 0x08
const RingGradientModeValue = 0x09
const VerticalGradientModeValue = 0x0A
const HorizontalGradientWaveModeValue = 0x0B
const AroundEdgesModeValue = 0x0C
const KeystrokeHorizontalLinesValue = 0x0D
const KeystrokeTitledLinesModeValue = 0x0E
const KeystrokeRipplesModeValue = 0x0F
const SequenceModeValue = 0x10
const WaveLineModeValue = 0x11
const TiltedLinesModeValue = 0x12
const BackAndForthModeValue = 0x13
const LightsOffModeValue = 0x80

const CommunicationEndCommand = 0x02
const WriteLEDSpecialEffectAreaCommand = 0x13
const TurnOnCustomizationCommand = 0x18
const TurnOffCustomizationCommand = 0x19
const LEDEffectStartCommand = 0xF0
