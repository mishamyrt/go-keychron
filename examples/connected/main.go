package main

import (
	"fmt"

	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/hid"
	"github.com/mishamyrt/go-keychron/pkg/keyboard"
)

// IsConnected checks if Keychron keyboard is connected to computer
func IsConnected(productId uint16) bool {
	_, err := keyboard.Open(keychron.K3v2Optical)
	return err == nil
}

func main() {
	hid.Init()

	var msg string
	if IsConnected(keychron.K3v2Optical) {
		msg = "Keyboard is connected ✌️"
	} else {
		msg = "Keyboard is not connected 🤔"
	}
	fmt.Println(msg)
}
