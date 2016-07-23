package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var productsFuncs = map[string]func() string{
	"Brand":       fake.Brand,
	"ProductName": fake.ProductName,
	"Product":     fake.Product,
	"Model":       fake.Model,
}

func TestProducts(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range productsFuncs {
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
