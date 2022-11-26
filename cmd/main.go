/*
Curr is a cli tool that accepts currency amounts (represented as floating point values)
and types (represented as strings), and returns human language string representations.

Usage:

	currency-cli [flags]

Flags:

	-amount float
		amount to translate (default 500)
	-currency string
		currency type (default "USD")
*/
package main

import (
	"flag"
	"fmt"
	c "github.com/nmiano1111/curr/currency"
)

func main() {
	amount := flag.Float64("amount", 500.00, "amount to translate")
	currency := flag.String("currency", "USD", "currency type")
	flag.Parse()

	curr, err := c.GetCurrencyStringFromNumber(*amount, *currency)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(curr)
}
