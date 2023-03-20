package main

import (
	"image/color"

	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/effect"
	"github.com/mishamyrt/go-keychron/pkg/hid"
	"github.com/mishamyrt/go-keychron/pkg/keyboard"
)

func main() {
	hid.Init()
	b, err := keyboard.Open(keychron.K3v2Optical)
	if err != nil {
		panic(err)
	}
	b.SetDebug(true)

	p := effect.NewPreset(
		&effect.RingGradientMode,
		color.RGBA{0, 255, 255, 0},
		effect.Slowest,
		effect.Brightest,
		0,
	)
	err = b.Set(p)
	if err != nil {
		panic(err)
	}
}
