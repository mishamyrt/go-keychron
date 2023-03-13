package keychron

import (
	"errors"
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

	if k.Debug {
		log.Println("Send", buf)
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
	if k.Debug {
		log.Println("Read", buf)
	}
	k.waitSync()
	return nil
}

func (k *Keyboard) Close() error {
	return k.dev.Close()
}

func (k *Keyboard) Set(mode Mode) {
	k.setCustomization(false)
	k.requestWrite(18)

	k.SendEffects()

	buf := make([]byte, PacketLength)
	buf[0] = mode.EffectValue
	buf[1] = mode.Color.Red
	buf[2] = mode.Color.Green
	buf[3] = mode.Color.Blue

	if mode.RandomColor {
		buf[8] = 1
	}
	buf[9] = mode.Brightness
	buf[10] = mode.Speed
	buf[11] = mode.Direction
	buf[14] = EffectPageCRCLow
	buf[15] = EffectPageCRCHigh

	k.Send(buf)
	k.endCommunication()
}

// Send effects to Keyboard.
func (k *Keyboard) SendEffects() error {
	var buf []byte
	var err error

	for i := 0; i < 5; i++ { // 5 packets
		buf = make([]byte, PacketLength)
		for j := 0; j < 4; j++ { // of 4 effects
			modeOffset := j + (i * 4)
			if modeOffset >= len(Modes) {
				continue
			}
			m := Modes[j+(i*4)]
			offset := j * EffectPageLength

			buf[offset+0] = m.EffectValue
			buf[offset+1] = m.Color.Red
			buf[offset+2] = m.Color.Green
			buf[offset+3] = m.Color.Blue

			buf[offset+8] = 1 // Random switch
			buf[offset+9] = m.Brightness
			buf[offset+10] = m.Speed
			buf[offset+11] = m.Direction

			buf[offset+14] = EffectPageCRCLow
			buf[offset+15] = EffectPageCRCHigh
		}
		err = k.Send(buf)
		if err != nil {
			return err
		}
	}

	// For some reason you have to send 3 empty packets between effects and colors ¯\_(ツ)_/¯
	buf = make([]byte, PacketLength)
	for i := 0; i < 3; i++ {
		err = k.Send(buf)
		if err != nil {
			return err
		}
	}

	// Colors for custom mode, 9 packets
	for i := 0; i < 9; i++ {
		buf = make([]byte, PacketLength)
		for j := 0; j < EffectPageLength; j++ {
			buf[j*4] = CustomColorHeader
			buf[j*4+1] = 0xFF // R
			buf[j*4+2] = 0xFF // G
			buf[j*4+3] = 0xFF // B
		}
		k.Send(buf)
	}
	return nil
}

// Waiting for data to be written. According to the documentation, it takes 10 ms to do this
func (k *Keyboard) waitSync() {
	time.Sleep(time.Millisecond * 10)
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
		err = k.tryOpen(productId)
		if err == nil {
			return k, nil
		}
	}
	hid.Exit()
	return k, err
}
