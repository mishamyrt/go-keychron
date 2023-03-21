package preset

import "github.com/mishamyrt/go-keychron/pkg/mode"

type PresetList []Preset

var List = PresetList{
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

func For(m *mode.Mode) *Preset {
	for i := range List {
		if List[i].mode.Code == m.Code {
			return &List[i]
		}
	}
	return nil
}
