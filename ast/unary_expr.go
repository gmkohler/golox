package ast

type unaryExpr[R any] struct {
	token Token
	right Expr[R]
}

func (u *unaryExpr[R]) Accept(v ExprVisitor[R]) R {
	return v.VisitUnaryExpr(u)
}
