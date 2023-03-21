package preset

import (
	"github.com/mishamyrt/go-keychron/pkg/color"
	"github.com/mishamyrt/go-keychron/pkg/mode"
)

type Preset struct {
	mode       *mode.Mode
	color      color.RGB
	speed      byte
	brightness byte
	direction  EffectDirection
}

// Mode returns preset mode
func (p *Preset) Mode() mode.Mode {
	return *p.mode
}

// SetModeByCode sets preset mode
func (p *Preset) SetMode(m *mode.Mode) {
	p.mode = m
}

// SetModeByCode sets preset mode by code
func (p *Preset) SetModeByCode(c byte) error {
	m, err := mode.ByCode(c)
	if err != nil {
		return err
	}
	p.mode = m
	return nil
}

// Color returns preset mode
func (p *Preset) Color() color.RGB {
	return p.color
}

// SetColor sets color to preset
func (p *Preset) SetColor(c color.RGB) error {
	if !p.mode.Features.Supports(mode.SpecificColor) {
		return NewErrNotSupported("specific color")
	}
	p.color = c
	return nil
}

// SetRandomColor sets random color to preset
func (p *Preset) SetRandomColor() {
	p.color = color.Random
}

// IsRandomColor checks if random color set to preset
func (p *Preset) IsRandomColor() bool {
	return p.color.IsRandom()
}

// Speed returns preset speed
func (p *Preset) Speed() byte {
	return p.speed
}

// SetSpeed sets preset speed
func (p *Preset) SetSpeed(s byte) error {
	if s < SpeedMin || s > SpeedMax {
		return NewErrOutOfRange(s, SpeedMin, SpeedMax)
	}
	p.speed = s
	return nil
}

// Brightness returns preset brightness
func (p *Preset) Brightness() byte {
	return p.brightness
}

// SetBrightness sets preset brightness
func (p *Preset) SetBrightness(b byte) error {
	if b < BrightnessMin || b > BrightnessMax {
		return NewErrOutOfRange(b, BrightnessMin, BrightnessMax)
	}
	p.brightness = b
	return nil
}

// Direction returns preset direction
func (p *Preset) Direction() EffectDirection {
	return p.direction
}

// SetDirection sets preset direction
func (p *Preset) SetDirection(d EffectDirection) error {
	if d.IsVertical() && !p.mode.Features.Supports(mode.VerticalDirection) {
		return NewErrNotSupported("vertical direction")
	} else if !p.mode.Features.Supports(mode.HorizontalDirection) {
		return NewErrNotSupported("horizontal direction")
	}
	p.direction = d
	return nil
}

func New(mode *mode.Mode, color color.RGB, speed uint8, direction EffectDirection) Preset {
	return Preset{
		mode:       mode,
		color:      color,
		speed:      speed,
		brightness: BrightnessMax,
		direction:  direction,
	}
}

var (
	Static                   = New(&mode.Static, color.New(255, 255, 255), 0, 0)
	KeystrokeLightUp         = New(&mode.KeystrokeLightUp, color.Random, SpeedMin, 0)
	KeystrokeDim             = New(&mode.KeystrokeDim, color.Random, SpeedMin, 0)
	Sparkle                  = New(&mode.Sparkle, color.Random, 5, 0)
	Rain                     = New(&mode.Rain, color.Random, 2, 0)
	RandomColors             = New(&mode.RandomColors, color.Random, 3, 0)
	Breathing                = New(&mode.Breathing, color.Random, 8, 0)
	SpectrumCycle            = New(&mode.SpectrumCycle, color.Random, SpeedMin, 0)
	RingGradient             = New(&mode.RingGradient, color.Random, 7, 0)
	VerticalGradient         = New(&mode.VerticalGradient, color.Random, 5, 0)
	HorizontalGradient       = New(&mode.HorizontalGradient, color.Random, 5, 0)
	AroundEdges              = New(&mode.AroundEdges, color.Random, 5, 0)
	KeystrokeHorizontalLines = New(&mode.KeystrokeHorizontalLines, color.Random, 8, 0)
	KeystrokeTiltedLines     = New(&mode.KeystrokeTiltedLines, color.Random, 8, 0)
	KeystrokeRipples         = New(&mode.KeystrokeRipples, color.Random, 4, 0)
	Sequence                 = New(&mode.Sequence, color.Random, 4, 0)
	WaveLine                 = New(&mode.WaveLine, color.Random, 6, 0)
	TiltedLines              = New(&mode.TiltedLines, color.Random, 3, 0)
	BackAndForth             = New(&mode.BackAndForth, color.Random, 6, 0)
)
