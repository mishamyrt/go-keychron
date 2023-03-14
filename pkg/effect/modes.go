package effect

import (
	"errors"
	"image/color"
)

// Mode represents Keychron keyboard backlight mode
type Mode struct {
	// Human readable mode name.
	Name string
	// Mode code. One of the effect code constants.
	Code byte
	// Mode color. The Alpha channel is ignored,
	// except when its value is 0x55 (85) with all other values equal to 0.
	Color color.RGBA
	// Mode speed. Number in the range from MinSpeed to MaxSpeed.
	Speed byte
	// Mode brightness. Number in the range from MinBrightness to MaxBrightness.
	Brightness byte
	// Mode direction. Value from direction constants.
	Direction EffectDirection
}

// Modes available on Keychron keyboards
var Modes = []Mode{
	{
		Name:       "Static",
		Code:       StaticMode,
		Color:      RandomColor,
		Brightness: Brightest,
	},
	{
		Name:       "Keystroke light up",
		Code:       KeystrokeLightUpMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      Slowest,
	},
	{
		Name:       "Keystroke dim",
		Code:       KeystrokeDimMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      Slowest,
	},
	{
		Name:       "Sparkle",
		Code:       SparkleMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      10,
	},
	{
		Name:       "Rain",
		Code:       RainMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Direction:  UpToDown,
		Speed:      Slowest,
	},
	{
		Name:       "Random colors",
		Code:       RandomColorsMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      Slowest,
	},
	{
		Name:       "Breathing",
		Code:       BreathingMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      7,
	},
	{
		Name:       "Spectrum cycle",
		Code:       SpectrumCycleMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      Slowest,
	},
	{
		Name:       "Ring gradient",
		Code:       RingGradientMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      Slowest,
		Direction:  LeftToRight,
	},
	{
		Name:       "Vertical gradient",
		Code:       VerticalGradientMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      5,
		Direction:  UpToDown,
	},
	{
		Name:       "Horizontal gradient",
		Code:       HorizontalGradientWaveMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      5,
		Direction:  LeftToRight,
	},
	{
		Name:       "Around edges",
		Code:       AroundEdgesMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      5,
		Direction:  LeftToRight,
	},
	{
		Name:       "Keystroke horizontal lines",
		Code:       KeystrokeHorizontalLinesMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      5,
	},
	{
		Name:       "Keystroke tilted lines",
		Code:       KeystrokeHorizontalLinesMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      9,
	},
	{
		Name:       "Keystroke ripples",
		Code:       KeystrokeRipplesMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      10,
	},
	{
		Name:       "Sequence",
		Code:       SequenceMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      Slowest,
		Direction:  LeftToRight,
	},
	{
		Name:       "Wave line",
		Code:       WaveLineMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      5,
	},
	{
		Name:       "Tilted lines",
		Code:       TiltedLinesMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      5,
	},
	{
		Name:       "Back and forth",
		Code:       BackAndForthMode,
		Color:      RandomColor,
		Brightness: Brightest,
		Speed:      5,
		Direction:  LeftToRight,
	},
}

// ErrModeNotFound means that the mode is not found.
var ErrModeNotFound = errors.New("mode with desired code is not found")

// Finds the mode by code. If the mode is not found, it returns an ErrModeNotFound error.
func Get(code byte) (Mode, error) {
	for i := range Modes {
		if Modes[i].Code == code {
			return Modes[i], nil
		}
	}
	return Mode{}, ErrModeNotFound
}
