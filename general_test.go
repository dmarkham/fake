package fake

import "testing"

var generalFuncs = map[string]func() string{
	"SimplePassword": SimplePassword,
	"HexColor":       HexColor,
	"HexColorShort":  HexColorShort,
	"Digits":         Digits,
}

func TestGeneral(t *testing.T) {
	t.Parallel()

	for name, funct := range generalFuncs {
		name, funct := name, funct // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if a := funct(); a == "" {
				t.Errorf("%s failed", name)
			}
		})
	}

}

func TestDigitsN(t *testing.T) {
	t.Parallel()

	v := DigitsN(2)
	if v == "" {
		t.Error("DigitsN failed")
	}
}
func TestPassword(t *testing.T) {
	t.Parallel()

	v := Password(4, 10, true, true, true)
	if v == "" {
		t.Error("Password failed")
	}
}
