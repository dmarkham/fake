package fake

import (
		"fmt"
		"math/rand"
	)

//Choose selects a random item from the given set
func Choose(v ...interface{}) interface{} {
	return v[rand.Intn(len(v))]
}

//Toss returns either a true or false
func Toss() int {
	return rand.Intn(2)
}

func Bool() bool {
	return Choose(true, false).(bool)
}

//MaritalStatus returns its namesake
func MaritalStatus() string {
	return Choose("Single", "Married", "Separated", "Divorced", "Widow", "Widower").(string)
}

func SexualOrientation() string {
	return Choose("Straight", "Gay", "Lesbian", "Bisexual", "Other").(string)
}

func Education() string {
	return Choose("None", "Home Schooled", "High School", "Diploma", "Associate Degree", "Bachelors", "Masters", "Doctorate", "Other").(string)
}

func Ethnicity() string {
	return Choose("Asian", "Black", "White", "Other").(string)
}

func Ethnicities(n int) (a []string) {
	for i := 0; i < n; i++ {
		a = append(a, Ethnicity())
	}
	return a
}

func Religion() string {
	return Choose("Hindu", "Atheist", "Islam", "Christian", "Jew", "Other").(string)
}

func Religions(n int) (a []string) {
	for i := 0; i < n; i++ {
		a = append(a, Religion())
	}
	return a
}

func Question() string {
	return Sentence() + "?"
}

func Amount() float32 {
	return rand.Float32()
}

func Between(from, to int) int {
	return rand.Intn(to-from) + from
}

func Version() string {
	ma := rand.Intn(10)
	mi := rand.Intn(100)
	dot := rand.Intn(100)
	return fmt.Sprintf("%v.%v.%v", ma, mi, dot)
}
