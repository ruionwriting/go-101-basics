package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_FizzBuzz(t *testing.T) {
	testCases := map[int]string{
		1:  "1",
		2:  "2",
		3:  "Fizz",
		5:  "Buzz",
		6:  "Fizz",
		9:  "Fizz",
		10: "Buzz",
		15: "FizzBuzz",
	}

	for k, v := range testCases {
		t.Run(fmt.Sprintf("FizzBuzz(%v->%v)", k, v), func(t *testing.T) {
			assert(t, v, FizzBuzz(k))
		})
	}
}

func Benchmark_Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(3)
	}
}

// assert is a helper function that takes an expected
// value and an actual value, performs a deep compare
// and if they're different, fails the test
func assert(tb testing.TB, expected interface{}, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	tb.Fatalf("expected %v but got %v", expected, actual)
}
