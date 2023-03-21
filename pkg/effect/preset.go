package effect

import "github.com/mishamyrt/go-keychron/pkg/color"

type Preset struct {
	mode       *Mode
	color      color.RGB
	speed      byte
	brightness byte
	direction  EffectDirection
}

// Mode returns preset mode
func (p *Preset) Mode() Mode {
	return *p.mode
}

// SetModeByCode sets preset mode
func (p *Preset) SetMode(m *Mode) {
	p.mode = m
}

// SetModeByCode sets preset mode by code
func (p *Preset) SetModeByCode(c byte) error {
	m, err := Modes.GetByCode(c)
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
	if !p.mode.Features.Supports(SpecificColor) {
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
	if s > Fastest || s < Slowest {
		return NewErrOutOfRange(s, Slowest, Fastest)
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
	if b > Brightest || b < Darkest {
		return NewErrOutOfRange(b, Darkest, Brightest)
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
	if d.IsVertical() && !p.mode.Features.Supports(VerticalDirection) {
		return NewErrNotSupported("vertical direction")
	} else if !p.mode.Features.Supports(HorizontalDirection) {
		return NewErrNotSupported("horizontal direction")
	}
	p.direction = d
	return nil
}

func NewPreset(mode *Mode, color color.RGB, speed uint8, direction EffectDirection) Preset {
	return Preset{
		mode:       mode,
		color:      color,
		speed:      speed,
		brightness: Brightest,
		direction:  direction,
	}
}
