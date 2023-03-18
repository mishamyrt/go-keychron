package keyboard

import (
	"github.com/mishamyrt/go-keychron/pkg/effect"
)

func fillEffect(m *effect.Mode, target []byte, offset int) {
	target[offset+OffsetCode] = m.Code

	if effect.IsRandomColor(m.Color) {
		target[offset+OffsetRandomColor] = 1
	} else {
		target[offset+OffsetR] = m.Color.R
		target[offset+OffsetG] = m.Color.G
		target[offset+OffsetB] = m.Color.B
	}

	target[offset+OffsetBrightness] = m.Brightness
	target[offset+OffsetSpeed] = m.Speed
	target[offset+OffsetDirection] = m.Direction.Code()

	target[offset+OffsetCRCLow] = EffectCRCLow
	target[offset+OffsetCRCHigh] = EffectCRCHigh
}

func parseEffects(buf []byte, count int) ([]effect.Mode, error) {
	var modes []effect.Mode = make([]effect.Mode, count)
	var err error
	for i := 0; i < count; i++ {
		offset := i * EffectPageLength
		if buf[offset+OffsetCRCLow] != EffectCRCLow || buf[offset+OffsetCRCHigh] != EffectCRCHigh {
			return modes, ErrCRCMismatch
		}

		modes[i], err = effect.Get(buf[offset+OffsetCode])
		if err != nil {
			return modes, err
		}

		if buf[offset+OffsetRandomColor] == 1 {
			modes[i].Color = effect.RandomColor
		} else {
			modes[i].Color.R = buf[offset+OffsetR]
			modes[i].Color.G = buf[offset+OffsetG]
			modes[i].Color.B = buf[offset+OffsetB]
		}

		modes[i].Brightness = buf[offset+OffsetBrightness]
		modes[i].Speed = buf[offset+OffsetSpeed]
		modes[i].Direction = effect.GetDirection(buf[offset+OffsetDirection])
	}
	return modes, nil
}
