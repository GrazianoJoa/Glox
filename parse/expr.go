package parse

import (
	"fmt"
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
	return fmt.Sprintf("%v", l.value)
}

func (p *VisitorPrint) VisitUnary(u *Unary) any {
	return fmt.Sprintf("")
}

func (p *VisitorPrint) VisitBinary(b *Binary) any {
	return ""
}

func (p *VisitorPrint) VisitGrouping(g *Grouping) any {
	return ""
}
 
// Binary

type Binary struct {
	left *Expr
	operator *scan.Token
	right *Expr
}

func (b *Binary) Accept(v Visitor) any {
	return v.VisitBinary(b)
}

// Grouping

type Grouping struct {
	expr *Expr
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
	right *Expr
	operator *scan.Token
}

func (u *Unary) Accept(v Visitor) any {
	return v.VisitUnary(u)
}

func main() {
	expr := &Literal{value: 45}
	printer := &VisitorPrint{}
	res := expr.Accept(printer)
	fmt.Println(res)
}
