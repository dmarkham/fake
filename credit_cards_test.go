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

	if len(v) < 15 {
		t.Error("CreditCardNum too short.")
	}
}

func TestCreditCardNumVisa(t *testing.T) {
	t.Parallel()

	v := CreditCardNum("visa")
	if v == "" {
		t.Error("CreditCardNum failed to generate visa card number.")
	}

	if len(v) < 16 {
		t.Error("Visa num too short.")
	}

}

func TestCreditCardNumLuhn(t *testing.T) {
	t.Parallel()

	var validCCs = []struct {
		num          string
		controlDigit int
	}{
		{"111111111", 6},
		{"811218987", 6},
		{"654166340827313", 2},
		{"34087773108580", 9},
	}

	for _, cc := range validCCs {
		i := generateControlDigit(cc.num)
		if i != cc.controlDigit {
			t.Errorf("CreditCardNum Luhn check failed for %s, exected %d, got %d", cc.num, cc.controlDigit, i)
		}
	}

}
