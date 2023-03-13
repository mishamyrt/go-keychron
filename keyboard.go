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

func (k *Keyboard) Send(packet []byte) error {
	buf := make([]byte, PacketLength+1)
	buf[0] = ReportID

	for x := 0; x < len(packet); x++ {
		buf[x+1] = packet[x]
	}

	log.Println("Sending", buf)

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

func (k *Keyboard) Set() {
	var buf []byte
	color := Color{
		Red:   255,
		Green: 0,
		Blue:  255,
	}
	brightness := mapBrightness(255)
	var mode uint8 = RandomColorsModeValue
	speed := 0x01

	k.setCustomization(false)
	k.requestWrite(18)

	fmt.Println("\n\nWriting effects")
	for i := 0; i < 5; i++ { // 5 packets
		buf = make([]byte, PacketLength) // of 4 effects
		for j := 0; j < 4; j++ {
			modeOffset := j + (i * 4)
			if modeOffset >= len(Modes) {
				continue
			}
			m := Modes[j+(i*4)]
			offset := j * EffectPageLength

			buf[offset+0] = m.EffectValue // mode value

			buf[offset+1] = m.Color.Red
			buf[offset+2] = m.Color.Green
			buf[offset+3] = m.Color.Blue

			buf[offset+8] = 1
			buf[offset+9] = m.Brightness
			buf[offset+10] = m.Speed
			buf[offset+11] = m.Direction

			buf[offset+14] = EffectPageCRCLow
			buf[offset+15] = EffectPageCRCHigh
		}
		log.Printf("Effects page %x: %v", i+1, buf)
		err := k.Send(buf)
		if err != nil {
			log.Println("Error while writing page", err)
		}
	}

	fmt.Println("\n\nWriting empty")
	buf = make([]byte, PacketLength)
	for i := 0; i < 3; i++ {
		err := k.Send(buf)
		if err != nil {
			log.Println("Error while writing empty packet", err)
		}
	}

	fmt.Println("\n\nWriting custom colors")
	for i := 0; i < 9; i++ {
		buf = make([]byte, PacketLength)
		for j := 0; j < EffectPageLength; j++ {
			buf[j*4] = CustomColorHeader
			buf[j*4+1] = color.Red
			buf[j*4+2] = color.Green
			buf[j*4+3] = color.Blue
		}
		k.Send(buf)
	}

	fmt.Println("\n\nWriting current effect")
	buf = make([]byte, PacketLength)
	buf[0x00] = mode
	buf[0x01] = color.Red
	buf[0x02] = color.Green
	buf[0x03] = color.Blue
	buf[0x09] = brightness
	buf[0x0A] = byte(speed)
	buf[0x0E] = 0xaa
	buf[0x0F] = 0x55
	k.Send(buf)

	fmt.Println("\n\nFinalizing")
	k.endCommunication()
	k.applyEffects()
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
