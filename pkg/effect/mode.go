package effect

type Feature byte

func (f Feature) Supports(x Feature) bool {
	return f&x != 0
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
