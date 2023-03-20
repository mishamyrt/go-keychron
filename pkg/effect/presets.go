package effect

type PresetList []Preset

var (
	Static                   = NewPreset(&StaticMode, RGB(255, 255, 255), 0, 0)
	KeystrokeLightUp         = NewPreset(&KeystrokeLightUpMode, RandomColorValue, Slowest, 0)
	KeystrokeDim             = NewPreset(&KeystrokeDimMode, RandomColorValue, Slowest, 0)
	Sparkle                  = NewPreset(&SparkleMode, RandomColorValue, 5, 0)
	Rain                     = NewPreset(&RainMode, RandomColorValue, 2, 0)
	RandomColors             = NewPreset(&RandomColorsMode, RandomColorValue, 3, 0)
	Breathing                = NewPreset(&BreathingMode, RandomColorValue, 8, 0)
	SpectrumCycle            = NewPreset(&SpectrumCycleMode, RandomColorValue, Slowest, 0)
	RingGradient             = NewPreset(&RingGradientMode, RandomColorValue, 7, 0)
	VerticalGradient         = NewPreset(&VerticalGradientMode, RandomColorValue, 5, 0)
	HorizontalGradient       = NewPreset(&HorizontalGradientMode, RandomColorValue, 5, 0)
	AroundEdges              = NewPreset(&AroundEdgesMode, RandomColorValue, 5, 0)
	KeystrokeHorizontalLines = NewPreset(&KeystrokeHorizontalLinesMode, RandomColorValue, 8, 0)
	KeystrokeTiltedLines     = NewPreset(&KeystrokeTiltedLinesMode, RandomColorValue, 8, 0)
	KeystrokeRipples         = NewPreset(&KeystrokeRipplesMode, RandomColorValue, 4, 0)
	Sequence                 = NewPreset(&SequenceMode, RandomColorValue, 4, 0)
	WaveLine                 = NewPreset(&WaveLineMode, RandomColorValue, 6, 0)
	TiltedLines              = NewPreset(&TiltedLinesMode, RandomColorValue, 3, 0)
	BackAndForth             = NewPreset(&BackAndForthMode, RandomColorValue, 6, 0)
)

var Presets = PresetList{
	Static,
	KeystrokeLightUp,
	KeystrokeDim,
	Sparkle,
	Rain,
	RandomColors,
	Breathing,
	SpectrumCycle,
	RingGradient,
	VerticalGradient,
	HorizontalGradient,
	AroundEdges,
	KeystrokeHorizontalLines,
	KeystrokeTiltedLines,
	KeystrokeRipples,
	Sequence,
	WaveLine,
	TiltedLines,
	BackAndForth,
}
