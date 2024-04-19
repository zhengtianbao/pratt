package main

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l: l,
	}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() string {
	switch p.curToken.Type {
	case NUMBER:
		return p.parseNumberExpression().String()
	default:
		return ""
	}
}

func (p *Parser) parseNumberExpression() Expression {
	return &Number{Token: p.curToken, Value: p.curToken.Literal}
}
