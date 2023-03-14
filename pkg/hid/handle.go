package hid

import (
	"log"
	"time"

	"github.com/sstallion/go-hid"
)

// Handle represents Keychron device handle
type Handle struct {
	Device *hid.Device
	// Debug flag. If true, then debugging data will be written to stdout when functions are executed.
	Debug bool
}

// Send packet to the device.
func (h *Handle) Send(payload []byte) error {
	if len(payload) > PacketLength {
		return ErrPayloadOverflow
	}

	packet := make([]byte, PacketLength+1)
	packet[0] = ReportID

	for i := 0; i < len(payload); i++ {
		packet[i+1] = payload[i]
	}

	if h.Debug {
		log.Printf("Send: %v", packet)
	}

	transferred, err := h.Device.SendFeatureReport(packet)
	if err != nil {
		return err
	}
	expected := len(packet)
	if transferred != len(packet) {
		return NewErrCountMismatch(expected, transferred)
	}

	h.waitSync()
	return nil
}

// Read packet from the device.
func (h *Handle) Read() error {
	packet := make([]byte, PacketLength+1)
	packet[0] = ReportID

	_, err := h.Device.GetFeatureReport(packet)
	if err != nil {
		return err
	}
	if h.Debug {
		log.Printf("Read: %v", packet)
	}
	h.waitSync()
	return nil
}

// Close device handle.
// The function should be called after the end of operation with the device.
func (h *Handle) Close() error {
	return h.Device.Close()
}

// Waiting for data to be written. According to the documentation, it takes 10 ms to do this
func (h *Handle) waitSync() {
	time.Sleep(time.Millisecond * 10)
}

// For some reason libhid doesn't let us open the right handle directly,
// so instead just try random handles until you get the right one.
func (h *Handle) tryOpen(productId uint16) error {
	d, err := hid.OpenFirst(VendorID, productId)
	if err != nil {
		return err
	}
	info, err := d.GetDeviceInfo()
	if err != nil {
		return err
	}
	if info.Usage != ControlUsage || info.UsagePage != ControlUsagePage {
		d.Close()
		return ErrWrongUsage
	}

	h.Device = d
	return nil
}

// Open tries to open the device to execute commands.
func Open(productId uint16) (h Handle, err error) {
	err = hid.Init()
	if err != nil {
		return
	}
	for i := 0; i < ConnectionAttempts; i++ {
		err = h.tryOpen(productId)
		if err == nil {
			return h, nil
		}
	}
	return
}
