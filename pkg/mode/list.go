package mode

var List = []Mode{
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

func ByName(n string) (*Mode, error) {
	for i := range List {
		if List[i].Name == n {
			return &List[i], nil
		}
	}
	return nil, ErrNotFound
}

func ByCode(c byte) (*Mode, error) {
	for i := range List {
		if List[i].Code == c {
			return &List[i], nil
		}
	}
	return nil, ErrNotFound
}
