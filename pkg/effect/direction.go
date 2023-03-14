package effect

// EffectDirection is the type used to define the direction of the effect.
type EffectDirection uint8

// Direction constants.
const (
	LeftToRight EffectDirection = iota
	RightToLeft
	DownToUp
	UpToDown
)

func (s EffectDirection) Code() byte {
	switch s {
	case LeftToRight:
		return 0x00
	case RightToLeft:
		return 0x01
	case DownToUp:
		return 0x02
	case UpToDown:
		return 0x03
	}
	return 0xFF
}
