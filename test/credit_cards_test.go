package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

func TestCreditCards(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		t.Run("CreditCardType", func(t *testing.T) {
			t.Parallel()
			v := fake.CreditCardType()
			if v == "" {
				t.Errorf("CreditCardType failed with lang %s", lang)
			}
		})

		t.Run("CreditCardNum-RandomVendor", func(t *testing.T) {
			t.Parallel()
			v := fake.CreditCardNum("")
			if v == "" {
				t.Errorf("CreditCardNum failed with lang %s", lang)
			}
		})

		t.Run("CreditCardNum-SpecificVendor", func(t *testing.T) {
			t.Parallel()
			v := fake.CreditCardNum("visa")
			if v == "" {
				t.Errorf("CreditCardNum failed with lang %s", lang)
			}
		})

	}
}
