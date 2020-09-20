package device

import (
	"github.com/go-ble/ble"
)

// NewDevice ...
func NewDevice(opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}

func NewDeviceWithName(name string, opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDeviceWithName(name, opts...)
}
