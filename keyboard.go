package keychron

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sstallion/go-hid"
)

var ErrWrongUsage = errors.New("wrong device usage")

type Keyboard struct {
	dev   *hid.Device
	Debug bool
}

func (k *Keyboard) SendRaw(packet []byte) error {
	buf := make([]byte, PacketLength)

	for x := 0; x < len(packet); x++ {
		buf[x] = packet[x]
	}

	_, err := k.dev.SendFeatureReport(buf)
	if err != nil {
		return err
	}

	k.waitSync()
	return nil
}

func (k *Keyboard) Send(packet []byte) error {
	buf := make([]byte, PacketLength+1)
	buf[0] = ReportID

	for x := 0; x < len(packet); x++ {
		buf[x+1] = packet[x]
	}

	_, err := k.dev.SendFeatureReport(buf)
	if err != nil {
		return err
	}

	k.waitSync()
	return nil
}

func (k *Keyboard) Read() error {
	buf := make([]byte, PacketLength+1)
	buf[0] = ReportID

	_, err := k.dev.GetFeatureReport(buf)
	if err != nil {
		return err
	}
	log.Println("Readed: ", buf)
	k.waitSync()
	return nil
}

func (k *Keyboard) Close() error {
	return k.dev.Close()
}

func (k *Keyboard) Test() {
	var buf []byte
	c := Color{
		Red:   255,
		Green: 0,
		Blue:  255,
	}

	k.setCustomization(true)
	k.requestWrite(18)
	// Effects
	k.Send([]byte{
		// 0,     1,       2,      3,    4,    5,    6,    7,    8,    9,   10,   11,   12,   13,   14,   15
		0x01, c.Red, c.Green, c.Blue, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x0c, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x02, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x03, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x04, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
	})
	k.Send([]byte{
		0x05, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x06, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x07, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x08, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
	})
	k.Send([]byte{
		0x09, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x0a, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x0b, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x01, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x0c, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
	})
	k.Send([]byte{
		0x0d, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x0e, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x0f, 0xff, 0x9b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x10, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
	})
	k.Send([]byte{
		0x11, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x12, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x01, 0x00, 0x00, 0xaa, 0x55,
		0x13, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x0F, 0x00, 0x00, 0x00, 0xaa, 0x55,
		0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0xaa, 0x55,
	})

	buf = make([]byte, PacketLength)
	for i := 0; i < 3; i++ {
		err := k.Send(buf)
		if err != nil {
			log.Println("Error while writing empty packet", err)
		}
	}

	for i := 0; i < 9; i++ {
		buf = make([]byte, PacketLength)
		for j := 0; j < EffectPageLength; j++ {
			buf[j*4] = 0x80
			buf[j*4+1] = c.Red
			buf[j*4+2] = c.Green
			buf[j*4+3] = c.Blue
		}
		k.Send(buf)
	}

	// Current effect
	buf = make([]byte, PacketLength)
	buf[0x00] = RingGradientModeValue
	buf[0x01] = c.Red
	buf[0x02] = c.Green
	buf[0x03] = c.Blue
	buf[0x09] = MaxBrightness
	buf[0x0A] = MinSpeed
	buf[0x0E] = 0xaa
	buf[0x0F] = 0x55
	k.Send(buf)

	k.endCommunication()
	k.applyEffects()
}

