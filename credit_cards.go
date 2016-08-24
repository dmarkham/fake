package fake

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type creditCard struct {
	vendor   string
	length   int
	prefixes []string
}

// https://en.wikipedia.org/wiki/Payment_card_number#Issuer_identification_number_.28IIN.29
var creditCards = map[string]creditCard{
	"amex":       {"American Express", 15, []string{"34", "37"}},
	"discover":   {"Discover", 16, []string{"6011", "622126", "622925", "644", "649", "65"}},
	"mastercard": {"MasterCard", 16, []string{"5"}},
	"visa":       {"VISA", 16, []string{"4"}},
}

// CreditCardType returns one of the following credit values:
// VISA, MasterCard, American Express and Discover
func CreditCardType() string {
	n := len(creditCards)
	var vendors []string
	for _, cc := range creditCards {
		vendors = append(vendors, cc.vendor)
	}

	return vendors[rand.Intn(n)]
}

// CreditCardNum generated credit card number according to the vendor's card number rules.
// Currently supports amex, discover, mastercard, and visa.
func CreditCardNum(vendor string) string {
	if vendor == "" {
		var vendors []string
		for v := range creditCards {
			vendors = append(vendors, v)
		}
		vendor = vendors[rand.Intn(len(vendors))]
	}
	vendor = strings.ToLower(vendor)

	card := creditCards[vendor]
	prefix := card.prefixes[rand.Intn(len(card.prefixes))]

	num := generateWithPrefix(card.length, prefix)

	return num
}

// Remainder of file adapted from github.com/joeljunstrom/go-luhn
// License: WTFPL

// generateWithPrefix creates and returns a string of the length of the argument targetSize
// but prefixed with the second argument.
// The returned string is valid according to the Luhn algorithm.
func generateWithPrefix(size int, prefix string) string {
	size = size - 1 - len(prefix)

	random := prefix + randomString(size)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

func randomString(size int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	source := make([]int, size)

	for i := 0; i < size; i++ {
		source[i] = rand.Intn(9)
	}

	result := make([]string, len(source))

	for i, number := range source {
		result[i] = strconv.Itoa(number)
	}

	return strings.Join(result, "")
}

func generateControlDigit(luhnString string) int {
	controlDigit := calculateChecksum(luhnString, true) % 10

	if controlDigit != 0 {
		controlDigit = 10 - controlDigit
	}

	return controlDigit
}

func calculateChecksum(luhnString string, double bool) int {
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}
