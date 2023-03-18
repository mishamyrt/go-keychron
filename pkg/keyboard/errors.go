package keyboard

import "errors"

var ErrCRCMismatch = errors.New("got wrong CRC")
var ErrNotInSync = errors.New("device is out of sync and is not responding as expected. Try to power cycle")
var ErrCmdNACK = errors.New("device responds with error code")
var ErrColorsMiscount = errors.New("wrong colors count. Should be exactly 144")
