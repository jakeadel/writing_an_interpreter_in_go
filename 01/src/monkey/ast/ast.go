package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node // This means anything that implements Statement must also implement Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// The root node of the ast
// Stores all program statements in Statements
type Program struct {
	Statements []Statement // slice
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// the "let" part
type LetStatement struct {
	Token token.Token // Going to be token.LET
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

// variable name
type Identifier struct {
	Token token.Token // Going to be token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
