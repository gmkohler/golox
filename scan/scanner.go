package scan

import "github.com/gmkohler/golox/ast"

type Scanner interface {
	ScanTokens() ([]*ast.Token, error)
}
