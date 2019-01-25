package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func fac(i int64) int64 {
	if i == 0 || i == 1 {
		return 1
	}
	return i * fac(i-1)
}

func pop2(stack []string) ([2]int64, []string, error) {
	var (
		el     [2]int64
		err    error
		stack1 []string
		stack2 []string
	)
	el[0], stack1, err = pop(stack)
	if err != nil {
		return el, stack, err
	}
	el[1], stack2, err = pop(stack1)
	if err != nil {
		return el, stack, err
	}
	return el, stack2, nil
}

// pop will pop of the stack a value evaluting it till it reaches a number
func pop(stack []string) (int64, []string, error) {
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
		return int64(math.Pow(float64(el[0]), float64(el[1]))), stack, nil
	case "%":
		el, stack, err := pop2(tail)
		if err != nil {
			return 0, stack, err
		}
		return el[0] % el[1], stack, nil

	case "!":
		el, stack, err := pop(tail)
		if err != nil {
			return 0, stack, err
		}
		return fac(el), stack, nil

	default: // assume number
		i, err := strconv.ParseInt(stack[0], 10, 64)
		return i, tail, err
	}
}

func Eval(s string) (int64, error) {
	i, _, err := pop(strings.Split(s, " "))
	return i, err
}
