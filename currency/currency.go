// Package currency implements functionality to transform floating point values
// along with currency type strings into human language strings
package currency

import (
	"fmt"
	"github.com/nmiano1111/curr/numconv"
	"math"
	"strings"
)

// Maximum and minimum accepted values
const (
	Max = 999999999.999
	Min = 0.0
)

var (
	// Accepted currencies
	currencies = map[string]Currency{
		"USD":  {"dollar", "dollars", "cent", "cents"},
		"EURO": {"euro", "euros", "cent", "cents"},
		"SEK":  {"krona", "kronor", "öre", "öre"},
		"BRL":  {"real", "reais", "centavo", "centavos"},
	}

	// Atomic number words used as building blocks for currency strings
	nums = []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens = []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	magnitudes = []string{
		"hundred",
		"thousand",
		"million",
	}
)

// Currency holds single and plural currency names,
// as well as single and plural 'partial' names (e.g. cents)
type Currency struct {
	name       string
	pluralName string

	partial       string
	pluralPartial string
}

func (c Currency) Name(amount float64) string {
	if amount == 1.0 {
		return c.name
	} else {
		return c.pluralName
	}
}

func (c Currency) Partial(amount float64) string {
	if amount == 1.0 {
		return c.partial
	} else {
		return c.pluralPartial
	}
}

// GetCurrencyStringFromNumber accepts a value x and currency type currency, and returns
// a human language string of words representing the currency amount and type.
// (e.g. 123.12 -> One hundred twenty three dollars and twelve cents)
func GetCurrencyStringFromNumber(x float64, currency string) (words string, _ error) {
	if x > Max || x < Min {
		return words, fmt.Errorf("%.3f is < %.3f or > %.3f", x, Max, Min)
	}

	translatedCurr, ok := currencies[currency]
	if !ok {
		return words, fmt.Errorf("%s is not a supported currency", currency)
	}

	integer, fractional := math.Modf(x)
	fractional = math.Round(fractional * 100)

	dollar, err := numToWords(int(integer))
	if err != nil {
		return words, err
	}

	cents, err := numToWords(int(fractional))
	if err != nil {
		return words, err
	}

	dollarName := translatedCurr.Name(integer)
	centName := translatedCurr.Partial(fractional)

	words = fmt.Sprintf("%s %s and %s %s", dollar, dollarName, cents, centName)
	return strings.ToUpper(words[:1]) + words[1:], nil
}

// numToWords accepts an integer value x, and
// returns a string of words
func numToWords(x int) (string, error) {
	if x < 20 {
		return nums[x], nil
	}

	chunks, err := numconv.IntToStrings(x, mapNumsToText)
	if err != nil {
		return "", err
	}

	switch {
	case len(chunks) == 3:
		chunks[0] = fmt.Sprintf("%s %s", chunks[0], magnitudes[2])
		chunks[1] = fmt.Sprintf("%s %s", chunks[1], magnitudes[1])
	case len(chunks) == 2:
		chunks[0] = fmt.Sprintf("%s %s", chunks[0], magnitudes[1])
	}
	return strings.Join(chunks, " "), nil
}

// mapNumsToText accepts a slice of ints s,
// and returns a string of those numbers as words
func mapNumsToText(s []int) (string, error) {
	if len(s) > 3 || len(s) == 0 {
		return "", fmt.Errorf("length of slice must be between 1 - 3, got %d", len(s))
	}

	// little helper to cut down on clutter
	buildTensFunc := func(a, b int, builder *strings.Builder) {
		combined := a*10 + b
		if combined < 20 {
			builder.WriteString(fmt.Sprintf("%s", nums[combined]))
		} else if b > 0 {
			builder.WriteString(fmt.Sprintf("%s %s", tens[a], nums[b]))
		} else {
			builder.WriteString(tens[a])
		}
	}

	builder := strings.Builder{}

	switch {
	case len(s) == 3:
		if s[0] > 0 {
			builder.WriteString(fmt.Sprintf("%s %s ", nums[s[0]], magnitudes[0]))
		}
		buildTensFunc(s[1], s[2], &builder)
	case len(s) == 2:
		buildTensFunc(s[0], s[1], &builder)
	case len(s) == 1:
		builder.WriteString(fmt.Sprintf("%s", nums[s[0]]))
	}
	return builder.String(), nil
}
