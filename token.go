package main

type TokenType string

const (
	// Single Character or Number
	NUMBER = "NUMBER"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	LPAREN   = "("
	RPAREN   = ")"

	// End Of File
	EOF = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
