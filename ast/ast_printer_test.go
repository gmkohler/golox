package ast

import (
	"testing"
)

type testCase struct {
	expr     Expr[string]
	expected string
}

var expr Expr[string] = &binaryExpr[string]{ // -123 * (45.67)
	left: &unaryExpr[string]{
		token: *NewToken(TokenTypeMinus, "-", 1),
		right: &literalExpr[string]{
			value: 123,
		},
	},
	operator: *NewToken(TokenTypeStar, "*", 1),
	right: &groupingExpr[string]{
		expression: &literalExpr[string]{
			value: 45.67,
		},
	},
}

var testCases = []testCase{
	{
		expr:     expr,
		expected: "(* (- 123) (group 45.67))",
	},
}

func TestAstPrinter_Print(t *testing.T) {
	var p = new(astPrinter)
	for _, tc := range testCases {
		var actual = p.Print(tc.expr)
		if actual != tc.expected {
			t.Fatalf("expected %s, got %s", tc.expected, actual)
		}
	}
}
