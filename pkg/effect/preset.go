package effect

import colorPkg "image/color"

type Preset struct {
	mode       *Mode
	color      colorPkg.RGBA
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
func (p *Preset) Color() colorPkg.RGBA {
	return p.color
}

// SetColor sets color to preset
func (p *Preset) SetColor(c colorPkg.RGBA) error {
	if !p.mode.Features.Supports(SpecificColor) {
		return NewErrNotSupported("specific color")
	}
	if c.A != 0 && !IsRandomColor(c) {
		return NewErrNotSupported("color alpha channel")
	}
	p.color = c
	return nil
}

// SetRandomColor sets random color to preset
func (p *Preset) SetRandomColor() {
	p.color = RandomColorValue
}

// IsRandomColor checks if random color set to preset
func (p *Preset) IsRandomColor() bool {
	return IsRandomColor(p.color)
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

func NewPreset(mode *Mode, color colorPkg.RGBA, speed byte, direction EffectDirection) Preset {
	return Preset{
		mode:      mode,
		color:     color,
		speed:     speed,
		direction: direction,
	}
}
