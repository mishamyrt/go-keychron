package keychron

import "github.com/mishamyrt/go-keychron/pkg/effect"

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
