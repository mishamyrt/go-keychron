package main

import (
	"fmt"

	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/hid"
	"github.com/mishamyrt/go-keychron/pkg/keyboard"
)

func main() {
	hid.Init()
	b, err := keyboard.Open(keychron.K3v2Optical)
	if err != nil {
		panic(err)
	}
	if b.GetSync() {
		fmt.Println("Keyboard is in sync")
	} else {
		fmt.Printf("Keyboard is not in sync: %v", err)
	}

}
