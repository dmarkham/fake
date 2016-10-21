package fake

import "math/rand"

//Choose selects a random item from the given set
func Choose(v ...interface{}) interface{} {
	return v[rand.Intn(len(v))]
}

//Toss returns either a true or false
func Toss() bool {
	return Choose(true, false).(bool)
}

//MaritalStatus returns its namesake
func MaritalStatus() string {
	return Choose("Single", "Married", "Separated", "Divorced", "Widow", "Widower").(string)
}

func Ethnicity() string {
	return Choose("Asian", "Black", "White").(string)
}
