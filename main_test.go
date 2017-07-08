package main

import (
	"testing"

	"errors"
	"fmt"
)

func Test_calculate(t *testing.T) {
	type testcase struct {
		input  string
		result int
		err    error
	}

	testcases := []testcase{
		testcase{"", 0, errors.New("")},
		testcase{"+", 0, errors.New("")},
		testcase{"1 +", 0, errors.New("")},
		testcase{"1+", 0, errors.New("")},
		testcase{"12+", 0, errors.New("")},
		testcase{"1 2+", 0, errors.New("")},
		testcase{"1 2 ++", 0, errors.New("")},
		testcase{"1", 1, nil},
		testcase{"1 2", 2, nil},
		testcase{"  1 2  ", 2, nil},
		testcase{"1 2 +", 3, nil},
		testcase{"1 2 3 +", 5, nil},
		testcase{"1 2 3 + +", 6, nil},
		testcase{"2 1 -", 1, nil},
		testcase{"1 2 -", -1, nil},
		testcase{"2 3 *", 6, nil},
		testcase{"6 3 /", 2, nil},
		testcase{"3 6 /", 0, nil},
		testcase{"9 2 /", 4, nil},
	}

	for _, j := range testcases {
		r, err := calculate(j.input)
		if r != j.result || fmt.Sprintf("%T", err) != fmt.Sprintf("%T", j.err) {
			t.Errorf("Case: %s \n"+
				"expected: %d, %v \n"+
				"got: %d, %v",
				j.input,
				j.result, fmt.Sprintf("%T", j.err),
				r, fmt.Sprintf("%T", j.err))
		}
	}
}
