package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var addressFuncs = map[string]func() string{
	"Continent":     fake.Continent,
	"Country":       fake.Country,
	"City":          fake.City,
	"Phone":         fake.Phone,
	"State":         fake.State,
	"Street":        fake.Street,
	"StreetAddress": fake.StreetAddress,
	"Zip":           fake.Zip,
}

func TestAddresses(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range addressFuncs {
			name, funct := name, funct // capture range variable
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				if a := funct(); a == "" {
					t.Errorf("%s failed with lang %s", name, lang)
				}
			})
		}

		// Special case for StateAbbrev
		// We don't have RU state abbreviations, and english fallbacks are inapproprirate
		v := fake.StateAbbrev()
		if v == "" && lang != "ru" {
			t.Errorf("StateAbbrev failed with lang %s", lang)
		}
	}
}
