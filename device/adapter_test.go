package device_test

import (
	"testing"

	"github.com/basilfx/go-ble-utilities/device"
	"github.com/stretchr/testify/assert"
)

// Test_ParseAdapterIndex_Valid tests the ParseAdapterIndex method for a valid
// input.
func Test_ParseAdapterIndex_Valid(t *testing.T) {
	index, err := device.ParseAdapterIndex("hci1")

	assert.Equal(t, 1, index)
	assert.Nil(t, err)
}

// Test_ParseAdapterIndex_Alternative tests the ParseAdapterIndex method for an
// alternative input.
func Test_ParseAdapterIndex_Alternative(t *testing.T) {
	index, err := device.ParseAdapterIndex("123")

	assert.Equal(t, 123, index)
	assert.Nil(t, err)
}

// Test_ParseAdapterIndex tests the ParseAdapterIndex method for a valid
// input.
func Test_ParseAdapterIndex_Invalid(t *testing.T) {
	_, err := device.ParseAdapterIndex("xyz1")

	assert.NotNil(t, err)
}
