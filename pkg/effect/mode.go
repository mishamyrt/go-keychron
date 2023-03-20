package effect

type Feature byte

type FeaturePredicate func(f Feature) bool

func (f Feature) assert(features []Feature, assert FeaturePredicate, defaultExit bool) bool {
	for i := range features {
		if assert(features[i]) {
			return !defaultExit
		}
	}
	return defaultExit
}

func (f Feature) Supports(features ...Feature) bool {
	return f.assert(features, func(feature Feature) bool {
		return f&feature == 0
	}, true)
}

func (f Feature) SupportsAny(features ...Feature) bool {
	return f.assert(features, func(feature Feature) bool {
		return f&feature != 0
	}, false)
}

const (
	SpecificColor Feature = 1 << iota
	RandomColor
	VerticalDirection
	HorizontalDirection
	Speed
)

// Mode represents Keychron keyboard backlight mode
type Mode struct {
	Name     string
	Code     byte
	Features Feature
}
