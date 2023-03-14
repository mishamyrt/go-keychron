package effect

import "image/color"

const randomValue = 0x55

// RandomColor represents a special color that, when applied,
// will make the colors of the effect random.
var RandomColor = color.RGBA{
	A: randomValue,
}

// IsRandomColor checks if the color is random.
func IsRandomColor(c color.RGBA) bool {
	return c.R == 0 && c.G == 0 && c.B == 0 && c.A == randomValue
}
