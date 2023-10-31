package ast

type Expr[R any] interface {
	Accept(v ExprVisitor[R]) R
}

type ExprVisitor[R any] interface {
	VisitBinaryExpr(b *binaryExpr[R]) R
	VisitUnaryExpr(u *unaryExpr[R]) R
	VisitLiteralExpr(l *literalExpr[R]) R
	VisitGroupingExpr(g *groupingExpr[R]) R
}
