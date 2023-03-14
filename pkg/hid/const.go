package hid

// VendorID is the Keychron vendor ID.
const VendorID uint16 = 0x05AC

// ReportID is the Keychron protocol hid report ID.
const ReportID byte = 0x00

// PacketLength is the maximum allowed packet length.
const PacketLength = 64

// ConnectionAttempts represents the number of attempts to connect to the device.
const ConnectionAttempts = 50

// Usage constants.
// Experimentally it was found that the operation of the device is possible only with these values
const (
	ControlUsage     = 6
	ControlUsagePage = 1
)
