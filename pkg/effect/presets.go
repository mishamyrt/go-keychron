package effect

var Presets = []Preset{}

var Static = NewPreset(&StaticMode, RandomColorValue, 0, 0)
var KeystrokeLightUp = NewPreset(&KeystrokeLightUpMode, RandomColorValue, Slowest, 0)
var KeystrokeDim = NewPreset(&KeystrokeLightUpMode, RandomColorValue, Slowest, 0)
var Sparkle = NewPreset(&SparkleMode, RandomColorValue, Slowest, 0)

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
