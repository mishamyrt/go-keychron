package keychron

import (
	"fmt"
	"log"

	"github.com/mishamyrt/go-keychron/pkg/effect"
	"github.com/mishamyrt/go-keychron/pkg/hid"
)

type Backlight struct {
	handle hid.Handle
	debug  bool
}

func (b *Backlight) SetDebug(enabled bool) {
	b.handle.Debug = enabled
	b.debug = enabled
}

func (b *Backlight) Close() error {
	return b.handle.Close()
}

func (b *Backlight) Set(m effect.Mode) {
	// TODO: Fix the custom mode
	// k.setCustomization(false)

	b.requestEffectsWrite(18)
	b.sendEffects()
	b.sendCurrentEffect(&m)
	b.endCommunication()
}

func (b *Backlight) printDebug(m string) {
	if b.debug {
		log.Println(m)
	}
}

func (b *Backlight) sendCurrentEffect(m *effect.Mode) {
	b.printDebug("Sending current effect")
	buf := make([]byte, EffectPageLength)
	fillEffect(m, buf, 0)
	b.handle.Send(buf)
}

// Send effects to keyboard.
func (b *Backlight) sendEffects() error {
	b.printDebug("Sending effects")
	var buf []byte
	var err error

	for i := 0; i < 5; i++ { // 5 packets
		buf = make([]byte, hid.PacketLength)
		for j := 0; j < 4; j++ { // of 4 effects
			modeOffset := j + (i * 4)
			if modeOffset >= len(effect.Modes) {
				continue
			}

			m := effect.Modes[modeOffset]

			fillEffect(&m, buf, j*EffectPageLength)
		}
		err = b.handle.Send(buf)
		if err != nil {
			return err
		}
	}

	// For some reason we have to send 3 empty packets between effects and colors ¯\_(ツ)_/¯
	buf = make([]byte, hid.PacketLength)
	for i := 0; i < 3; i++ {
		err = b.handle.Send(buf)
		if err != nil {
			return err
		}
	}

	// Colors for custom mode, 9 packets
	for i := 0; i < 9; i++ {
		buf = make([]byte, hid.PacketLength)
		for j := 0; j < EffectPageLength; j++ {
			buf[j*4] = CustomColorHeader
			buf[j*4+1] = 0xFF // R
			buf[j*4+2] = 0xFF // G
			buf[j*4+3] = 0xFF // B
		}
		b.handle.Send(buf)
	}
	return nil
}

// requestEffectsWrite tells the device that N packets of data will be transmitted next
func (b *Backlight) requestEffectsWrite(n uint8) error {
	b.printDebug(
		fmt.Sprintf("Requesting effects write, %d pages", n),
	)
	payload := make([]byte, 9)

	payload[0] = PacketHeader
	payload[1] = WriteLEDEffects
	payload[8] = n

	err := b.handle.Send(payload)
	if err != nil {
		return err
	}
	return b.handle.Read()
}

func (b *Backlight) endCommunication() error {
	b.printDebug("Ending communication")
	err := b.handle.Send([]byte{PacketHeader, CommunicationEnd})
	if err != nil {
		return err
	}
	return b.handle.Read()
}

func (b *Backlight) applyEffects() error {
	b.printDebug("Applying effects")
	err := b.handle.Send([]byte{PacketHeader, ApplyLEDEffects})
	if err != nil {
		return err
	}
	return b.handle.Read()
}

func Open(productId uint16) (Backlight, error) {
	h, err := hid.Open(productId)
	return Backlight{handle: h}, err
}
