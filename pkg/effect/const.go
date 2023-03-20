package effect

// Speed constants.
const (
	Slowest byte = 0x00
	Fastest byte = 0x0F
)

// Brightness constants.
const (
	// According to the documentation, the minimum brightness is 0x0,
	// but experimentally it was found that already at 0x01 the backlight turns off.
	Darkest   byte = 0x01
	Brightest byte = 0x0F
)
