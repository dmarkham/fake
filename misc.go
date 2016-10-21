package fake

import "math/rand"

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
