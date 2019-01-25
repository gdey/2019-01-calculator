package calculator_test

import (
	"testing"

	calculator "github.com/sdgophers/2019-01-calculator"
)

func TestCalculator(t *testing.T) {

	type TestCase struct {
		e   float64
		err error
	}

	fn := func(input string, tc TestCase) func(t *testing.T) {
		return func(t *testing.T) {
			got, err := calculator.Eval(input)
			if tc.err != err {
				t.Errorf("err, expected %v got %v", tc.err, err)
			}
			// if we expect and error expected is not valid
			if tc.err != nil {
				return
			}
			if got != tc.e {
				t.Errorf("number, expected %v got %v", tc.e, got)
			}
		}
	}

	tests := map[string]TestCase{
		"+ 1 2":         TestCase{e: 3},
		"+ 1 + 1 1":     TestCase{e: 3},
		"+ 2 * 5 8":     TestCase{e: 42},
		"+ 1.5 + 4.6 1": TestCase{e: 7.1},
		"! 2":           TestCase{e: 2},
	}

	for name, tc := range tests {
		t.Run(name, fn(name, tc))
	}
}
