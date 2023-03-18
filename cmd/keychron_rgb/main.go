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
	// b.SetDebug(true)
	m, _ := effect.Get(effect.RainMode)
	m.Color = color.RGBA{0, 255, 255, 0}
	// current, err := b.Get()
	if err != nil {
		panic(err)
	}
	// fmt.Println(current)
	// m.Code = 0
	colors := make([]color.RGBA, 144)
	selectedNumbers := []int{
		19, 35, 36, 51, 21, 66,
	}
	for i := 0; i < len(selectedNumbers); i++ {
		colors[selectedNumbers[i]] = color.RGBA{255, 255, 255, 0}
	}
	err = b.SetCustom(colors, effect.Brightest)
	if err != nil {
		panic(err)
	}
	// b.ReadEffect()
	// b.TestEnd()
}
