package calculator_test

import (
	"meetup/2019-01-calculator"
	"testing"
)

func TestASTCalculator(t *testing.T) {

	type TestCase struct {
		e   float64
		err error
	}

	fn := func(input string, tc TestCase) func(t *testing.T) {
		return func(t *testing.T) {
			got, _  := calculator.EvalAsAST(input)
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
