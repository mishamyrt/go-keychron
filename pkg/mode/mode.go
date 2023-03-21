package mode

// Mode represents Keychron keyboard backlight mode
type Mode struct {
	Name     string
	Code     byte
	Features Feature
}

var (
	Static = Mode{
		"Static",
		StaticCode,
		SpecificColor | RandomColor,
	}
	KeystrokeLightUp = Mode{
		"Keystroke light up",
		KeystrokeLightUpCode,
		SpecificColor | RandomColor,
	}
	KeystrokeDim = Mode{
		"Keystroke dim",
		KeystrokeDimCode,
		SpecificColor | RandomColor,
	}
	Sparkle = Mode{
		"Sparkle",
		SparkleCode,
		SpecificColor | RandomColor,
	}
	Rain = Mode{
		"Rain",
		RainCode,
		SpecificColor | RandomColor,
	}
	RandomColors = Mode{
		"Random colors",
		RandomColorsCode,
		RandomColor,
	}
	Breathing = Mode{
		"Breathing",
		BreathingCode,
		SpecificColor | RandomColor,
	}
	SpectrumCycle = Mode{
		"Spectrum cycle",
		SpectrumCycleCode,
		RandomColor,
	}
	RingGradient = Mode{
		"Ring gradient",
		RingGradientCode,
		SpecificColor | RandomColor | VerticalDirection,
	}
	VerticalGradient = Mode{
		"Vertical gradient",
		VerticalGradientCode,
		SpecificColor | RandomColor | VerticalDirection,
	}
	HorizontalGradient = Mode{
		"Horizontal gradient",
		HorizontalGradientCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	AroundEdges = Mode{
		"Around edges",
		AroundEdgesCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	KeystrokeHorizontalLines = Mode{
		"Keystroke horizontal lines",
		KeystrokeHorizontalLinesCode,
		SpecificColor | RandomColor,
	}
	KeystrokeTiltedLines = Mode{
		"Keystroke tilted lines",
		KeystrokeTitledLinesCode,
		SpecificColor | RandomColor,
	}
	KeystrokeRipples = Mode{
		"Keystroke ripples",
		KeystrokeRipplesCode,
		SpecificColor | RandomColor,
	}
	Sequence = Mode{
		"Sequence",
		SequenceCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	WaveLine = Mode{
		"Wave line",
		WaveLineCode,
		SpecificColor | RandomColor,
	}
	TiltedLines = Mode{
		"Tilted lines",
		TiltedLinesCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
	BackAndForth = Mode{
		"Back and forth",
		BackAndForthCode,
		SpecificColor | RandomColor | HorizontalDirection,
	}
)
