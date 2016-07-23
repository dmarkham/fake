package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

func TestCreditCardType(t *testing.T) {
	t.Parallel()

	v := fake.CreditCardType()
	if v == "" {
		t.Error("CreditCardType failed.")
	}
}

func TestCreditCardNumRandom(t *testing.T) {
	t.Parallel()

	v := fake.CreditCardNum("")
	if v == "" {
		t.Error("CreditCardNum failed to generate random card number.")
	}
}

func TestCreditCardNumVisa(t *testing.T) {
	t.Parallel()

	v := fake.CreditCardNum("visa")
	if v == "" {
		t.Error("CreditCardNum failed to generate visa card number.")
	}
}
