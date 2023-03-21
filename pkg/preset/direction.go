package preset

// EffectDirection is the type used to define the direction of the effect.
type EffectDirection byte

// Direction constants.
const (
	LeftToRight EffectDirection = iota
	RightToLeft
	BottomToUp
	TopToBottom
)

func (d EffectDirection) String() string {
	switch d {
	case LeftToRight:
		return "left-to-right"
	case RightToLeft:
		return "right-to-left"
	case BottomToUp:
		return "bottom-to-top"
	case TopToBottom:
		return "top-to-bottom"
	}
	return "<unknown>"
}

func (d EffectDirection) Code() byte {
	switch d {
	case LeftToRight:
		return 0x00
	case RightToLeft:
		return 0x01
	case BottomToUp:
		return 0x02
	case TopToBottom:
		return 0x03
	}
	return 0xFF
}

func (d EffectDirection) IsHorizontal() bool {
	return d == LeftToRight || d == RightToLeft
}

func (d EffectDirection) IsVertical() bool {
	return d == TopToBottom || d == BottomToUp
}

func GetDirection(v byte) EffectDirection {
	switch v {
	case 0x01:
		return RightToLeft
	case 0x02:
		return BottomToUp
	case 0x03:
		return TopToBottom
	default:
		return LeftToRight
	}
}
