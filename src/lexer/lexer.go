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

func (l *Lexer) peekChar() byte {
	if l.currReadPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.currReadPosition]
	}
}

func (l *Lexer) NextToken() token.Token {

	l.skipWhiteSpaces()

	var tok token.Token

	switch l.ch {
	case '+':
		tok = createToken(token.PLUS, l.ch)
	case '-':
		tok = createToken(token.MINUS, l.ch)
	case '*':
		tok = createToken(token.ASTERISK, l.ch)
	case '/':
		tok = createToken(token.SLASH, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = createToken(token.BANG, l.ch)
		}
	case '<':
		tok = createToken(token.LT, l.ch)
	case '>':
		tok = createToken(token.GT, l.ch)
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
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = createToken(token.ASSIGN, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = createToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func createToken(Type token.TokenType, ch byte) token.Token {
	return token.Token{Type: Type, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}

func (l *Lexer) readIdentifier() string {
	initialPosition := l.lastReadPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[initialPosition:l.lastReadPosition]
}

func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')
}

func (l *Lexer) readNumber() string {
	initialPosition := l.lastReadPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[initialPosition:l.lastReadPosition]
}
