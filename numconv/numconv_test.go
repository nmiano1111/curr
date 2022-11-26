package numconv

import (
	"fmt"
	"strings"
	"testing"
)

func TestIntToStrings(t *testing.T) {
	x := 123456789

	m := func(ns []int) (string, error) {
		builder := strings.Builder{}
		for _, n := range ns {
			builder.WriteString(fmt.Sprintf("%d", n))
		}
		return builder.String(), nil
	}

	expected := []string{"123", "456", "789"}
	actual, err := IntToStrings(x, m)
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if !equalStringSlice(expected, actual) {
		t.Errorf("expected '%v', got '%v'", expected, actual)
	}

	// odd
	x = 23456789
	expected = []string{"23", "456", "789"}
	actual, err = IntToStrings(x, m)
	if err != nil {
		t.Errorf("should not return error, got %e", err)
	} else if !equalStringSlice(expected, actual) {
		t.Errorf("expected '%v', got '%v'", expected, actual)
	}
}

func TestIntToStringsError(t *testing.T) {
	x := 12345678

	m := func(ns []int) (string, error) {
		return "", fmt.Errorf("oops")
	}

	// -- error on first pass

	_, err := IntToStrings(x, m)
	if err == nil {
		t.Errorf("when mapper returns error, so should function")
	}

	// -- error at end

	m = func(ns []int) (string, error) {
		if ns[0] == 1 {
			return "", fmt.Errorf("oops")
		}
		return "", nil
	}

	_, err = IntToStrings(x, m)
	if err == nil {
		t.Errorf("when mapper returns error, so should function")
	}
}

// helper  to compare slices of strings
func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
