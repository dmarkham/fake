package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var generalFuncs = map[string]func() string{
	"SimplePassword": fake.SimplePassword,
	"Color":          fake.Color,
	"HexColor":       fake.HexColor,
	"HexColorShort":  fake.HexColorShort,
	"Digits":         fake.Digits,
}

func TestGeneral(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range generalFuncs {
			name, funct := name, funct // capture range variable
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				if a := funct(); a == "" {
					t.Errorf("%s failed with lang %s", name, lang)
				}
			})
		}

		v := fake.Password(4, 10, true, true, true)
		if v == "" {
			t.Errorf("Password failed with lang %s", lang)
		}

		v = fake.DigitsN(2)
		if v == "" {
			t.Errorf("DigitsN failed with lang %s", lang)
		}

	}
}
