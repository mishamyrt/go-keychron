package main

import (
	"image/color"

	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/effect"
	"github.com/mishamyrt/go-keychron/pkg/hid"
)

func main() {
	hid.Init()
	b, err := keychron.Open(hid.K3V2Optical)
	if err != nil {
		panic(err)
	}
	// b.SetDebug(true)
	m, _ := effect.Get(effect.RainMode)
	m.Color = color.RGBA{0, 255, 255, 0}
	// current, err := b.Get()
	if err != nil {
		panic(err)
	}
	// fmt.Println(current)
	err = b.Set(m)
	if err != nil {
		panic(err)
	}
	// b.ReadEffect()
	// b.TestEnd()
}
