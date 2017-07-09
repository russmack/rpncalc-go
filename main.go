// Package rpncalc-go is a reverse polish notation calculator.
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// usage displays the correct program usage.
func usage() {
	fmt.Println(`Usage: rpncalc-go "182 7 27 2 3 4 5 + + + - * /"`)
	os.Exit(1)
}

// pop takes an array, removes the last value, and returns both
// the last value and the shortened array.
func pop(v []int) (int, []int, error) {
	if len(v) == 0 {
		return 0, nil, fmt.Errorf("Syntax error: too few values.")
	}
	return v[len(v)-1], v[:len(v)-1], nil
}

// main iterates over the values and operators,
// adding values to the stack, and processing the stack values
// according to the operators.
func main() {
	if len(os.Args) < 2 {
		usage()
	}

	s := strings.TrimSpace(os.Args[1])

	fmt.Println("Calculating:", s)

	r, err := calculate(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Result:", r)
}

func calculate(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("Syntax error: no values.")
	}

	stack := []int{}

	args := strings.Split(s, " ")

	for _, j := range args {
		n, err := strconv.Atoi(j)
		if err == nil {
			stack = append(stack, n)
			continue
		}

		var a, b int

		a, stack, err = pop(stack)
		if err != nil {
			return 0, err
		}

		b, stack, err = pop(stack)
		if err != nil {
			return 0, err
		}

		switch j {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, b-a)
		case "*":
			stack = append(stack, a*b)
		case "/":
			stack = append(stack, b/a)
		default:
			return 0, fmt.Errorf("Syntax error: %s", j)
		}
	}

	return stack[len(stack)-1], nil
}
