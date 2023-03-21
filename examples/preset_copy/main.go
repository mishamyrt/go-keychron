package main

import (
	"github.com/mishamyrt/go-keychron"
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

	p := preset.For(&mode.Rain)
	var SlowRain preset.Preset = *p
	SlowRain.SetSpeed(0)

	err = b.Set(&SlowRain)
	if err != nil {
		panic(err)
	}
}
