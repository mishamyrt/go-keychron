package color

type RGB struct {
	R, G, B uint8
	random  bool
}

func (c *RGB) IsRandom() bool {
	return c.random
}

// Random represents a special color that, when applied,
// will make the colors of the effect random.
var Random = RGB{
	random: true,
}

func New(R, G, B uint8) RGB {
	return RGB{R, G, B, false}
}
