package fizzbuzz

import (
	"strconv"
)

func fizzbuzz(input int) string {
	fizzbuzz := "FizzBuzz"
	fizz := "Fizz"
	buzz := "Buzz"

	if input%15 == 0 {
		return fizzbuzz
	} else if input%3 == 0 {
		return fizz
	} else if input%5 == 0 {
		return buzz
	}
	return strconv.Itoa(input)
}
