package device

import (
	"errors"
	"strconv"
	"strings"
)

// ParseAdapterIndex tries to parse the adapter name (e.g. hci0) and return its
// index (e.g. 0). It returns an error if it fails to parse.
func ParseAdapterIndex(adapter string) (int, error) {
	adapter = strings.TrimSpace(strings.ToLower(adapter))

	if strings.HasPrefix(adapter, "hci") {
		adapter = strings.TrimPrefix(adapter, "hci")
	}

	index, err := strconv.ParseInt(adapter, 10, 8)

	if err != nil {
		return 0, errors.New("parsing error")
	}

	return int(index), nil
}
