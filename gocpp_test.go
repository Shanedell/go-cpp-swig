package main

import (
	"testing"

	"github.com/Shanedell/go-cpp-swig/gocpp"
)

func TestSum(t *testing.T) {
	expected := 6
	given := gocpp.Sum(4, 2)

	if expected != given {
		t.Fatalf(`Expected: %v != Given: %v`, expected, given)
	}
}

func TestDiff(t *testing.T) {
	expected := 2
	given := gocpp.Diff(4, 2)

	if expected != given {
		t.Fatalf(`Expected: %v != Given: %v`, expected, given)
	}
}

func TestProduct(t *testing.T) {
	expected := 8
	given := gocpp.Product(4, 2)

	if expected != given {
		t.Fatalf(`Expected: %v != Given: %v`, expected, given)
	}
}

func TestQuotient(t *testing.T) {
	expected := 2
	given := gocpp.Quotient(4, 2)

	if expected != given {
		t.Fatalf(`Expected: %v != Given: %v`, expected, given)
	}
}

func TestModulus(t *testing.T) {
	expected := 0
	given := gocpp.Modulus(4, 2)

	if expected != given {
		t.Fatalf(`Expected: %v != Given: %v`, expected, given)
	}
}
