package scan

import (
	"github.com/gmkohler/golox/ast"
	"slices"
	"testing"
)

type testCase struct {
	src      string
	expected []*ast.Token
}

var testCases = []testCase{
	{
		src: "var language = \"lox\"",
		expected: []*ast.Token{
			ast.NewToken(ast.TokenTypeVar, "var", 1),
			ast.NewToken(ast.TokenTypeIdentifier, "language", 1),
			ast.NewToken(ast.TokenTypeEqual, "=", 1),
			ast.NewTokenWithLiteral(ast.TokenTypeString, "\"lox\"", "lox", 1),
			ast.NewToken(ast.TokenTypeEof, "", 1),
		},
	},
}

func TestSourceScanner_ScanTokens(t *testing.T) {
	for _, tc := range testCases {
		s := NewSourceScanner(tc.src)
		tokens, err := s.ScanTokens()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !slices.EqualFunc(
			tokens,
			tc.expected,
			func(t1 *ast.Token, t2 *ast.Token) bool {
				return *t1 == *t2
			},
		) {
			t.Fatalf("\nExpected: %+v\nGot: %+v", tc.expected, tokens)
		}
	}

}
