package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

var loremIpsumFuncs = map[string]func() string{
	"Character":  fake.Character,
	"Characters": fake.Characters,
	"Word":       fake.Word,
	"Words":      fake.Words,
	"Title":      fake.Title,
	"Sentence":   fake.Sentence,
	"Sentences":  fake.Sentences,
	"Paragraph":  fake.Paragraph,
	"Paragraphs": fake.Paragraphs,
}

func TestLoremIpsum(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range loremIpsumFuncs {
			name, funct := name, funct // capture range variable
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				if a := funct(); a == "" {
					t.Errorf("%s failed with lang %s", name, lang)
				}
			})
		}

		v := fake.CharactersN(2)
		if v == "" {
			t.Errorf("CharactersN failed with lang %s", lang)
		}

		v = fake.WordsN(2)
		if v == "" {
			t.Errorf("WordsN failed with lang %s", lang)
		}

		v = fake.SentencesN(2)
		if v == "" {
			t.Errorf("SentencesN failed with lang %s", lang)
		}

		v = fake.ParagraphsN(2)
		if v == "" {
			t.Errorf("ParagraphsN failed with lang %s", lang)
		}

		vs := fake.WordsNUnique(251)
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
