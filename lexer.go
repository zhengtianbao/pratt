package main

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = NewToken(PLUS, l.ch)
	case '-':
		tok = NewToken(MINUS, l.ch)
	case '/':
		tok = NewToken(SLASH, l.ch)
	case '*':
		tok = NewToken(ASTERISK, l.ch)
	case '!':
		tok = NewToken(BANG, l.ch)
	case '(':
		tok = NewToken(LPAREN, l.ch)
	case ')':
		tok = NewToken(RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		tok = NewToken(NUMBER, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
