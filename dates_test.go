package fake

import (
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	t.Parallel()

	n := Day()
	if n < 0 || n > 31 {
		t.Errorf("Day failed, got %d", n)
	}
}

func TestMonthNum(t *testing.T) {
	t.Parallel()

	n := MonthNum()
	if n < 0 || n > 31 {
		t.Errorf("MonthNum failed, got %d", n)
	}
}

func TestWeekdayNum(t *testing.T) {
	t.Parallel()

	n := WeekdayNum()
	if n < 0 || n > 7 {
		t.Errorf("WeekdayNum failed, got %d", n)
	}
}

func TestYear(t *testing.T) {
	t.Parallel()

	n := Year(1950, 2020)
	if n < 1950 || n > 2020 {
		t.Errorf("Year failed, got %d", n)
	}
}

func TestTime(t *testing.T) {
	t.Parallel()

	from, _ := time.Parse("2006-01-02T15:04:05", "2016-01-01T00:00:00")
	to, _ := time.Parse("2006-01-02T15:04:05", "2016-01-31T23:59:59")
	d := Time(from, to)
	if !from.Equal(d) && !from.Before(d) {
		t.Errorf("Expected time from %s, got %s", from, d)
	}
	if !to.Equal(d) && !to.After(d) {
		t.Errorf("Expected time to %s, got %s", to, d)
	}
}

func TestBirthdate(t *testing.T) {
	t.Parallel()

	for age := 0; age <= 120; age++ {
		now := time.Now()
		birthdate := Birthdate(age)
		diffDate := now.Sub(birthdate)
		calcAge := int(diffDate.Hours() / (365 * 24))
		diff := calcAge - age
		// TODO improve this to be more precise. Right now we don't need precision
		if diff > 1 {
			t.Errorf("Birthdate (%s) is not valid according to age (%d) = %d", birthdate, age, diff)
		}
	}
}
