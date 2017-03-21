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
		num          []int
		controlDigit int
	}{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 6},
		{[]int{8, 1, 1, 2, 1, 8, 9, 8, 7}, 6},
		{[]int{6, 5, 4, 1, 6, 6, 3, 4, 0, 8, 2, 7, 3, 1, 3}, 2},
		{[]int{3, 4, 0, 8, 7, 7, 7, 3, 1, 0, 8, 5, 8, 0}, 9},
	}

	for _, cc := range validCCs {
		i := generateControlDigit(cc.num)
		if i != cc.controlDigit {
			t.Errorf("CreditCardNum Luhn check failed for %s, exected %d, got %d", cc.num, cc.controlDigit, i)
		}
	}

}
