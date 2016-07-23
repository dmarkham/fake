package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var personalFuncs = map[string]func() string{
	"Gender":       fake.Gender,
	"GenderAbbrev": fake.GenderAbbrev,
	"Language":     fake.Language,
}

func TestPersonal(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range personalFuncs {
			name, funct := name, funct // capture range variable
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				if a := funct(); a == "" {
					t.Errorf("%s failed with lang %s", name, lang)
				}
			})
		}
	}
}
