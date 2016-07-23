package fake

import "testing"

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
