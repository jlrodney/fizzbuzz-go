package main

import(
  "testing"
)

func TestFizzBuzz(t *testing.T) {
  result := fizzbuzz(15)
  expect := "FizzBuzz"
  if result != expect {
  t.Error("Test failed: 45 inputted, expected: FizzBuzz")
  }
}

func TestBuzz(t *testing.T) {
  result := fizzbuzz(10)
  expect := "Buzz"
  if result != expect {
  t.Error("Test failed: 10 inputted, expected: Buzz")
  }
}

func TestFizz(t *testing.T) {
  result := fizzbuzz(9)
  expect := "Fizz"
  if result != expect {
  t.Error("Test failed: 9 inputted, expected: Buzz")
  }
}

func TestInt(t *testing.T) {
  result := fizzbuzz(4)
  expect := "4"
  if result != expect {
  t.Error("Test failed: 4 inputted, expected: 4")
  }
}

func TestFizzBuzz100(t *testing.T) {
	var testCases = []struct {
		integer  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{6, "Fizz"},
		{9, "Fizz"},
		{8, "8"},
		{10, "Buzz"},
		{12, "Fizz"},
		{15, "FizzBuzz"},
		{25, "Buzz"},
		{30, "FizzBuzz"},
	}

	for _, test := range testCases {
		result := fizzbuzz(test.integer)
		expect := test.out

		if result != expect {
			t.Errorf("Test failed at input: %v \n  Actual result: %v \n  Expected: %v", test.integer, result, expect)
		}
	}
}
