package fake

import "testing"

func TestLoremIpsum(t *testing.T) {
	for _, lang := range GetLangs() {
		err := SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		v := CharactersN(2)
		if v == "" {
			t.Errorf("CharactersN failed with lang %s", lang)
		}

		v = WordsN(2)
		if v == "" {
			t.Errorf("WordsN failed with lang %s", lang)
		}

		v = SentencesN(2)
		if v == "" {
			t.Errorf("SentencesN failed with lang %s", lang)
		}

		v = ParagraphsN(2)
		if v == "" {
			t.Errorf("ParagraphsN failed with lang %s", lang)
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
