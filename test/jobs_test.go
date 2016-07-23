package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var jobsFuncs = map[string]func() string{
	"Company":  fake.Company,
	"JobTitle": fake.JobTitle,
	"Industry": fake.Industry,
}

func TestJobs(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range jobsFuncs {
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
