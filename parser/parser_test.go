package parser

import (
	"strconv"
	"strings"
	"testing"
)

// FuzzParseComplex fuzzes the ParseComplex function.
func FuzzParseComplex(f *testing.F) {
	f.Fuzz(func(t *testing.T, data string) {
		_ = ParseComplex(data)
	})
}

// FuzzEvalExpr fuzzes the EvalExpr function. It should never panic on any input
func FuzzEvalExpr(f *testing.F) {
	f.Add("1+3")
	f.Add("-2+3")

	f.Fuzz(func(t *testing.T, expr string) {
		result, err := EvalExpr(expr)
		if err != nil {
			return
		}
		// If parsing succeeded, verify the result for validity
		parts := strings.Split(expr, "+")
		expectedA, _ := strconv.Atoi(parts[0])
		expectedB, _ := strconv.Atoi(parts[1])
		if result != expectedA+expectedB {
			t.Errorf("unexpected sum for %q: got %d, want %d", expr,
				result, expectedA+expectedB)
		}
	})
}
