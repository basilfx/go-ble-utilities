package device

import (
	"errors"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}

// DefaultDeviceWithName ...
func DefaultDeviceWithName(name string, opts ...ble.Option) (d ble.Device, err error) {
	return nil, errors.New("Not supported")
}
