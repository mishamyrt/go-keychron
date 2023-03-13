package main

import (
	"github.com/mishamyrt/go-keychron"
)

const k3v2optical = 0x024F

func main() {
	k, err := keychron.Open(k3v2optical)
	if err != nil {
		panic(err)
	}
	defer k.Close()
	k.Set(keychron.Mode{
		Color: keychron.Color{
			Red:   0,
			Green: 255,
			Blue:  0,
		},
		Brightness:  keychron.MaxBrightness,
		Speed:       keychron.MinSpeed,
		EffectValue: keychron.RingGradientModeValue,
		Direction:   keychron.DirectionDTU,
	})
}
