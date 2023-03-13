package main

import (
	"log"

	"github.com/mishamyrt/go-keychron"
	"github.com/sstallion/go-hid"
)

const k3vid = 0x024F

func main() {
	k, err := keychron.Open(k3vid)
	if err != nil {
		hid.Exit()
		panic(err)
	}
	defer k.Close()
	log.Println(k)
	// modes := keychron.CreateModes(keychron.Color{0, 0, 255}, 255)
	k.Set()
	// // hid.Enumerate(keychron.VendorID, keychron.K3V2RGBOptical, func(info *hid.DeviceInfo) error {
	// // 	if info.Usage == 6 && info.UsagePage == 1 {}
	// // 	log. ("%#04x %#04x %s\n", info.Usage, info.UsagePage, info.Path)
	// // 	return nil
	// // })
	// // fmt.Println("after")
	// kb := tryOpen(50)
	// defer kb.Close()
	// kb.SetMode(2, keychron.Color{0, 255, 0})
	// // kb.SetCustomization(false)
}
