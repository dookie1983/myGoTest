package fizzbuzz

import "testing"

func TestFizzBuzzOne(t *testing.T) {
	testExpected := "1"
	if observed := FizzBuzz("1"); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", "1", observed, testExpected)
	}
}
func TestFizzBuzzTwo(t *testing.T) {
	testExpected := "2"
	if observed := FizzBuzz("2"); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", "2", observed, testExpected)
	}
}

func TestFizzBuzzTree(t *testing.T) {
	testExpected := "Fizz"
	inputParams := "3"
	if observed := FizzBuzz(inputParams); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", inputParams, observed, testExpected)
	}
}

func TestFizzBuzzFive(t *testing.T) {
	testExpected := "Buzz"
	inputParams := "5"
	if observed := FizzBuzz(inputParams); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", inputParams, observed, testExpected)
	}
}

func TestFizzBuzzSix(t *testing.T) {
	testExpected := "Fizz"
	inputParams := "6"
	if observed := FizzBuzz(inputParams); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", inputParams, observed, testExpected)
	}
}
func TestFizzBuzzNine(t *testing.T) {
	testExpected := "Fizz"
	inputParams := "9"
	if observed := FizzBuzz(inputParams); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", inputParams, observed, testExpected)
	}
}

func TestFizzBuzzTen(t *testing.T) {
	testExpected := "Buzz"
	inputParams := "10"
	if observed := FizzBuzz(inputParams); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", inputParams, observed, testExpected)
	}
}

func TestFizzBuzzTwel(t *testing.T) {
	testExpected := "Fizz"
	inputParams := "12"
	if observed := FizzBuzz(inputParams); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", inputParams, observed, testExpected)
	}
}
func TestFizzBuzzFifteen(t *testing.T) {
	testExpected := "FizzBuzz"
	inputParams := "15"
	if observed := FizzBuzz(inputParams); observed != testExpected {
		t.Fatalf("ShareWith(%s) = %v, want %v", inputParams, observed, testExpected)
	}
}
