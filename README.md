# curr

Curr is a cli tool that accepts currency amounts (represented as floating point values)
and types (represented as strings), and returns human language string representations of the given
currency value and type. It is written with the [Go](https://go.dev/) programming language.

####Usage:

	curr-cli [flags]

####Flags:

	-amount float
		amount to translate (default 500)
	-currency string
		currency type (default "USD")


For example: `curr-cli --amount 1234.12 --currency EURO`
produces `One thousand two hundred thirty four euros and twelve cents` as output.

Example of code usage:
```go
curr, err := GetCurrencyStringFromNumber(1234.12, "EURO")
```

Curr accepts `USD` (US Dollar), `EURO` (Euro), `SEK` (Swedish Krona), `BRL` (Brazilian Real) as currency types.

##Build

To compile an executable, run `go build -o curr-cli  cmd/main.go`.

_note: a precompiled executable named `curr-cli` is included in the repo, in case you
don't have Go set up on your machine._

##Test

Tests can be run with the following command: `go test ./...`

### Project Layout
* `cmd` directory contains logic for CLI
* `currency` contains currency related functions / types / variables / constants
* `numconv` contains logic for transforming numerical data


