package calculator

import (
	"math"
	"strings"
	"errors"
	"strconv"
)

//AST
type World *struct{}

type ExprNode interface {
	Eval(World) ExprNode
}

type Number float64

func (n Number) Eval(w World) ExprNode {
	return n
}

type Addition [2]ExprNode

func (a Addition) Eval(w World) ExprNode {
	a1 := a[0].Eval(w)
	a2 := a[1].Eval(w)

	v1, ok := a1.(Number)
	if !ok {
		panic("expected number")
	}

	v2, ok := a2.(Number)
	if !ok {
		panic("expected number")
	}

	v := float64(v1) + float64(v2)

	return Number(v)
}

type Subtraction [2]ExprNode

func (s Subtraction) Eval(w World) ExprNode {
	s1 := s[0].Eval(w)
	s2 := s[1].Eval(w)

	v1, ok := s1.(Number)
	if !ok {
		panic("expected number")
	}

	v2, ok := s2.(Number)
	if !ok {
		panic("expected number")
	}

	v := float64(v1) - float64(v2)

	return Number(v)
}

type Division [2]ExprNode

func (a Division) Eval(w World) ExprNode {
	a1 := a[0].Eval(w)
	a2 := a[1].Eval(w)

	v1, ok := a1.(Number)
	if !ok {
		panic("expected number")
	}

	v2, ok := a2.(Number)
	if !ok {
		panic("expected number")
	}

	v := float64(v1) / float64(v2)

	return Number(v)
}

type Multiplication [2]ExprNode

func (s Multiplication) Eval(w World) ExprNode {
	s1 := s[0].Eval(w)
	s2 := s[1].Eval(w)

	v1, ok := s1.(Number)
	if !ok {
		panic("expected number")
	}

	v2, ok := s2.(Number)
	if !ok {
		panic("expected number")
	}

	v := float64(v1) * float64(v2)

	return Number(v)
}

type Exponents [2]ExprNode

func (s Exponents) Eval(w World) ExprNode {
	s1 := s[0].Eval(w)
	s2 := s[1].Eval(w)

	v1, ok := s1.(Number)
	if !ok {
		panic("expected number")
	}

	v2, ok := s2.(Number)
	if !ok {
		panic("expected number")
	}

	v := math.Pow(float64(v1), float64(v2))

	return Number(v)
}

type Modulus [2]ExprNode

func (s Modulus) Eval(w World) ExprNode {
	s1 := s[0].Eval(w)
	s2 := s[1].Eval(w)

	v1, ok := s1.(Number)
	if !ok {
		panic("expected number")
	}

	v2, ok := s2.(Number)
	if !ok {
		panic("expected number")
	}

	v := float64(int64(v1) % int64(v2))

	return Number(v)
}

type Fac struct{ node ExprNode }

func (s Fac) Eval(w World) ExprNode {
	s1 := s.node.Eval(w)

	v1, ok := s1.(Number)
	if !ok {
		panic("expected number")
	}

	v := fac(float64(v1))
	return Number(v)
}

func EvalAsAST(s string) (float64, ExprNode){
	el, _, err := popAST(strings.Split(s, " "))
	if err != nil {
		panic("invalid input")
	}

	v, ok := el.Eval(nil).(Number)
	if ! ok {
		panic("expected a number")

	}

	return float64(v), el

}

// pop will pop of the stack a value evaluting it till it reaches a number
func popAST(stack []string) (ExprNode, []string, error) {
	head, tail := stack[0], stack[1:]
	if len(stack) == 0 {
		return nil, stack, errors.New("Not enough values to eval")
	}
	switch strings.ToLower(head) {
	case "+":
		el, stack, err := pop2AST(tail)
		if err != nil {
			return nil, stack, err
		}
		return Addition(el), stack, nil
	case "-":
		el, stack, err := pop2AST(tail)
		if err != nil {
			return nil, stack, err
		}
		return Subtraction(el), stack, nil
	case "*":
		el, stack, err := pop2AST(tail)
		if err != nil {
			return nil, stack, err
		}
		return Multiplication(el), stack, nil
	case "/":
		el, stack, err := pop2AST(tail)
		if err != nil {
			return nil, stack, err
		}
		return Division(el), stack, nil
	case "^":
		el, stack, err := pop2AST(tail)
		if err != nil {
			return nil, stack, err
		}
		return Exponents(el), stack, nil
	case "%":
		el, stack, err := pop2AST(tail)
		if err != nil {
			return nil, stack, err
		}
		return Modulus(el), stack, nil

	case "!":
		el, stack, err := popAST(tail)
		if err != nil {
			return nil, stack, err
		}
		return Fac{node: el}, stack, nil

	default: // assume number
		i, err := strconv.ParseFloat(stack[0], 64)
		return Number(i), tail, err
	}
}

func pop2AST(stack []string) ([2]ExprNode, []string, error) {
	var (
		el  [2]ExprNode
		err error
	)
	el[0], stack, err = popAST(stack)
	if err != nil {
		return el, stack, err
	}
	el[1], stack, err = popAST(stack)
	if err != nil {
		return el, stack, err
	}
	return el, stack, nil
}
