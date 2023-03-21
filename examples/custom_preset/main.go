package main

import (
	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/color"
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

	MintRain := effect.NewPreset(
		&effect.RainMode,
		color.New(0, 235, 47),
		effect.Slowest,
		0,
	)
	err = b.Set(&MintRain)
	if err != nil {
		panic(err)
	}
}
