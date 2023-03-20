package effect

import "image/color"

const randomValue = 0x55

// RandomColorValue represents a special color that, when applied,
// will make the colors of the effect random.
var RandomColorValue = color.RGBA{
	A: randomValue,
}

// IsRandomColor checks if the color is random.
func IsRandomColor(c color.RGBA) bool {
	return c.R+c.G+c.B == 0 && c.A == randomValue
}

// RGB creates rgb color
func RGB(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 0}
}
