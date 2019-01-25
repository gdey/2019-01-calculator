package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func fac(i float64) float64 {
	if i == 0 || i == 1 {
		return 1
	}
	return i * fac(i-1)
}

func pop2(stack []string) ([2]float64, []string, error) {
	var (
		el  [2]float64
		err error
	)
	el[0], stack, err = pop(stack)
	if err != nil {
		return el, stack, err
	}
	el[1], stack, err = pop(stack)
	if err != nil {
		return el, stack, err
	}
	return el, stack, nil
}

// pop will pop of the stack a value evaluting it till it reaches a number
func pop(stack []string) (float64, []string, error) {
	head, tail := stack[0], stack[1:]
	if len(stack) == 0 {
		return 0, stack, errors.New("Not enough values to eval")
	}
	switch strings.ToLower(head) {
	case "+":
		el, stack, err := pop2(tail)
		if err != nil {
			return 0, stack, err
		}
		return el[0] + el[1], stack, nil
	case "-":
		el, stack, err := pop2(tail)
		if err != nil {
			return 0, stack, err
		}
		return el[0] - el[1], stack, nil
	case "*":
		el, stack, err := pop2(tail)
		if err != nil {
			return 0, stack, err
		}
		return el[0] * el[1], stack, nil
	case "/":
		el, stack, err := pop2(tail)
		if err != nil {
			return 0, stack, err
		}
		return el[0] / el[1], stack, nil
	case "^":
		el, stack, err := pop2(tail)
		if err != nil {
			return 0, stack, err
		}
		return math.Pow(el[0], el[1]), stack, nil
	case "%":
		el, stack, err := pop2(tail)
		if err != nil {
			return 0, stack, err
		}
		return float64(int64(el[0]) % int64(el[1])), stack, nil

	case "!":
		el, stack, err := pop(tail)
		if err != nil {
			return 0, stack, err
		}
		return fac(el), stack, nil

	default: // assume number
		i, err := strconv.ParseFloat(stack[0], 64)
		return i, tail, err
	}
}

func Eval(s string) (float64, error) {
	i, _, err := pop(strings.Split(s, " "))
	return i, err
}
