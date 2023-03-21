package preset

// Speed constants.
const (
	SpeedMin uint8 = 0
	SpeedMax uint8 = 15
)

// Brightness constants.
const (
	// According to the documentation, the minimum brightness is 0x0,
	// but experimentally it was found that already at 0x01 the backlight turns off.
	BrightnessMin byte = 1
	BrightnessMax byte = 15
)
