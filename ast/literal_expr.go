package ast

type literalExpr[R any] struct {
	value any
}

func (l *literalExpr[R]) Accept(v ExprVisitor[R]) R {
	return v.VisitLiteralExpr(l)
}
