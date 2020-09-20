package device

import (
	"github.com/go-ble/ble"
	"github.com/go-ble/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}

// DefaultDeviceWithName ...
func DefaultDeviceWithName(name string, opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDeviceWithName(name, opts...)
}
