package main

import (
	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/color"
	"github.com/mishamyrt/go-keychron/pkg/hid"
	"github.com/mishamyrt/go-keychron/pkg/keyboard"
	"github.com/mishamyrt/go-keychron/pkg/mode"
	"github.com/mishamyrt/go-keychron/pkg/preset"
)

func main() {
	hid.Init()
	b, err := keyboard.Open(keychron.K3v2Optical)
	if err != nil {
		panic(err)
	}
	b.SetDebug(true)

	MintRain := preset.New(
		&mode.Rain,
		color.New(0, 235, 47),
		0,
		0,
	)
	err = b.Set(&MintRain)
	if err != nil {
		panic(err)
	}
}
