package fake

import "testing"

func TestIPv4(t *testing.T) {
	t.Parallel()

	if ip := IPv4(); ip == "" {
		t.Error("IPv4 failed")
	}
}
