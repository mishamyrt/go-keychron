package keyboard

import (
	"github.com/mishamyrt/go-keychron/pkg/color"
	"github.com/mishamyrt/go-keychron/pkg/mode"
	"github.com/mishamyrt/go-keychron/pkg/preset"
)

type EffectPage preset.PresetList

func fillPreset(p *preset.Preset, target []byte, offset int) {
	target[offset+OffsetCode] = p.Mode().Code

	if p.IsRandomColor() {
		target[offset+OffsetRandomColor] = 1
	} else {
		c := p.Color()
		target[offset+OffsetR] = c.R
		target[offset+OffsetG] = c.G
		target[offset+OffsetB] = c.B
	}

	target[offset+OffsetBrightness] = p.Brightness()
	target[offset+OffsetSpeed] = p.Speed()
	target[offset+OffsetDirection] = p.Direction().Code()

	target[offset+OffsetCRCLow] = EffectCRCLow
	target[offset+OffsetCRCHigh] = EffectCRCHigh
}

func parsePresets(buf []byte, count int) (preset.PresetList, error) {
	var presets = make(preset.PresetList, count)
	for i := 0; i < count; i++ {
		offset := i * EffectPageLength
		if buf[offset+OffsetCRCLow] != EffectCRCLow || buf[offset+OffsetCRCHigh] != EffectCRCHigh {
			return presets, ErrCRCMismatch
		}

		m, err := mode.ByCode(buf[offset+OffsetCode])
		if err != nil {
			return presets, err
		}

		presets[i].SetMode(m)

		if buf[offset+OffsetRandomColor] == 1 {
			presets[i].SetRandomColor()
		} else {
			presets[i].SetColor(color.New(
				buf[offset+OffsetR],
				buf[offset+OffsetG],
				buf[offset+OffsetB],
			))
		}

		presets[i].SetBrightness(buf[offset+OffsetBrightness])
		presets[i].SetSpeed(buf[offset+OffsetSpeed])

		presets[i].SetDirection(
			preset.GetDirection(buf[offset+OffsetDirection]),
		)
	}
	return presets, nil
}
