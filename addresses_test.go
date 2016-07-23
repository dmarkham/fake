package fake

import (
	"math/rand"
	"testing"
)

func TestCity(t *testing.T) {
	rand.Seed(0)
	for i := 0; i < 10; i++ {
		t.Run("City-deterministic", func(t *testing.T) {
			c := City()
			if c == "" {
				t.Fail()
			}
		})
	}
}

func TestAddressesStateAbbrev(t *testing.T) {
	for _, lang := range GetLangs() {
		err := SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		// Special case for StateAbbrev
		// We don't have RU state abbreviations, and english fallbacks are inappropriate
		v := StateAbbrev()
		if v == "" && lang != "ru" {
			t.Errorf("StateAbbrev failed with lang %s", lang)
		}
	}
}
