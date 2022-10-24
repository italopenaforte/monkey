package parser

import (
	"github.com/italopenaforte/monkey/ast"
	"github.com/italopenaforte/monkey/lexer"
	"github.com/italopenaforte/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // current token
	peekToken token.Token // see next token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Will read two of the tokens, so curToken get the current token and peek token get the next token they are both set initially
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {

	// will get the current and the next token to make a decision what to do with
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// build the parse to pass the tests
	return nil
}
