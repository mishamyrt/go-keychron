package main

import (
	"github.com/mishamyrt/go-keychron"
	"github.com/sstallion/go-hid"
)

const k3v2optical = 0x024F

func main() {
	k, err := keychron.Open(k3v2optical)
	if err != nil {
		hid.Exit()
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
		EffectValue: keychron.RainModeValue,
		Direction:   keychron.DirectionDTU,
	})
}
