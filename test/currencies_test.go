package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var currencyFuncs = map[string]func() string{
	"Currency":     fake.Currency,
	"CurrencyCode": fake.CurrencyCode,
}

func TestCurrencies(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range currencyFuncs {
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
