package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

func TestGeoLongLat(t *testing.T) {
	t.Parallel()

	f := fake.Latitute()
	if f == 0 {
		t.Error("Latitude zeroed")
	}

	f = fake.Longitude()
	if f == 0 {
		t.Error("Longitude zeroed")
	}
}

func TestGeoMinSecs(t *testing.T) {
	t.Parallel()

	i := fake.LatitudeMinutes()
	if i < 0 || i >= 60 {
		t.Errorf("LatitudeMinutes failed, got %v", i)
	}

	i = fake.LatitudeSeconds()
	if i < 0 || i >= 60 {
		t.Errorf("LatitudeSeconds failed, got %v", i)
	}
	i = fake.LongitudeMinutes()
	if i < 0 || i >= 60 {
		t.Errorf("LongitudeMinutes failed, got %v", i)
	}

	i = fake.LongitudeSeconds()
	if i < 0 || i >= 60 {
		t.Errorf("LongitudeSeconds failed, got %v", i)
	}
}

func TestGeoDegrees(t *testing.T) {
	t.Parallel()

	i := fake.LatitudeDegrees()
	if i < -180 || i > 180 {
		t.Errorf("LatitudeDegrees failed, got %v", i)
	}

	i = fake.LongitudeDegrees()
	if i < -180 || i > 180 {
		t.Errorf("LongitudeDegrees failed, got %v", i)
	}

}

func TestGeoDirection(t *testing.T) {
	t.Parallel()

	s := fake.LatitudeDirection()
	if s != "N" && s != "S" {
		t.Errorf("LatitudeDirection failed, got %v", s)
	}

	s = fake.LongitudeDirection()
	if s != "W" && s != "E" {
		t.Errorf("LongitudeDirection failed, got %v", s)
	}
}
