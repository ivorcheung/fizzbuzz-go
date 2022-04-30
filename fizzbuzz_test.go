package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFizzBuzz(t *testing.T) {
	got := 15
	want := "Fizzbuzz"

	assert.Equal(t, fizzbuzz(got), want)
}