func (k *Keyboard) SetModes(modes []Mode, active uint8) error {
	var buf []byte
	var selectedMode = make([]byte, EffectPageLength)
	var err error

	k.setCustomization(active == CustomModeValue)
	k.requestWrite(2)

	for i := 0; i < 5; i++ { // 5 packets
		buf = make([]byte, PacketLength) // of 4 effects
		for j := 0; j < 4; j++ {
			if len(modes) <= j+(i*4) {
				continue
			}
			m := modes[j+(i*4)]
			offset := j * EffectPageLength

			buf[offset+0] = m.EffectValue // mode value

			buf[offset+1] = m.Color.Red
			buf[offset+2] = m.Color.Green
			buf[offset+3] = m.Color.Blue

			buf[offset+8] = 0 // Random color, disable for now
			buf[offset+9] = m.Brightness
			buf[offset+10] = m.Speed
			buf[offset+11] = m.Direction

			buf[offset+14] = EffectPageCRCLow
			buf[offset+15] = EffectPageCRCHigh

			if m.EffectValue == active {
				buf[offset+9] = MaxBrightness
				for x := 0; x < EffectPageLength; x++ {
					selectedMode[x] = buf[offset+x]
				}
			}
		}
		log.Printf("Effects page %x: %v", i+1, buf)
		err = k.Send(buf)
		if err != nil {
			log.Println("Error while writing page", err)
			return err
		}
		k.Read()
	}

	// 3 times an empty packet - guess why...
	buf = make([]byte, PacketLength)
	for i := 0; i < 3; i++ {
		err = k.Send(buf)
		if err != nil {
			log.Println("Error while writing empty packet", err)
			return err
		}
	}

	// Customization stuff
	// 9 times * 16 blocks 80 RR GG BB

	// colorBuf := make([]byte, ColorBufSize)
	// for i := 0; i < ColorBufSize; i += 4 {
	// 	colorBuf[i] = 0x80
	// 	colorBuf[i+1] = color.Red
	// 	colorBuf[i+2] = color.Green
	// 	colorBuf[i+3] = color.Blue
	// }

	// for p := 0; p < 9; p++ {
	// 	for x := 0; x < EffectPageLength; x++ {
	// 		buf[x] = colorBuf[p*PacketLength+x]
	// 		err = k.Send(buf)
	// 		if err != nil {
	// 			log.Println("Error while writing color page", err)
	// 			return err
	// 		}
	// 	}
	// }

	log.Println("Selected:", selectedMode)
	for i := 0; i < EffectPageLength; i++ {
		buf[i] = selectedMode[i]
	}
	err = k.Send(buf)
	if err != nil {
		log.Println("Error while writing selected mode", err)
		return err
	}

	err = k.endCommunication()
	if err != nil {
		log.Println("Error while ending communication", err)
		return err
	}

	// err = k.applyEffects()
	// if err != nil {
	// 	log.Println("Error while starting effect", err)
	// 	return err
	// }
	// buf = make([]byte, PacketLength)
	// for i := 0; i < 3; i++ {
	// 	err = k.Send(buf)
	// 	if err != nil {
	// 		log.Println("Error while writing empty packet", err)
	// 		return err
	// 	}
	// }
	return nil
}

func (k *Keyboard) waitSync() {
	time.Sleep(time.Millisecond * 15)
}

func (k *Keyboard) setCustomization(state bool) error {
	buf := make([]byte, PacketLength)
	buf[0] = PacketHeader
	if state {
		buf[1] = TurnOnCustomizationCommand
	} else {
		buf[1] = TurnOffCustomizationCommand
	}
	err := k.Send(buf)
	if err != nil {
		return err
	}
	return k.Read()
}

func (k *Keyboard) requestWrite(count uint8) error {
	buf := make([]byte, PacketLength)

	buf[0] = PacketHeader
	buf[1] = WriteLEDSpecialEffectAreaCommand
	buf[8] = count

	err := k.Send(buf)
	if err != nil {
		return err
	}
	return k.Read()
}

func (k *Keyboard) endCommunication() error {
	err := k.Send([]byte{PacketHeader, CommunicationEndCommand})
	if err != nil {
		return err
	}
	return k.Read()
}

func (k *Keyboard) applyEffects() error {
	err := k.Send([]byte{PacketHeader, LEDEffectStartCommand})
	if err != nil {
		return err
	}
	return k.Read()
}

func (k *Keyboard) tryOpen(productId uint16) error {
	dev, err := hid.OpenFirst(VendorID, productId)
	if err != nil {
		return err
	}
	info, err := dev.GetDeviceInfo()
	if err != nil {
		return err
	}
	if info.Usage != 6 || info.UsagePage != 1 {
		dev.Close()
		return ErrWrongUsage
	}

	k.dev = dev
	return nil
}

func Open(productId uint16) (Keyboard, error) {
	k := Keyboard{}
	var err error
	err = hid.Init()
	if err != nil {
		return k, err
	}
	hid.SetOpenExclusive(false)
	for i := 0; i < 50; i++ {
		fmt.Println("Attempt", i)
		err = k.tryOpen(productId)
		if err == nil {
			fmt.Println("Found!")
			return k, nil
		}
	}
	fmt.Println("Not found(")
	return k, err
}
