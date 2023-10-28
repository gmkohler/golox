package parse

import "github.com/gmkohler/golox/ast"

type parser struct {
	tokens  []*ast.Token
	current int
}
