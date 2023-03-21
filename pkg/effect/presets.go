package effect

import "github.com/mishamyrt/go-keychron/pkg/color"

type PresetList []Preset

// Paginate splits the list of presets into pages of a set number.
func (p PresetList) Paginate(n int) []PresetList {
	count := len(p)
	size := count / n
	if size%n > 0 {
		size += 1
	}
	// 0:4 – 0, 1, 2, 3
	// 4:8

	// 2:4 – 2,3,4
	pages := make([]PresetList, size)
	for i := 0; i < size; i++ {
		start := i * size
		end := start + size
		if end > count {
			end = count
		}
		pages[i] = p[start:end]
	}
	return pages
}

var (
	Static                   = NewPreset(&StaticMode, color.New(255, 255, 255), 0, 0)
	KeystrokeLightUp         = NewPreset(&KeystrokeLightUpMode, color.Random, Slowest, 0)
	KeystrokeDim             = NewPreset(&KeystrokeDimMode, color.Random, Slowest, 0)
	Sparkle                  = NewPreset(&SparkleMode, color.Random, 5, 0)
	Rain                     = NewPreset(&RainMode, color.Random, 2, 0)
	RandomColors             = NewPreset(&RandomColorsMode, color.Random, 3, 0)
	Breathing                = NewPreset(&BreathingMode, color.Random, 8, 0)
	SpectrumCycle            = NewPreset(&SpectrumCycleMode, color.Random, Slowest, 0)
	RingGradient             = NewPreset(&RingGradientMode, color.Random, 7, 0)
	VerticalGradient         = NewPreset(&VerticalGradientMode, color.Random, 5, 0)
	HorizontalGradient       = NewPreset(&HorizontalGradientMode, color.Random, 5, 0)
	AroundEdges              = NewPreset(&AroundEdgesMode, color.Random, 5, 0)
	KeystrokeHorizontalLines = NewPreset(&KeystrokeHorizontalLinesMode, color.Random, 8, 0)
	KeystrokeTiltedLines     = NewPreset(&KeystrokeTiltedLinesMode, color.Random, 8, 0)
	KeystrokeRipples         = NewPreset(&KeystrokeRipplesMode, color.Random, 4, 0)
	Sequence                 = NewPreset(&SequenceMode, color.Random, 4, 0)
	WaveLine                 = NewPreset(&WaveLineMode, color.Random, 6, 0)
	TiltedLines              = NewPreset(&TiltedLinesMode, color.Random, 3, 0)
	BackAndForth             = NewPreset(&BackAndForthMode, color.Random, 6, 0)
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
