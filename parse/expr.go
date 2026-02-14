package parse

import (
	"fmt"
	"strings"
	"github.com/GrazianoJoa/Glox/scan"
)

// This implementation uses the Visitor Pattern.
// In Go, it's often more idiomatic to use a type switch instead.

type Expr interface {
	Accept(v Visitor) any
}

type Visitor interface {
	VisitLiteral(l *Literal) any
	VisitBinary(b *Binary) any
	VisitGrouping(g *Grouping) any
	VisitUnary(u *Unary) any
}

// Concrete Visitor

type VisitorPrint struct {}

func (p *VisitorPrint) VisitLiteral(l *Literal) any {
	if l.value == nil { 
		return "nil"
	}
	return l.value
}

func (p *VisitorPrint) VisitUnary(u *Unary) any {
	return p.parenthesize(u.operator.Lexeme, u.right)
}

func (p *VisitorPrint) VisitBinary(b *Binary) any {
	return p.parenthesize(b.operator.Lexeme, b.right, b.left)
}

func (p *VisitorPrint) VisitGrouping(g *Grouping) any {
	return p.parenthesize("grouping", g.expr)
}

func (p *VisitorPrint) parenthesize(name string, exprs... Expr) string {
	var b strings.Builder
	b.WriteString("(")
	for _, expr := range exprs {
		b.WriteString(" ")
		e := expr.Accept(p)
		b.WriteString(fmt.Sprint(e))
	}
	b.WriteString(")")
	return b.String()
}
 
// Binary

type Binary struct {
	left Expr
	operator *scan.Token
	right Expr
}

func (b *Binary) Accept(v Visitor) any {
	return v.VisitBinary(b)
}

// Grouping

type Grouping struct {
	expr Expr
}

func (g *Grouping) Accept(v Visitor) any {
	return v.VisitGrouping(g)
}

// Literal

type Literal struct {
	value any
}

func (l *Literal) Accept(v Visitor) any {
	return v.VisitLiteral(l)
}

// Unary

type Unary struct {
	right Expr
	operator *scan.Token
}

func (u *Unary) Accept(v Visitor) any {
	return v.VisitUnary(u)
}

