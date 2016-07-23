package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var generalFuncs = map[string]func() string{
	"SimplePassword": fake.SimplePassword,
	"HexColor":       fake.HexColor,
	"HexColorShort":  fake.HexColorShort,
	"Digits":         fake.Digits,
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

func TestColor(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		if c := fake.Color(); c == "" {
			t.Errorf("Color failed with language %s", lang)
		}
	}
}

func TestDigitsN(t *testing.T) {
	t.Parallel()

	v := fake.DigitsN(2)
	if v == "" {
		t.Error("DigitsN failed")
	}
}
func TestPassword(t *testing.T) {
	t.Parallel()

	v := fake.Password(4, 10, true, true, true)
	if v == "" {
		t.Error("Password failed")
	}
}
