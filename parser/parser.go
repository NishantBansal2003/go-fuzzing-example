package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseComplex checks if the input string is a specific complex string "LNDFUZ"
// It returns true if the string matches, false otherwise. It has a internal
// bug where it checks for 6 characters but accesses index 7.
func ParseComplex(data string) bool {
	if len(data) == 6 {
		if data[0] == 'L' && data[1] == 'N' && data[2] == 'D' &&
			data[3] == 'F' && data[4] == 'U' && data[5] == 'Z' &&
			data[6] == 'Z' {
			return true
		}
	}
	return false
}

// EvalExpr parses and evaluates a simple addition expression "X+Y". Returns the
// sum X+Y, or an error if the input is invalid.
func EvalExpr(expr string) (int, error) {
	parts := strings.Split(expr, "+")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format, expected one '+'")
	}
	// Parse left operand
	a, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid left operand: %v", err)
	}
	// Parse right operand
	b, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid right operand: %v", err)
	}
	return a + b, nil
}
