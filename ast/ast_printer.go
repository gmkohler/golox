package ast

import (
	"fmt"
	"strings"
)

type astPrinter struct{}

func (p *astPrinter) Print(e Expr[string]) string {
	return e.Accept(p)
}

func (p *astPrinter) VisitBinaryExpr(b *binaryExpr[string]) string {
	return p.parenthesize(b.operator.lexeme, b.left, b.right)
}

func (p *astPrinter) VisitUnaryExpr(u *unaryExpr[string]) string {
	return p.parenthesize(u.token.lexeme, u.right)
}

func (p *astPrinter) VisitLiteralExpr(l *literalExpr[string]) string {
	if l.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", l.value)
}

func (p *astPrinter) VisitGroupingExpr(g *groupingExpr[string]) string {
	return p.parenthesize("group", g.expression)
}
func (p *astPrinter) parenthesize(
	name string,
	exprs ...Expr[string],
) string {
	var b strings.Builder
	b.WriteString("(")
	b.WriteString(name)
	for _, expr := range exprs {
		b.WriteString(" ")
		b.WriteString(expr.Accept(p))
	}
	b.WriteString(")")
	return b.String()
}
