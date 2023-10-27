package golox

import (
	"slices"
	"testing"
)

type testCase struct {
	src      string
	expected []Token
}

var testCases = []testCase{
	{
		src: "var language = \"lox\"",
		expected: []Token{
			{
				tType:  tokenTypeVar,
				lexeme: "var",
				line:   1,
			},
			{
				tType:  tokenTypeIdentifier,
				lexeme: "language",
				line:   1,
			},
			{
				tType:  tokenTypeEqual,
				lexeme: "=",
				line:   1,
			},
			{
				tType:   tokenTypeString,
				lexeme:  "\"lox\"",
				literal: "lox",
				line:    1,
			},
			{
				tType: tokenTypeEof,
				line:  1,
			},
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
		if !slices.Equal(tokens, tc.expected) {
			t.Fatalf("\nExpected: %+v\nGot: %+v", tc.expected, tokens)
		}
	}

}
