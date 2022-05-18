package fizzbuzz

import (
	"testing"
)

func TestIsFizzBuzz(t *testing.T) {

	testCases := []struct {
		input int
		want  string
	}{
		{0, "FizzBuzz"},
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
	}

	for _, testCase := range testCases {
		fizzBuzzResults := fizzbuzz(testCase.input)
		if fizzBuzzResults != testCase.want {
			t.Errorf("Got %s, expected %s", fizzBuzzResults, testCase.want)
		}
	}

}
