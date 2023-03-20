package effect

var Presets = []Preset{}

var Static = NewPreset(&StaticMode, RandomColorValue, 0x5, Brightest, 0)
var KeystrokeLightUp = NewPreset(&StaticMode, RandomColorValue, 0x5, Brightest, 0)

// StaticMode,
// KeystrokeLightUpMode,
// KeystrokeDimMode,
// SparkleMode,
// RainMode,
// RandomColorsMode,
// BreathingMode,
// SpectrumCycleMode,
// RingGradientMode,
// VerticalGradientMode,
// HorizontalGradientMode,
// AroundEdgesMode,
// KeystrokeHorizontalLinesMode,
// KeystrokeTiltedLinesMode,
// KeystrokeRipplesMode,
// SequenceMode,
// WaveLineMode,
// TiltedLinesMode,
// BackAndForthMode,
