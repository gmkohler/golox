package ast

type visitor[R, E Expr[R]] interface {
	visit(expr E) R
}
