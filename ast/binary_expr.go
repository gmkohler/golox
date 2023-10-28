package ast

type binaryExpr[R any] struct {
	left     Expr[R]
	operator Token
	right    Expr[R]
}

func (b *binaryExpr[R]) Accept(v ExprVisitor[R]) R {
	return v.VisitBinaryExpr(b)
}
