package effect

// Speed constants.
const (
	Slowest byte = 0x00
	Fastest byte = 0x0F
)

// Brightness constants.
const (
	// According to the documentation, the minimum brightness is 0x0,
	// but experimentally it was found that already at 0x01 the backlight turns off.
	Darkest   byte = 0x01
	Brightest byte = 0x0F
)

// Mode code constants.
const (
	CustomMode                   byte = 0x00
	StaticMode                   byte = 0x01
	KeystrokeLightUpMode         byte = 0x02
	KeystrokeDimMode             byte = 0x03
	SparkleMode                  byte = 0x04
	RainMode                     byte = 0x05
	RandomColorsMode             byte = 0x06
	BreathingMode                byte = 0x07
	SpectrumCycleMode            byte = 0x08
	RingGradientMode             byte = 0x09
	VerticalGradientMode         byte = 0x0A
	HorizontalGradientWaveMode   byte = 0x0B
	AroundEdgesMode              byte = 0x0C
	KeystrokeHorizontalLinesMode byte = 0x0D
	KeystrokeTitledLinesMode     byte = 0x0E
	KeystrokeRipplesMode         byte = 0x0F
	SequenceMode                 byte = 0x10
	WaveLineMode                 byte = 0x11
	TiltedLinesMode              byte = 0x12
	BackAndForthMode             byte = 0x13
	LightsOffMode                byte = 0x80
)
