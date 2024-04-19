package main

import (
	"bytes"
)

// The base Node interface
type Node interface {
	String() string
}

// All expression nodes implement this
type Expression interface {
	Node
}

type Number struct {
	Token Token
	Value string
}

func (i *Number) String() string { return i.Value }

type PrefixExpression struct {
	Token    Token // The prefix token, e.g. -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
