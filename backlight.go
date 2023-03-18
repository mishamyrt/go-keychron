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

// Get current effect
func (b *Backlight) Get() (effect.Mode, error) {
	current, _, err := b.GetEffects()
	return current, err
}

// GetEffects reads all modes from device
func (b *Backlight) GetEffects() (current effect.Mode, effects []effect.Mode, err error) {
	err = b.requestEffectPages(ReadEffects)
	if err != nil {
		return
	}

	// For some reason the reading does not work consistently,
	// the keyboard may respond with a confirmation or immediately with the first page of effects.
	buf, err := b.handle.Read()
	if err != nil {
		return
	}
	var remaining int
	if buf[0] == PacketHeader && buf[1] == ReadEffects && buf[3] == CmdACK {
		// Got ACK, effects should be on next page
		remaining = 5
	} else if buf[OffsetCRCLow] == EffectCRCLow && buf[OffsetCRCHigh] == EffectCRCHigh {
		// Got first page, acting
		m, err := parseEffects(buf, 4)
		if err != nil {
			return current, m, err
		}
		effects = append(effects, m...)
		remaining = 4
	} else {
		// Not in sync
		return current, effects, ErrNotInSync
	}

	// Read modes, 4 or 5 pages
	m, err := b.readEffectPages(remaining)
	if err != nil {
		return current, effects, err
	}
	effects = append(effects, m...)

	// Skip padding and custom colors
	for i := 0; i < 12; i++ {
		b.handle.Read()
	}
	m, err = b.readEffectPage(1)
	if err != nil {
		return
	}
	current = m[0]
	b.endCommunication()
	return current, effects, nil
}

func (b *Backlight) SetDebug(enabled bool) {
	b.handle.Debug = enabled
	b.debug = enabled
}

func (b *Backlight) Close() error {
	return b.handle.Close()
}

func (b *Backlight) Set(m effect.Mode) error {
	// TODO: Fix the custom mode
	// k.setCustomization(false)

	err := b.requestEffectPages(WriteLEDEffects)
	if err != nil {
		return err
	}
	resp, err := b.handle.Read()
	if err != nil {
		return err
	}
	switch resp[3] {
	case CmdNACK:
		return ErrCmdNACK
	case CmdACK:
		break
	default:
		return ErrNotInSync
	}
	err = b.sendEffects()
	if err != nil {
		return err
	}
	err = b.sendCurrentEffect(&m)
	if err != nil {
		return err
	}
	return b.endCommunication()
}

func (b *Backlight) printDebug(m string) {
	if b.debug {
		log.Println(m)
	}
}

func (b *Backlight) sendCurrentEffect(m *effect.Mode) error {
	b.printDebug("Sending current effect")
	buf := make([]byte, EffectPageLength)
	fillEffect(m, buf, 0)
	return b.handle.Send(buf)
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

// requestEffectPages tells the device that EffectPages packets of data will be transmitted next
func (b *Backlight) requestEffectPages(command byte) error {
	b.printDebug(
		fmt.Sprintf("Requesting effects transmit, %d pages", EffectPages),
	)
	payload := make([]byte, 9)

	payload[0] = PacketHeader
	payload[1] = command
	payload[8] = EffectPages

	err := b.handle.Send(payload)
	if err != nil {
		return err
	}
	return nil
}

func (b *Backlight) endCommunication() error {
	b.printDebug("Ending communication")
	err := b.handle.Send([]byte{PacketHeader, CommunicationEnd})
	if err != nil {
		return err
	}
	buf, err := b.handle.Read()
	if err != nil {
		return err
	}
	switch buf[3] {
	case CmdACK:
		return nil
	case CmdNACK:
		return ErrCmdNACK
	default:
		return ErrNotInSync
	}
}

func (b *Backlight) applyEffects() error {
	b.printDebug("Applying effects")
	err := b.handle.Send([]byte{PacketHeader, ApplyLEDEffects})
	if err != nil {
		return err
	}
	_, err = b.handle.Read()
	return err
}

func (b *Backlight) readEffectPages(n int) ([]effect.Mode, error) {
	var effects []effect.Mode
	for i := 0; i < n; i++ {
		count := 0
		if i == n-1 {
			count = 3
		} else {
			count = 4
		}
		m, err := b.readEffectPage(count)
		if err != nil {
			return effects, err
		}
		effects = append(effects, m...)
	}
	return effects, nil
}

func (b *Backlight) readEffectPage(n int) ([]effect.Mode, error) {
	buf, err := b.handle.Read()
	if err != nil {
		return []effect.Mode{}, err
	}
	return parseEffects(buf, n)
}

func Open(productId uint16) (Backlight, error) {
	h, err := hid.Open(productId)
	return Backlight{handle: h}, err
}
