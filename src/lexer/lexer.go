package lexer

import (
	"github.com/RAMESHSM9/interpreter_go/src/token"
)

type Lexer struct {
	input            string
	lastReadPosition int
	currReadPosition int
	ch               byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.currReadPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.currReadPosition]
	}
	l.lastReadPosition = l.currReadPosition
	l.currReadPosition += 1
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	switch l.ch {
	case '+':
		tok = createToken(token.PLUS, l.ch)
	case '(':
		tok = createToken(token.LPAREN, l.ch)
	case ')':
		tok = createToken(token.RPAREN, l.ch)
	case '{':
		tok = createToken(token.LBRACES, l.ch)
	case '}':
		tok = createToken(token.RBRACES, l.ch)
	case ',':
		tok = createToken(token.COMMA, l.ch)
	case ';':
		tok = createToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func createToken(Type token.TokenType, ch byte) token.Token {
	return token.Token{Type: Type, Literal: string(ch)}
}
