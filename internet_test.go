package fake

import "testing"

func TestUsername(t *testing.T) {
	for i := 0; i < 6; i++ {
		t.Run("Username-deterministic", func(t *testing.T) {
			u := UserName()
			if u == "" {
				t.Fail()
			}
		})
	}
}

func TestIPv4(t *testing.T) {
	t.Parallel()

	if ip := IPv4(); ip == "" {
		t.Error("IPv4 failed")
	}
}
