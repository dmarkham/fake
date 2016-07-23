package fake

import "testing"

var simpleFuncs = map[string]func() string{
	// Addresses
	"Continent":     Continent,
	"Country":       Country,
	"City":          City,
	"Phone":         Phone,
	"State":         State,
	"Street":        Street,
	"StreetAddress": StreetAddress,
	"Zip":           Zip,

	// Currencies
	"Currency":     Currency,
	"CurrencyCode": CurrencyCode,

	// Dates
	"WeekDay":      WeekDay,
	"WeekDayShort": WeekDayShort,
	"Month":        Month,
	"MonthShort":   MonthShort,

	// General
	"Color": Color,

	// Internet
	"UserName":       UserName,
	"TopLevelDomain": TopLevelDomain,
	"DomainName":     DomainName,
	"EmailAddress":   EmailAddress,
	"EmailSubject":   EmailSubject,
	"EmailBody":      EmailBody,
	"DomainZone":     DomainZone,

	// Jobs
	"Company":  Company,
	"JobTitle": JobTitle,
	"Industry": Industry,

	// Lorem Ipsum
	"Character":  Character,
	"Characters": Characters,
	"Word":       Word,
	"Words":      Words,
	"Title":      Title,
	"Sentence":   Sentence,
	"Sentences":  Sentences,
	"Paragraph":  Paragraph,
	"Paragraphs": Paragraphs,

	// Names
	"MaleFirstName":            MaleFirstName,
	"FemaleFirstName":          FemaleFirstName,
	"FirstName":                FirstName,
	"MaleLastName":             MaleLastName,
	"FemaleLastName":           FemaleLastName,
	"LastName":                 LastName,
	"MalePatronymic":           MalePatronymic,
	"FemalePatronymic":         FemalePatronymic,
	"Patronymic":               Patronymic,
	"MaleFullNameWithPrefix":   MaleFullNameWithPrefix,
	"FemaleFullNameWithPrefix": FemaleFullNameWithPrefix,
	"FullNameWithPrefix":       FullNameWithPrefix,
	"MaleFullNameWithSuffix":   MaleFullNameWithSuffix,
	"FemaleFullNameWithSuffix": FemaleFullNameWithSuffix,
	"FullNameWithSuffix":       FullNameWithSuffix,
	"MaleFullName":             MaleFullName,
	"FemaleFullName":           FemaleFullName,
	"FullName":                 FullName,

	// Personal
	"Gender":       Gender,
	"GenderAbbrev": GenderAbbrev,
	"Language":     Language,

	// Products
	"Brand":       Brand,
	"ProductName": ProductName,
	"Product":     Product,
	"Model":       Model,
}

func TestSimpleFuncs(t *testing.T) {
	for _, lang := range GetLangs() {
		err := SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		for name, funct := range simpleFuncs {
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
