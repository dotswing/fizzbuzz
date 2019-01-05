package fizzbuzz_test

import (
	"testing"

	"github.com/dotswing/fizzbuzz"
)

func TestFizzBuzzShouldSayOne(t *testing.T) {

	expected := "1"

	result := fizzbuzz.Say(1)

	if expected != result {
		t.Errorf("it should say %q but go %q", expected, result)
	}
}

func TestFizzBuzzShouldSayTwo(t *testing.T) {
	expected := "2"

	result := fizzbuzz.Say(2)

	if expected != result {
		t.Errorf("it should say %q but got %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFizz(t *testing.T) {
	expected := "Fizz"

	result := fizzbuzz.Say(3)

	if expected != result {
		t.Errorf("it should say %q but got %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFour(t *testing.T) {
	expected := "4"

	result := fizzbuzz.Say(4)

	if expected != result {
		t.Errorf("it should say %q but got %q", expected, result)
	}
}

func TestFizzBuzzShouldSayBuzz(t *testing.T) {
	expected := "Buzz"

	result := fizzbuzz.Say(5)

	if expected != result {
		t.Errorf("it should say %q but got %q", expected, result)
	}
}
