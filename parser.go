package main

const (
	_ int = iota
	LOWEST
	SUM     // +
	PRODUCT // *
	PREFIX  // -X
	POSTFIX // X!
)

var precedences = map[TokenType]int{
	PLUS:     SUM,
	MINUS:    SUM,
	ASTERISK: PRODUCT,
	SLASH:    PRODUCT,
	BANG:     POSTFIX,
}

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
	expr := p.parseExpression(LOWEST)
	if expr == nil {
		return ""
	}
	return expr.String()
}

func (p *Parser) parseExpression(precedence int) Expression {
	var leftExp Expression
	switch p.curToken.Type {
	case NUMBER:
		leftExp = p.parseNumberExpression()
	case MINUS:
		leftExp = p.parsePrefixExpression()
	case LPAREN:
		leftExp = p.parseParenExpression()
	default:
		leftExp = nil
	}

	for p.peekToken.Type != EOF {
		switch p.peekToken.Type {
		case PLUS, MINUS, ASTERISK, SLASH:
			if precedence >= p.peekPrecedence() {
				return leftExp
			}
			p.nextToken()
			leftExp = p.parseInfixExpression(leftExp)
		case BANG:
			if precedence >= p.peekPrecedence() {
				return leftExp
			}
			p.nextToken()
			leftExp = p.parsePostExpression(leftExp)
		default:
			return leftExp
		}
	}
	return leftExp
}

func (p *Parser) parseParenExpression() Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if p.peekToken.Type == RPAREN {
		p.nextToken()
	}
	return exp
}

func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parsePostExpression(left Expression) Expression {
	expression := &PostfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	return expression
}

func (p *Parser) parseNumberExpression() Expression {
	return &Number{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
