package effect

type ModeList []Mode

func (m ModeList) GetByName(n string) (*Mode, error) {
	for i := range m {
		if m[i].Name == n {
			mode := &m[i]
			return mode, nil
		}
	}
	return nil, ErrNotFound
}

func (m ModeList) GetByCode(c byte) (*Mode, error) {
	for i := range m {
		if m[i].Code == c {
			mode := &m[i]
			return mode, nil
		}
	}
	return nil, ErrNotFound
}

// Mode code constants.
const (
	CustomCode                   byte = 0x00
	StaticCode                   byte = 0x01
	KeystrokeLightUpCode         byte = 0x02
	KeystrokeDimCode             byte = 0x03
	SparkleCode                  byte = 0x04
	RainCode                     byte = 0x05
	RandomColorsCode             byte = 0x06
	BreathingCode                byte = 0x07
	SpectrumCycleCode            byte = 0x08
	RingGradientCode             byte = 0x09
	VerticalGradientCode         byte = 0x0A
	HorizontalGradientCode       byte = 0x0B
	AroundEdgesCode              byte = 0x0C
	KeystrokeHorizontalLinesCode byte = 0x0D
	KeystrokeTitledLinesCode     byte = 0x0E
	KeystrokeRipplesCode         byte = 0x0F
	SequenceCode                 byte = 0x10
	WaveLineCode                 byte = 0x11
	TiltedLinesCode              byte = 0x12
	BackAndForthCode             byte = 0x13
	LightsOffCode                byte = 0x80
)

var (
	StaticMode = Mode{
		"Static",
		StaticCode,
		SpecificColor | RandomColor,
	}
	KeystrokeLightUpMode = Mode{
		"Keystroke light up",
		KeystrokeLightUpCode,
		SpecificColor | RandomColor,
	}
	KeystrokeDimMode = Mode{
		"Keystroke dim",
		KeystrokeDimCode,
		SpecificColor | RandomColor,
	}
	SparkleMode = Mode{
		"Sparkle",
		SparkleCode,
		SpecificColor | RandomColor,
	}
	RainMode = Mode{
		"Rain",
		RainCode,
		SpecificColor | RandomColor,
	}
	RandomColorsMode = Mode{
		"Random colors",
		RandomColorsCode,
		RandomColor,
	}
	BreathingMode = Mode{
		"Breathing",
		BreathingCode,
		SpecificColor | RandomColor,
	}
	SpectrumCycleMode = Mode{
		"Spectrum cycle",
		SpectrumCycleCode,
		RandomColor,
	}
	RingGradientMode = Mode{
		"Ring gradient",
		RingGradientCode,
		SpecificColor | RandomColor | VerticalDirection,
	}
	VerticalGradientMode = Mode{
		"Vertical gradient",
		VerticalGradientCode,
		SpecificColor | RandomColor | VerticalDirection,
	}
	HorizontalGradientMode = Mode{
		"Horizontal gradient",
		HorizontalGradientCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	AroundEdgesMode = Mode{
		"Around edges",
		AroundEdgesCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	KeystrokeHorizontalLinesMode = Mode{
		"Keystroke horizontal lines",
		KeystrokeHorizontalLinesCode,
		SpecificColor | RandomColor,
	}
	KeystrokeTiltedLinesMode = Mode{
		"Keystroke tilted lines",
		KeystrokeTitledLinesCode,
		SpecificColor | RandomColor,
	}
	KeystrokeRipplesMode = Mode{
		"Keystroke ripples",
		KeystrokeRipplesCode,
		SpecificColor | RandomColor,
	}
	SequenceMode = Mode{
		"Sequence",
		SequenceCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	WaveLineMode = Mode{
		"Wave line",
		WaveLineCode,
		SpecificColor | RandomColor,
	}
	TiltedLinesMode = Mode{
		"Tilted lines",
		TiltedLinesCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	BackAndForthMode = Mode{
		"Back and forth",
		BackAndForthCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
)

var Modes = ModeList{
	StaticMode,
	KeystrokeLightUpMode,
	KeystrokeDimMode,
	SparkleMode,
	RainMode,
	RandomColorsMode,
	BreathingMode,
	SpectrumCycleMode,
	RingGradientMode,
	VerticalGradientMode,
	HorizontalGradientMode,
	AroundEdgesMode,
	KeystrokeHorizontalLinesMode,
	KeystrokeTiltedLinesMode,
	KeystrokeRipplesMode,
	SequenceMode,
	WaveLineMode,
	TiltedLinesMode,
	BackAndForthMode,
}
