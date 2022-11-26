package currency

import (
	"testing"
)

func TestErrors(t *testing.T) {
	_, err := GetCurrencyStringFromNumber(-1.0, "USD")
	if err == nil {
		t.Errorf("attempting to convert number < 0 should return error")
	}

	_, err = GetCurrencyStringFromNumber(1000000000.0, "USD")
	if err == nil {
		t.Errorf("attempting to convert number > 999,999,999.999 should return error")
	}

	_, err = GetCurrencyStringFromNumber(100.0, "LIRA")
	if err == nil {
		t.Errorf("attempting to use an unsupported currency should return error")
	}
}

func TestMillions(t *testing.T) {
	// -- hundred

	expected := "Four hundred three million two hundred ten thousand thirty three dollars and four cents"
	actual, err := GetCurrencyStringFromNumber(403210033.04, "USD")

	if err != nil {
		t.Errorf("should not return error, got %v", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	// -- ten

	expected = "Fifty six million four hundred nineteen thousand three hundred eighty " +
		"four euros and fourteen cents"
	actual, err = GetCurrencyStringFromNumber(56419384.14, "EURO")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	// -- one

	expected = "One million one hundred two thousand three hundred fifty kronor and ninety two öre"
	actual, err = GetCurrencyStringFromNumber(1102350.92, "SEK")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

func TestThousands(t *testing.T) {
	// -- hundred

	expected := "Nine hundred forty eight thousand ten reais and ninety centavos"
	actual, err := GetCurrencyStringFromNumber(948010.90, "BRL")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	// -- ten

	expected = "Sixty four thousand fifteen dollars and twelve cents"
	actual, err = GetCurrencyStringFromNumber(64015.12, "USD")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	// -- ten

	expected = "Two thousand eight hundred twenty one dollars and forty six cents"
	actual, err = GetCurrencyStringFromNumber(2821.46, "USD")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

}

func TestTens(t *testing.T) {
	expected := "Eighty four dollars and fourteen cents"
	actual, err := GetCurrencyStringFromNumber(84.14, "USD")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	expected = "Nineteen dollars and twelve cents"
	actual, err = GetCurrencyStringFromNumber(19.12, "USD")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

func TestOnes(t *testing.T) {
	expected := "Five dollars and thirty three cents"
	actual, err := GetCurrencyStringFromNumber(5.33, "USD")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	expected = "One dollar and one cent"
	actual, err = GetCurrencyStringFromNumber(1.01, "USD")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	expected = "One euro and one cent"
	actual, err = GetCurrencyStringFromNumber(1.01, "EURO")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	expected = "One krona and one öre"
	actual, err = GetCurrencyStringFromNumber(1.01, "SEK")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

	expected = "One real and one centavo"
	actual, err = GetCurrencyStringFromNumber(1.01, "BRL")
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

func TestCurrencyTypes(t *testing.T) {
	_, err := GetCurrencyStringFromNumber(1102350.92, "USD")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}

	_, err = GetCurrencyStringFromNumber(1102350.92, "EURO")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}

	_, err = GetCurrencyStringFromNumber(1102350.92, "SEK")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}

	_, err = GetCurrencyStringFromNumber(1102350.92, "BRL")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}
}
