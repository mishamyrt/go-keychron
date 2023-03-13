package keychron

type Features uint8

func (f Features) Has(flag Features) bool {
	return f&flag != 0
}

const (
	SpecificColor Features = 1 << iota
	RandomColor
	DirectionHorizontal
	DirectionVertical
)

type Effect struct {
	name  string
	value byte
	flags Features
}

var Effects = []Effect{
	{
		"Static",
		StaticModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Keystroke light up",
		KeystrokeLightUpModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Keystroke dim",
		KeystrokeDimModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Sparkle",
		SparkleModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Rain",
		RainModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Random colors",
		RandomColorsModeValue,
		RandomColor,
	},
	{
		"Breathing",
		BreathingModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Spectrum cycle",
		SpectrumCycleModeValue,
		RandomColor,
	},
	{
		"Ring gradient",
		RingGradientModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Vertical gradient",
		VerticalGradientModeValue,
		RandomColor | SpecificColor | DirectionVertical,
	},
	{
		"Horizontal gradient / Rainbow wave",
		HorizontalGradientWaveModeValue,
		RandomColor | SpecificColor | DirectionHorizontal,
	},
	{
		"Around edges",
		AroundEdgesModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Keystroke horizontal lines",
		KeystrokeHorizontalLinesValue,
		RandomColor | SpecificColor,
	},
	{
		"Keystroke tilted lines",
		KeystrokeTitledLinesModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Keystroke ripples",
		KeystrokeRipplesModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Sequence",
		SequenceModeValue,
		RandomColor | SpecificColor | DirectionHorizontal,
	},
	{
		"Wave line",
		WaveLineModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Tilted lines",
		TiltedLinesModeValue,
		RandomColor | SpecificColor,
	},
	{
		"Back and forth",
		BackAndForthModeValue,
		RandomColor | SpecificColor | DirectionHorizontal,
	},
}

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Mode struct {
	Color       Color
	RandomColor bool
	EffectValue uint8
	Brightness  uint8
	Speed       uint8
	Direction   uint8
}

func CreateModes() []Mode {
	modes := make([]Mode, len(Effects))
	for i := 0; i < len(Effects); i++ {
		e := Effects[i]
		modes[i].EffectValue = e.value
		if e.flags.Has(SpecificColor) {
			modes[i].Color = Color{
				Red:   255,
				Green: 255,
				Blue:  255,
			}
		}
		if e.flags.Has(DirectionHorizontal) {
			modes[i].Direction = DirectionLTR
		} else if e.flags.Has(DirectionVertical) {
			modes[i].Direction = DirectionUTD
		}

		modes[i].Speed = MaxSpeed
		modes[i].Brightness = MaxBrightness
	}
	return modes
}

// func mapBrightness(value uint8) uint8 {
// 	var output float64 = MinBrightness + (float64(MaxBrightness-MinBrightness)/255)*float64(value)
// 	return uint8(output)
// }

var Modes = CreateModes()
