package fake

import "testing"

func TestCreditCardType(t *testing.T) {
	t.Parallel()

	v := CreditCardType()
	if v == "" {
		t.Error("CreditCardType failed.")
	}
}

func TestCreditCardNumRandom(t *testing.T) {
	t.Parallel()

	v := CreditCardNum("")
	if v == "" {
		t.Error("CreditCardNum failed to generate random card number.")
	}
}

func TestCreditCardNumVisa(t *testing.T) {
	t.Parallel()

	v := CreditCardNum("visa")
	if v == "" {
		t.Error("CreditCardNum failed to generate visa card number.")
	}
}
