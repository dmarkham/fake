package fake

import "testing"

var loremIpsumFuncs = map[string]func(int) string{
	"CharactersN": CharactersN,
	"WordsN":      WordsN,
	"SentencesN":  SentencesN,
	"ParagraphsN": ParagraphsN,
}

func TestLoremIpsumN(t *testing.T) {
	for _, lang := range GetLangs() {
		err := SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range loremIpsumFuncs {
			name, funct := name, funct // capture range variable
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				if a := funct(2); a == "" {
					t.Errorf("%s failed with lang %s", name, lang)
				}
			})
		}
	}
}

func TestLoremIpsumUnique(t *testing.T) {
	for _, lang := range GetLangs() {
		err := SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		vs := WordsNUnique(251)
		if len(vs) != 251 {
			t.Errorf("WordsNUnique failed with lang %s", lang)
		}
		lw := ""
		for n, w := range vs {
			if w == lw {
				t.Errorf("Duplicate words found")
			}
			if w == "" {
				t.Errorf("Word %d is empty", n)
				continue
			}
			lw = w
		}
	}
}
