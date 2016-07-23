package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var namesFuncs = map[string]func() string{
	"MaleFirstName":            fake.MaleFirstName,
	"FemaleFirstName":          fake.FemaleFirstName,
	"FirstName":                fake.FirstName,
	"MaleLastName":             fake.MaleLastName,
	"FemaleLastName":           fake.FemaleLastName,
	"LastName":                 fake.LastName,
	"MalePatronymic":           fake.MalePatronymic,
	"FemalePatronymic":         fake.FemalePatronymic,
	"Patronymic":               fake.Patronymic,
	"MaleFullNameWithPrefix":   fake.MaleFullNameWithPrefix,
	"FemaleFullNameWithPrefix": fake.FemaleFullNameWithPrefix,
	"FullNameWithPrefix":       fake.FullNameWithPrefix,
	"MaleFullNameWithSuffix":   fake.MaleFullNameWithSuffix,
	"FemaleFullNameWithSuffix": fake.FemaleFullNameWithSuffix,
	"FullNameWithSuffix":       fake.FullNameWithSuffix,
	"MaleFullName":             fake.MaleFullName,
	"FemaleFullName":           fake.FemaleFullName,
	"FullName":                 fake.FullName,
}

func TestNames(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range namesFuncs {
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
