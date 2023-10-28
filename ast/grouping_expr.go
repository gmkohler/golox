package ast

type groupingExpr[R any] struct {
	expression Expr[R]
}

func (g *groupingExpr[R]) Accept(v ExprVisitor[R]) R {
	return v.VisitGroupingExpr(g)
}
