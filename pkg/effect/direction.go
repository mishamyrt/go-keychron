package effect

// EffectDirection is the type used to define the direction of the effect.
type EffectDirection byte

// Direction constants.
const (
	LeftToRight EffectDirection = iota
	RightToLeft
	DownToUp
	UpToDown
)

func (d EffectDirection) Code() byte {
	switch d {
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

func GetDirection(v byte) EffectDirection {
	switch v {
	case 0x01:
		return RightToLeft
	case 0x02:
		return DownToUp
	case 0x03:
		return UpToDown
	default:
		return LeftToRight
	}
}
