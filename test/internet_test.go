package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var internetFuncs = map[string]func() string{
	"UserName":       fake.UserName,
	"TopLevelDomain": fake.TopLevelDomain,
	"DomainName":     fake.DomainName,
	"EmailAddress":   fake.EmailAddress,
	"EmailSubject":   fake.EmailSubject,
	"EmailBody":      fake.EmailBody,
	"DomainZone":     fake.DomainZone,
}

func TestInternet(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range internetFuncs {
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

func TestIPv4(t *testing.T) {
	t.Parallel()

	if ip := fake.IPv4(); ip == "" {
		t.Error("IPv4 failed")
	}
}
