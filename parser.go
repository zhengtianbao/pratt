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
	expr := p.parseExpression()
	if expr == nil {
		return ""
	}
	return expr.String()
}

func (p *Parser) parseExpression() Expression {
	var leftExp Expression
	switch p.curToken.Type {
	case NUMBER:
		leftExp = p.parseNumberExpression()
	case MINUS:
		leftExp = p.parsePrefixExpression()
	default:
		leftExp = nil
	}
	return leftExp
}

func (p *Parser) parseNumberExpression() Expression {
	return &Number{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression()
	return expression
}
